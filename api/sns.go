package api

import (
	"github.com/cocktail18/wx-helper-go/util"
)

func (api *Api) GetSnsFirstPage() error {
	url, err := api.getUrl(GetSNSFirstPageUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, nil)
	return err
}

func (api *Api) GetSNSNextPage() error {
	url, err := api.getUrl(GetSNSNextPageUrl)
	if err != nil {
		return err
	}
	_, err = util.Request(url, nil)
	return err
}
