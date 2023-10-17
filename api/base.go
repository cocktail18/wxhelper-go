package api

import (
	"encoding/json"
	"github.com/cocktail18/wxhelper-go/proto"
	"github.com/cocktail18/wxhelper-go/util"
	"github.com/pkg/errors"
	"path/filepath"
)

func (api *Api) CheckLogin() (bool, error) {
	if api.ApiVersion == ApiVersionV1 {
		return false, errors.New("v1有崩溃，不能调用")
	}
	url, err := api.getUrl(CheckLoginUrl)
	if err != nil {
		return false, err
	}
	resp, err := util.Request(url, nil)
	if err != nil {
		return false, err
	}
	return resp.Code == 1, nil
}

func (api *Api) GetUserInfo() (*proto.WxUserInfo, error) {
	url, err := api.getUrl(UserInfoUrl)
	if err != nil {
		return nil, err
	}
	resp, err := util.Request(url, nil)
	if err != nil {
		return nil, err
	}
	var userInfo proto.WxUserInfo
	err = json.Unmarshal(resp.Data, &userInfo)
	return &userInfo, err
}

func (api *Api) SendTextMsg(wxid, msg string) error {
	url, err := api.getUrl(SendTextMsgUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"wxid": wxid, "msg": msg,
	})
	return err
}

func (api *Api) SendFileMsg(wxid, filePath string) error {
	filePath, _ = filepath.Abs(filePath)
	url, err := api.getUrl(SendFileMsgUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"wxid": wxid, "filePath": filePath,
	})
	return err
}

func (api *Api) SendImagesMsg(wxid, imagePath string) error {
	imagePath, _ = filepath.Abs(imagePath)
	url, err := api.getUrl(SendImagesMsgUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"wxid": wxid, "imagePath": imagePath,
	})
	return err
}

// 表情路径，可以直接查询CustomEmotion表的MD5字段
// "filePath":"C:\\wechatDir\\WeChat Files\\wxid_123\\FileStorage\\CustomEmotion\\8F\\8F6423BC2E69188DCAC797E279C81DE9"
func (api *Api) SendCustomEmotion(wxid, filePath string) error {
	filePath, _ = filepath.Abs(filePath)
	url, err := api.getUrl(SendCustomEmotionUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"wxid": wxid, "filePath": filePath,
	})
	return err
}

// 发送小程序（待完善，不稳定）,相关参数可以参考示例的滴滴小程序的内容自行组装。
func (api *Api) SendApplet(req *proto.SendAppletMsgReq) error {
	url, err := api.getUrl(SendAppletUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, req)
	return err
}

// wxid	string	被拍人id
// receiver	string	接收人id，可以是自己wxid，私聊好友wxid，群id
func (api *Api) SendPatMsg(wxid string, receiver string) error {
	url, err := api.getUrl(SendPatMsgUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"wxid": wxid, "receiver": receiver,
	})
	return err
}

func (api *Api) Ocr(imagePath string) (string, error) {
	url, err := api.getUrl(OcrUrl)
	if err != nil {
		return "", err
	}
	resp, err := util.Request(url, map[string]interface{}{
		"imagePath": imagePath,
	})
	if err != nil {
		return "", err
	}
	return string(resp.Data), nil
}

func (api *Api) ForwardMsg(wxid, msgId string) error {
	url, err := api.getUrl(ForwardMsgUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"wxid": wxid, "msgId": msgId,
	})
	return err
}

func (api *Api) HookMsg(port int) error {
	for i := 0; i < 3; i++ {
		url, err := api.getUrl(HookSyncMsgUrl)
		if err != nil {
			return err
		}
		resp, err := util.Request(url, map[string]interface{}{
			"port": port, "ip": "127.0.0.1", "enableHttp": 0, "timeout": 60 * 1000, "url": "",
		})
		if err != nil {
			return err
		}
		if resp.Code == 2 { // 已经hook过
			err = api.UnHookMsg()
			if err != nil {
				return err
			}
		} else { // hook 成功
			break
		}
	}

	return nil
}

func (api *Api) UnHookMsg() error {
	url, err := api.getUrl(UnHookSyncMsgUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, nil)
	return err
}

func (api *Api) GetContactList() ([]proto.ContactInfo, error) {
	url, err := api.getUrl(GetContactListUrl)
	if err != nil {
		return nil, err
	}
	resp, err := util.Request(url, nil)
	if err != nil {
		return nil, err
	}
	var contactInfoList = make([]proto.ContactInfo, 0)
	err = json.Unmarshal(resp.Data, &contactInfoList)
	return contactInfoList, err
}

func (api *Api) GetContactProfile(wxid string) (*proto.ContactProfile, error) {
	url, err := api.getUrl(GetContactProfileUrl)
	if err != nil {
		return nil, err
	}
	resp, err := util.Request(url, map[string]interface{}{
		"wxid": wxid,
	})
	if err != nil {
		return nil, err
	}
	var roomMember proto.ContactProfile
	err = json.Unmarshal(resp.Data, &roomMember)
	return &roomMember, err
}

func (api *Api) GetContactNickname(wxidOrGroupId string) (string, error) {
	// 看起来有bug，返回都是空
	url, err := api.getUrl(GetContactNicknameUrl)
	if err != nil {
		return "", err
	}
	resp, err := util.Request(url, map[string]interface{}{
		"id": wxidOrGroupId,
	})
	if err != nil {
		return "", err
	}
	return string(resp.Data), err
}

// ConfirmFriendRequest
// permission 好友权限，0.全部 8.仅聊天。（好友超过5000人后，会验证权限，传8即可）
func (api *Api) ConfirmFriendRequest(encryptUsername, ticket string, permission int) error {
	// 看起来有bug，返回都是空
	url, err := api.getUrl(ConfirmFriendRequest)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"v3":         encryptUsername,
		"v4":         ticket,
		"permission": permission,
	})
	return err
}
