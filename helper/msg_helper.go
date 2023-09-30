package helper

import (
	"encoding/json"
	"github.com/beevik/etree"
	"github.com/cocktail18/wxhelper-go/proto"
	"github.com/spf13/cast"
	"strings"
)

func DecodePrivateMsg(bs []byte) (*proto.PrivateMsg, error) {
	var privateMsg proto.PrivateMsg
	err := json.Unmarshal(bs, &privateMsg)
	if err != nil {
		return nil, err
	}
	// 处理群消息
	switch privateMsg.Type {
	case proto.MsgTypeChat:
		if strings.Contains(privateMsg.FromUser, "@") {
			privateMsg.IsGroup = true
			privateMsg.GroupId = privateMsg.FromUser
			privateMsg.FromUser, privateMsg.Content = getWxIdAndContentFromMsgContent(privateMsg.Content)
			doc := etree.NewDocument()
			if err := doc.ReadFromString(privateMsg.Signature); err != nil {
				privateMsg.GroupMemberCount = cast.ToInt(doc.FindElement("/msgsource/membercount").Text())
			}

		}
	}
	return &privateMsg, nil
}

func getWxIdAndContentFromMsgContent(content string) (wxId string, realContent string) {
	idx := strings.Index(content, ":\n")
	return content[:idx], content[idx+2:]
}
