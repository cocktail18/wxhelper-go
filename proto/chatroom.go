package proto

type ChatroomDetail struct {
	ChatRoomId string `json:"chatRoomId"` //	群id
	Notice     string `json:"notice"`     //	公告通知
	Admin      string `json:"admin"`      //	群管理
	Xml        string `json:"xml"`        //	xml信息
}

type ChatroomMember struct {
	ChatRoomId     string `json:"chatRoomId"`     //	群id
	Members        string `json:"members"`        //	成员id
	MemberNickname string `json:"memberNickname"` //	成员昵称
	Admin          string `json:"admin"`          //	群管理
	AdminNickname  string `json:"adminNickname"`  //	管理昵称
}
