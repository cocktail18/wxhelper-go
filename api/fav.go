package api

import (
	"github.com/cocktail18/wxhelper-go/util"
)

func (api *Api) AddFavFromMsg(msgId string) error {
	url, err := api.getUrl(AddFavFromMsgUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"msgId": msgId,
	})
	return err
}

func (api *Api) AddFavFromImage(wxId string, imagePath string) error {
	url, err := api.getUrl(AddFavFromImageUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, map[string]interface{}{
		"wxId": wxId, "imagePath": imagePath,
	})
	return err
}
