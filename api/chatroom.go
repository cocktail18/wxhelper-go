package api

import (
	"encoding/json"
	"github.com/cocktail18/wxhelper-go/proto"
	"github.com/cocktail18/wxhelper-go/util"
	"strings"
)

// GetChatRoomDetailInfoUrl
// param chatRoomId: "12222@chatroom"
func (api *Api) GetChatRoomDetailInfo(chatRoomId string) (*proto.ChatroomDetail, error) {
	url, err := api.getUrl(GetChatRoomDetailInfoUrl)
	if err != nil {
		return nil, err
	}
	resp, err := util.Request(url, map[string]interface{}{
		"chatRoomId": chatRoomId,
	})
	if err != nil {
		return nil, err
	}
	var roomDetail proto.ChatroomDetail
	err = json.Unmarshal(resp.Data, &roomDetail)
	return &roomDetail, err
}

func (api *Api) GetMemberFromChatRoom(chatRoomId string) (*proto.ChatroomMember, error) {
	url, err := api.getUrl(GetMemberFromChatRoomUrl)
	if err != nil {
		return nil, err
	}
	resp, err := util.Request(url, map[string]interface{}{
		"chatRoomId": chatRoomId,
	})
	if err != nil {
		return nil, err
	}
	var roomMember proto.ChatroomMember
	err = json.Unmarshal(resp.Data, &roomMember)
	return &roomMember, err
}

func (api *Api) GetNicknameFromChatRoom(chatRoomId, memberId string) (string, error) {
	if api.ApiVersion == ApiVersionV1 {
		url, err := api.getUrl(GetNicknameFromChatRoomUrl)
		if err != nil {
			return "", err
		}
		resp, err := util.Request(url, map[string]interface{}{
			"chatRoomId": chatRoomId,
			"memberId":   memberId,
		})
		if err != nil {
			return "", err
		}
		return resp.Nickname, err
	} else {
		data, err := api.GetMemberFromChatRoom(chatRoomId)
		if err != nil {
			return "", err
		}
		if chatRoomId == memberId { //获取群主
			memberId = data.Admin
		}
		memberList := make([]string, 0)
		if !strings.Contains(data.Members, "^G") {
			memberList = append(memberList, data.Members)
		} else {
			memberList = strings.Split(data.Members, "^G")
		}
		nicknameList := make([]string, 0)
		if !strings.Contains(data.MemberNickname, "^G") {
			nicknameList = append(nicknameList, data.MemberNickname)
		} else {
			nicknameList = strings.Split(data.MemberNickname, "^G")
		}
		nickname := ""
		for i := 0; i < len(memberList); i++ {
			if memberList[i] == memberId {
				nickname = nicknameList[i]
			}
		}
		if nickname == "" {
			profile, err := api.GetContactProfile(memberId)
			if err != nil {
				return "", err
			}
			return profile.Nickname, nil
		}
		return nickname, nil
	}

}

func (api *Api) AddMemberToChatRoom(chatRoomId string, members ...string) error {
	if len(members) <= 0 {
		return nil
	}
	url, err := api.getUrl(AddMemberToChatRoomUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"chatRoomId": chatRoomId, "memberIds": strings.Join(members, ","),
	})
	return err
}

func (api *Api) ModifyChatroomNickname(chatRoomId, wxid, nickName string) error {
	url, err := api.getUrl(ModifyNicknameUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"chatRoomId": chatRoomId, "nickName": nickName, "wxid": wxid,
	})
	return err
}

func (api *Api) DelMemberFromChatRoom(chatRoomId string, members ...string) error {
	if len(members) <= 0 {
		return nil
	}
	url, err := api.getUrl(DelMemberFromChatRoomUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"chatRoomId": chatRoomId, "memberIds": strings.Join(members, ","),
	})
	return err
}

func (api *Api) TopMsg(msgId string) error {
	url, err := api.getUrl(TopMsgUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"msgId": msgId,
	})
	return err
}

func (api *Api) RemoveTopMsg(chatRoomId, msgId string) error {
	url, err := api.getUrl(RemoveTopMsgUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"chatRoomId": chatRoomId, "msgId": msgId,
	})
	return err
}

func (api *Api) InviteMemberToChatRoom(chatRoomId string, members ...string) error {
	url, err := api.getUrl(InviteMemberToChatRoomUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"chatRoomId": chatRoomId, "memberIds": strings.Join(members, ","),
	})
	return err
}

func (api *Api) SendAtText(chatRoomId string, msg string, wxids ...string) error {
	url, err := api.getUrl(SendAtTextUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"chatRoomId": chatRoomId, "msg": msg, "wxids": strings.Join(wxids, ","),
	})
	return err
}

func (api *Api) SendAtAllText(chatRoomId string, msg string) error {
	return api.SendAtText(chatRoomId, msg, "notify@all")
}
