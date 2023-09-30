package injector

import (
	"github.com/pkg/errors"
	"go.zoe.im/injgo/pkg/w32"
	"golang.org/x/sys/windows/registry"
	"strings"
	"syscall"
	"unsafe"
)

// Process ...
type Process struct {
	ProcessID int
	Name      string
	ExePath   string
}

var (
	// ErrCreateSnapshot ...
	ErrCreateSnapshot        = errors.New("create snapshot error")
	ErrWxInstallPathNotFound = errors.New("找不到微信安装路径")
)

// FindProcessListByName get process information by name
func FindProcessListByName(name string) ([]*Process, error) {
	handle, _ := syscall.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPPROCESS, 0)
	if handle == 0 {
		return nil, ErrCreateSnapshot
	}
	defer syscall.CloseHandle(handle)

	ret := make([]*Process, 0)

	var entry = syscall.ProcessEntry32{}
	entry.Size = uint32(unsafe.Sizeof(entry))
	var process Process

	for true {
		if nil != syscall.Process32Next(handle, &entry) {
			break
		}

		_exeFile := w32.UTF16PtrToString(&entry.ExeFile[0])
		if strings.ToLower(name) == strings.ToLower(_exeFile) {
			process.Name = _exeFile
			process.ProcessID = int(entry.ProcessID)
			process.ExePath = _exeFile
			ret = append(ret, &process)
		}

	}
	return ret, nil
}

func GetWxInstallPath() (string, error) {
	// 打开 "HKEY_CURRENT_USER\Software\Tencent\WeChat" 键
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Tencent\WeChat`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	// 读取 "InstallPath" 值
	val, valKind, err := k.GetStringValue("InstallPath")
	if err != nil {
		return "", errors.Wrap(err, "无法读取注册表值")
	}

	// 判断值类型是否为字符串
	if valKind != registry.SZ && valKind != registry.EXPAND_SZ {
		return "", errors.Wrap(err, "注册表值类型错误")
	}

	// 转换成Windows API所需的LPWSTR格式
	installPathPtr, err := syscall.UTF16PtrFromString(val)
	if err != nil {
		return "", errors.Wrap(err, "路径转换出错")
	}

	// 使用Windows API获取真实路径
	realPath := make([]uint16, syscall.MAX_PATH)
	_, err = syscall.GetLongPathName(installPathPtr, &realPath[0], syscall.MAX_PATH)
	if err != nil {
		return "", errors.Wrap(err, "无法获取真实路径")
	}

	// 将UTF16编码转换为Go字符串
	realPathStr := syscall.UTF16ToString(realPath)
	return realPathStr, nil
}
