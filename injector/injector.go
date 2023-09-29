package injector

import (
	"github.com/cocktail18/wx-helper-go/api"
	"github.com/pkg/errors"
	"go.zoe.im/injgo"
	"go.zoe.im/injgo/pkg/w32"
	"golang.org/x/exp/slog"
	"math"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"
)

var (
	ErrWxProcessNotFound  = errors.New("微信进程不存在")
	kernel32              = syscall.NewLazyDLL("kernel32.dll")
	procGetExitCodeThread = kernel32.NewProc("GetExitCodeThread")
)

func InjectWx(version api.ApiVersion, dllPath string, httpPort int) error {
	processList, err := FindProcessListByName("wechat.exe")
	if err != nil {
		return errors.Wrap(err, "查找进程失败")
	}
	if len(processList) <= 0 {
		return ErrWxProcessNotFound
	}
	for _, process := range processList {
		if version == api.ApiVersionV1 {
			err = InjectAndStartHttpServer(process.ProcessID, httpPort, dllPath, true)
		} else {
			err = injgo.Inject(process.ProcessID, dllPath, true)
		}
		if err != nil {
			return errors.Wrap(err, "注入失败")
		}
		return nil
	}
	return nil
}

func InjectByProcess(process *Process, dllPath string) error {
	err := injgo.Inject(process.ProcessID, dllPath, true)
	if err != nil {
		return errors.Wrap(err, "注入失败")
	}
	return nil
}

func StartWxProcess() (*Process, error) {
	// 启动微信 并且注入
	wxInstallPath, err := GetWxInstallPath()
	if err != nil {
		return nil, err
	}
	if wxInstallPath == "" {
		return nil, errors.New("找不到微信安装路径，请检查是否已经正确安装微信")
	}
	process, err := injgo.CreateProcess(wxInstallPath)
	if err != nil {
		return nil, err
	}
	return &Process{
		ProcessID: process.ProcessID,
		Name:      process.Name,
		ExePath:   process.ExePath,
	}, nil
}

func InjectAndStartHttpServer(pid, httpPort int, dllname string, replace bool) error {
	dllname, _ = filepath.Abs(dllname)
	_, err := os.Stat(dllname)
	if os.IsNotExist(err) {
		return err
	}
	// check is already injected
	if !replace && injgo.IsInjected(pid, dllname) {
		return injgo.ErrAlreadyInjected
	}

	// open process
	hdlr, err := w32.OpenProcess(w32.PROCESS_ALL_ACCESS, true, ptr(pid))
	if err != nil {
		return err
	}
	defer w32.CloseHandle(hdlr)

	startHttpFunc := "http_start"
	// malloc space to write dll name
	startHttpFuncLen := len(startHttpFunc) + 1
	dlllen := len(dllname) + 1
	dllnameaddr, err := w32.VirtualAllocEx(hdlr, 0, ptr(dlllen), ptr(w32.MEM_RESERVE_AND_COMMIT), ptr(w32.PAGE_READWRITE))
	if err != nil {
		return err
	}
	startHttpFuncAddr, err := w32.VirtualAllocEx(hdlr, 0, ptr(startHttpFuncLen), ptr(w32.MEM_RESERVE_AND_COMMIT), ptr(w32.PAGE_READWRITE))
	if err != nil {
		return err
	}
	params := []uintptr{dllnameaddr, startHttpFuncAddr}
	paramsAddr, err := w32.VirtualAllocEx(hdlr, 0, ptr(8), ptr(w32.MEM_RESERVE_AND_COMMIT), ptr(w32.PAGE_READWRITE))
	if err != nil {
		return err
	}

	// write dll name
	err = w32.WriteProcessMemory(hdlr, dllnameaddr, ptr(dllname), ptr(dlllen))
	if err != nil {
		return err
	}

	defer func() {
		_, err = w32.VirtualFreeEx(hdlr, dllnameaddr, ptr(0), w32.MEM_RELEASE)
		slog.Error("VirtualFreeEx dll error:%v", err)
	}()

	// write start http function
	err = w32.WriteProcessMemory(hdlr, startHttpFuncAddr, ptr(startHttpFunc), ptr(startHttpFuncLen))
	if err != nil {
		return err
	}

	defer func() {
		_, err = w32.VirtualFreeEx(hdlr, startHttpFuncAddr, ptr(0), w32.MEM_RELEASE)
		slog.Error("VirtualFreeEx dll error:%v", err)
	}()

	// write start http function
	err = w32.WriteProcessMemory(hdlr, paramsAddr, uintptr(unsafe.Pointer(&params[0])), ptr(8))
	if err != nil {
		return err
	}

	// test
	tecase, _ := w32.ReadProcessMemory(hdlr, dllnameaddr, ptr(dlllen))
	if string(tecase[:len(tecase)-1]) != dllname {
		return errors.New("write dll name error")
	}

	// get LoadLibraryA address in target process
	// TODO: can we get the address at from this process?
	lddladdr, err := w32.LoadLibraryAddress(ptr("LoadLibraryA"))
	if err != nil {
		return err
	}

	// call remote process
	dllthread, _, err := w32.CreateRemoteThread(hdlr, nil, ptr(0), ptr(lddladdr), paramsAddr, ptr(0))
	if err != nil {
		return err
	}
	defer w32.CloseHandle(dllthread)
	err = w32.WaitForSingleObj(dllthread, math.MaxInt)
	if err != nil {
		return err
	}

	exitCode, err := GetExitCodeThread(syscall.Handle(dllthread))
	if err != nil {
		return err
	}

	hStartHttp, _, err := w32.CreateRemoteThread(hdlr, nil, ptr(0), ptr(exitCode), ptr(httpPort), ptr(0))
	if err != nil {
		return err
	}
	defer w32.CloseHandle(hStartHttp)
	err = w32.WaitForSingleObj(hStartHttp, math.MaxInt)
	if err != nil {
		return err
	}
	return nil
}

func ptr(val interface{}) uintptr {
	switch val.(type) {
	case byte:
		return uintptr(val.(byte))
	case bool:
		isTrue := val.(bool)
		if isTrue {
			return uintptr(1)
		}
		return uintptr(0)
	case string:
		bytePtr, _ := syscall.BytePtrFromString(val.(string))
		return uintptr(unsafe.Pointer(bytePtr))
	case int:
		return uintptr(val.(int))
	case uint:
		return uintptr(val.(uint))
	case uintptr:
		return val.(uintptr)
	default:
		return uintptr(0)
	}
}

func GetExitCodeThread(handle syscall.Handle) (uintptr, error) {
	var exitCode *uint32
	_, _, err := procGetExitCodeThread.Call(uintptr(handle), uintptr(unsafe.Pointer(exitCode)))
	if w32.IsErrSuccess(err) {
		return uintptr(unsafe.Pointer(exitCode)), nil
	}
	return 0, err
}
