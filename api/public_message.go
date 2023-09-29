package api

import (
	"github.com/cocktail18/wx-helper-go/util"
)

func (api *Api) ForwardPublicMsgByMsgId(wxId string, msgId string) error {
	url, err := api.getUrl(ForwardPublicMsgByMsgIdUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"wxId": wxId, "msgId": msgId,
	})
	return err
}

/*
appName	string	公众号id，消息内容里的appname
userName	string	公众号昵称，消息内容里的username
title	string	链接地址，消息内容里的title
url	string	链接地址，消息内容里的url
thumbUrl	string	缩略图地址，消息内容里的thumburl
digest	string	摘要，消息内容里的digest
wxid	string	wxid
*/
func (api *Api) ForwardPublicMsg(appName, userName, title, url, thumbUrl, digest, wxid string) error {
	url, err := api.getUrl(ForwardPublicMsgUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"appName": appName, "userName": userName, "title": title, "url": url, "thumbUrl": thumbUrl, "digest": digest, "wxid": wxid,
	})
	return err
}
