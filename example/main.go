package main

import (
	"fmt"
	"github.com/cocktail18/wxhelper-go/api"
	"github.com/cocktail18/wxhelper-go/helper"
	"github.com/cocktail18/wxhelper-go/injector"
	"github.com/cocktail18/wxhelper-go/proto"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"time"
)

const (
	apiVersion = api.ApiVersionV1 // 3.9.5.81 使用v2
	dllPath    = "wxhelper.dll"
	port       = 10086
)

func injectWx() {
	err := injector.InjectWx(apiVersion, dllPath, port)
	if errors.Is(err, injector.ErrWxProcessNotFound) {
		process, err2 := injector.StartWxProcess()
		if err2 != nil {
			panic(err2)
		}
		<-time.After(time.Second * 1)
		err = injector.InjectByProcess(apiVersion, process, dllPath, port)
	}
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	injectWx()

	ins := api.NewApi(apiVersion, "http://127.0.0.1:"+cast.ToString(port))
	err = ins.SetMsgCallback(func(bs []byte) {
		msg, err := helper.DecodePrivateMsg(apiVersion, bs)
		if err != nil {
			fmt.Println("解析消息失败：", err.Error())
			return
		}
		switch msg.Type {
		case proto.MsgTypeChat:
			fmt.Println("isGroup", msg.IsGroup(), "from", msg.FromUser, "groupId", msg.FromGroup, "content", msg.Content)
			if !msg.IsGroup() && msg.Content == "hello" {
				ins.SendTextMsg(msg.FromUser, "world")
			}
		default:
			fmt.Println("接收消息", string(bs))
		}
	})
	if err != nil {
		panic(err)
	}

	userInfo, err := ins.GetUserInfo()
	if err != nil {
		panic(err)
	}
	fmt.Println("userInfo", userInfo.Wxid, userInfo.Name)

	contactList, err := ins.GetContactList()
	if err != nil {
		panic(err)
	}
	for _, info := range contactList {
		fmt.Println("contact", info.Wxid, info.Nickname)

	}
	time.Sleep(time.Hour)
}
