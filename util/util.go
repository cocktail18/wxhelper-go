package util

import (
	"errors"
	"fmt"
	"github.com/cocktail18/wxhelper-go/proto"
	"github.com/imroc/req/v3"
	"golang.org/x/exp/slog"
	"net"
	"runtime/debug"
)

var (
	defaultHttpClient = req.C() // Use C() to create a client.
)

func Request(url string, data interface{}) (*proto.Response, error) {
	result := proto.Response{}
	resp, err := defaultHttpClient.R().
		SetBody(data).
		SetSuccessResult(&result).
		Post(url)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccessState() {
		return nil, fmt.Errorf("statusCode:%v msg:%v", resp.StatusCode, string(resp.Bytes()))
	}
	if result.Code == 0 && result.Result == "ERROR" {
		return nil, errors.New("接口调用失败：" + result.Msg)
	}
	return &result, nil
}

func GetRandomAvailablePort() (int, error) {
	// 监听一个随机的端口
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	defer l.Close()

	// 获取监听的地址
	addr := l.Addr().(*net.TCPAddr)
	return addr.Port, nil
}

func RecoveryGo(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("recovery: ", "err", err, "stack", debug.Stack())
			}
		}()
		f()
	}()
}
