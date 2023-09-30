package api

import (
	"github.com/cocktail18/wxhelper-go/tcpserver"
	"github.com/cocktail18/wxhelper-go/util"
	"github.com/pkg/errors"
	"net"
	"strings"
	"sync"
)

var (
	ErrMethodNotImplement = errors.New("方法未实现")
)

type Api struct {
	BaseUrl           string
	ApiVersion        ApiVersion
	MsgListenInstance *MsgListenServer
}

type MsgListenServer struct {
	Lock        sync.Mutex
	Port        int
	Listener    net.Listener
	HookSuccess bool
	Callback    func(bs []byte)
	MsgChannel  chan []byte
}

func NewApi(apiVersion ApiVersion, baseUrl string) *Api {
	msgListenServer := &MsgListenServer{
		Callback:    nil,
		MsgChannel:  make(chan []byte, 512),
		HookSuccess: false,
	}
	go func() {
		for bs := range msgListenServer.MsgChannel {
			if msgListenServer.Callback == nil {
				continue
			}
			msgListenServer.Callback(bs)
		}
	}()
	return &Api{BaseUrl: strings.TrimRight(baseUrl, "/"), ApiVersion: apiVersion, MsgListenInstance: msgListenServer}
}

func (api *Api) getUrl(url ApiUrl) (string, error) {
	if str, ok := getUrlMapper(api.ApiVersion)[url]; ok {
		return api.BaseUrl + str, nil
	}
	return "", ErrMethodNotImplement
}

func (api *Api) SetMsgCallback(callback func(bs []byte)) error {
	api.MsgListenInstance.Lock.Lock()
	defer api.MsgListenInstance.Lock.Unlock()

	if api.MsgListenInstance.Listener == nil {
		var ln net.Listener
		var randPort int
		var err error
		for i := 0; i < 3; i++ {
			randPort, err = util.GetRandomAvailablePort()
			if err != nil {
				continue
			}
			ln, err = tcpserver.Listen(randPort, func(bs []byte) {
				api.MsgListenInstance.MsgChannel <- bs
			})
			if err == nil {
				break
			}
		}
		if err != nil {
			return errors.WithMessage(err, "监听随机端口失败")
		}
		api.MsgListenInstance.Listener = ln
		api.MsgListenInstance.Port = randPort
	}
	if !api.MsgListenInstance.HookSuccess {
		err := api.HookMsg(api.MsgListenInstance.Port)
		if err != nil {
			return errors.WithMessage(err, "hook接收消息失败")
		}
		api.MsgListenInstance.HookSuccess = true
	}

	api.MsgListenInstance.Callback = callback
	return nil
}
