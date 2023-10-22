package helper

import (
	"encoding/json"
	"github.com/beevik/etree"
	"github.com/cocktail18/wxhelper-go/api"
	"github.com/cocktail18/wxhelper-go/proto"
	"github.com/spf13/cast"
	"strings"
	"time"
)

func DecodePrivateMsg(apiVersion api.ApiVersion, bs []byte) (*proto.WxPrivateMsg, error) {
	var privateMsg proto.WxPrivateMsg
	// 处理群消息
	var err error
	if apiVersion == api.ApiVersionV1 {
		var msgV1 proto.PrivateMsgV1
		err = json.Unmarshal(bs, &msgV1)
		if err != nil {
			return nil, err
		}
		privateMsg.Content = msgV1.Content
		privateMsg.FromGroup = msgV1.FromGroup
		privateMsg.FromUser = msgV1.FromUser
		privateMsg.IsSendMsg = msgV1.IsSendMsg
		privateMsg.MsgId = msgV1.MsgId
		privateMsg.Pid = msgV1.Pid
		privateMsg.Sign = msgV1.Sign
		privateMsg.Signature = msgV1.Signature
		privateMsg.Time = msgV1.Time
		privateMsg.Timestamp = msgV1.Timestamp
		privateMsg.Type = msgV1.Type
		privateMsg.DisplayFullContent = msgV1.Content
	} else {
		var msgV2 proto.PrivateMsgV2
		err = json.Unmarshal(bs, &msgV2)
		if err != nil {
			return nil, err
		}
		privateMsg.Content = msgV2.Content
		privateMsg.FromUser = msgV2.FromUser
		privateMsg.IsSendMsg = 0
		privateMsg.MsgId = msgV2.MsgId
		privateMsg.Pid = msgV2.Pid
		privateMsg.Sign = msgV2.Signature
		privateMsg.Signature = msgV2.Signature
		privateMsg.Time = time.Unix(int64(msgV2.CreateTime), 0).Format("2006:01:02 15:04:05")
		privateMsg.Timestamp = msgV2.CreateTime
		privateMsg.Type = msgV2.Type
		privateMsg.DisplayFullContent = msgV2.Content
		if strings.Contains(privateMsg.FromUser, "@") {
			privateMsg.FromGroup = privateMsg.FromUser
			privateMsg.FromUser, privateMsg.Content = getWxIdAndContentFromMsgContent(privateMsg.Content)

		}
	}
	privateMsg.AtWxIds = make([]string, 0)
	doc := etree.NewDocument()
	if err := doc.ReadFromString(privateMsg.Signature); err == nil {
		if doc.FindElement("/msgsource/membercount") != nil {
			privateMsg.GroupMemberCount = cast.ToInt(doc.FindElement("/msgsource/membercount").Text())
		}
		if doc.FindElement("/msgsource/atuserlist") != nil {
			atWxIds := doc.FindElement("/msgsource/atuserlist").Text()
			if atWxIds != "" {
				privateMsg.AtWxIds = strings.Split(atWxIds, ",")
			}
		}
	}
	return &privateMsg, nil
}

func getWxIdAndContentFromMsgContent(content string) (wxId string, realContent string) {
	idx := strings.Index(content, ":\n")
	return content[:idx], content[idx+2:]
}
