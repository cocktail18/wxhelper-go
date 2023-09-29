package proto

type DbInfo struct {
	DatabaseName string        `json:"databaseName"` //	数据库名称
	Handle       int64         `json:"handle"`       //	句柄
	Tables       []DbTableInfo `json:"tables"`       //	表信息
}

type DbTableInfo struct {
	Name      string `json:"name"`      //	表名
	RootPage  string `json:"rootpage"`  //	rootpage
	Sql       string `json:"sql"`       //	ddl语句
	TableName string `json:"tableName"` //	表名
}
