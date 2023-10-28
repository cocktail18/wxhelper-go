package api

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

func (api *Api) GetAvatars(wxids []string) (map[string]string, error) {
	dbInfoList, err := api.GetDBInfo()
	if err != nil {
		return nil, err
	}
	// 获取handler
	var handler int64
	for _, dbInfo := range dbInfoList {
		if dbInfo.DatabaseName == "Misc.db" {
			handler = dbInfo.Handle
		}
	}
	if handler == 0 {
		return nil, errors.New("找不到对应的数据库句柄")
	}
	params := make([]string, len(wxids))
	for i := 0; i < len(wxids); i++ {
		params[i] = "'" + wxids[i] + "'"
	}
	exexResult, err := api.ExecSql(handler, fmt.Sprintf("select usrName,smallHeadBuf from ContactHeadImg1 where usrName in (%s)", strings.Join(params, ",")))
	if err != nil {
		return nil, err
	}
	ret := make(map[string]string)
	for i, row := range exexResult {
		if i == 0 {
			continue
		}
		ret[row[0]] = row[1]
	}
	return ret, nil
}
