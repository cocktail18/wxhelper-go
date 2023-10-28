package api

import (
	"encoding/json"
	"github.com/cocktail18/wxhelper-go/proto"
	"github.com/cocktail18/wxhelper-go/util"
)

func (api *Api) GetDBInfo() ([]*proto.DbInfo, error) {
	url, err := api.getUrl(GetDBInfoUrl)
	if err != nil {
		return nil, err
	}
	resp, err := util.Request(url, nil)
	if err != nil {
		return nil, err
	}
	dbInfo := make([]*proto.DbInfo, 0)
	err = json.Unmarshal(resp.Data, &dbInfo)
	return dbInfo, err
}

func (api *Api) ExecSql(dbHandle int64, sql string) ([][]string, error) {
	url, err := api.getUrl(ExecSqlUrl)
	if err != nil {
		return nil, err
	}
	resp, err := util.Request(url, map[string]interface{}{
		"dbHandle": dbHandle, "sql": sql,
	})
	if err != nil {
		return nil, err
	}
	var ret = make([][]string, 0)
	err = json.Unmarshal(resp.Data, &ret)
	return ret, err
}
