package util

import (
	"fmt"
	"github.com/cocktail18/wx-helper-go/proto"
	"github.com/imroc/req/v3"
	"net"
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
