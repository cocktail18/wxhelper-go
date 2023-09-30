package api

import (
	"github.com/cocktail18/wxhelper-go/util"
)

// 下载附件，保存在微信文件目录下 wxid_xxx/wxhelper 目录下
func (api *Api) DownloadAttach(msgId string) error {
	url, err := api.getUrl(DownloadAttachUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"msgId": msgId,
	})
	return err
}

func (api *Api) DecodeImage(src, storeDir string) error {
	url, err := api.getUrl(DecodeImageUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"filePath": src, "storeDir": storeDir,
	})
	return err
}

func (api *Api) GetVoiceByMsgId(msgId, storeDir string) error {
	url, err := api.getUrl(GetVoiceByMsgIdUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"msgId": msgId, "storeDir": storeDir,
	})
	return err
}
