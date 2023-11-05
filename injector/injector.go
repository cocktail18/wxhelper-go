package injector

import (
	"fmt"
	"github.com/cocktail18/wxhelper-go/api"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"go.zoe.im/injgo"
	"golang.org/x/exp/slog"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

var (
	ErrWxProcessNotFound = errors.New("微信进程不存在")
)

func createConfigFile(port int) error {
	filePath, err := GetWxInstallPath()
	content := fmt.Sprintf("[config]\nport=%d", port)
	file, err := os.OpenFile(filePath+"/config.ini", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	// 相对路径也创建一份， 测试的时候发现用的是这份
	file2, err := os.OpenFile("./config.ini", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file2.Close()
	_, err = file2.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func InjectWx(injectorExePath string, version api.ApiVersion, dllPath string, httpPort int) error {
	processList, err := FindProcessListByName("wechat.exe")
	if err != nil {
		return errors.Wrap(err, "查找进程失败")
	}
	if len(processList) <= 0 {
		return ErrWxProcessNotFound
	}
	for _, process := range processList {
		return InjectByProcess(injectorExePath, version, process, dllPath, httpPort)
	}
	return nil
}

func InjectByProcess(injectorExePath string, version api.ApiVersion, process *Process, dllPath string, httpPort int) error {
	var err error
	// 写入配置文件，然后启动
	err = createConfigFile(httpPort)
	if err != nil {
		return errors.Wrap(err, "写入配置文件失败，请检查是否有权限")
	}
	slog.Info("开始注入", "pid", process.ProcessID, "port", httpPort)

	if version == api.ApiVersionV1 || injectorExePath != "" {
		err = InjectByCmd(injectorExePath, process.ProcessID, dllPath)
		if err != nil {
			return errors.Wrap(err, "注入失败")
		}
	} else if version == api.ApiVersionV2 {
		err = injgo.Inject(process.ProcessID, dllPath, true)
		if err != nil {
			return errors.Wrap(err, "注入失败")
		}
	} else {
		return errors.New("不支持的版本")
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
	process, err := injgo.CreateProcess(wxInstallPath + "/WeChat.exe")
	if err != nil {
		return nil, err
	}
	return &Process{
		ProcessID: process.ProcessID,
		Name:      process.Name,
		ExePath:   process.ExePath,
	}, nil
}

func InjectByCmd(injectorExePath string, pid int, dllname string) error {
	dllname, _ = filepath.Abs(dllname)
	_, err := os.Stat(dllname)
	if os.IsNotExist(err) {
		return err
	}
	// injector.exe 下载地址 https://github.com/nefarius/Injector  , 要使用win32版本
	robotCmd := exec.Command(injectorExePath, "-p", cast.ToString(pid), "-i", dllname)
	robotCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return robotCmd.Run()
}
