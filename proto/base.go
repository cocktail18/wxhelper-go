package proto

import "strings"

type MsgType int

const (

	/**
	 *
	 */
	MsgTypeChat                      MsgType = 1     // 私聊消息
	MsgTypeFriendRequest             MsgType = 37    // 好友邀请
	MsgTypeGetCard                   MsgType = 42    // 收到名片
	MsgTypeExpression                MsgType = 47    // 表情
	MsgTypeTransfer                  MsgType = 49    // 转账，引用消息
	MsgTypeAfterTransferOrFileHelper MsgType = 51    // 收到转账或者文件助手消息等
	MsgTypeJoinGroup                 MsgType = 10000 // 进群消息 或者 添加好友 后的打招呼
	MsgTypeScan                      MsgType = 10002 // 扫码触发,会触发2次, 有一次有编号,一次没有,还有登陆之后也有,很多情况都会调用这个

)

type WxUserInfo struct {
	Account         string `json:"account"`         //账号
	HeadImage       string `json:"headImage"`       //头像
	City            string `json:"city"`            //城市
	Country         string `json:"country"`         //国家
	CurrentDataPath string `json:"currentDataPath"` //当前数据目录,登录的账号目录
	DataSavePath    string `json:"dataSavePath"`    //微信保存目录
	Mobile          string `json:"mobile"`          //手机
	Name            string `json:"name"`            //昵称
	Province        string `json:"province"`        //省
	Wxid            string `json:"wxid"`            //wxid
	Signature       string `json:"signature"`       //个人签名
	DbKey           string `json:"dbKey"`           //	数据库的SQLCipher的加密key，可以使用该key配合decrypt.py解密数据库
}

type ContactInfo struct {
	CustomAccount string `json:"customAccount"` //自定义账号
	EncryptName   string `json:"encryptName"`   //昵称
	Nickname      string `json:"nickname"`      //昵称
	Pinyin        string `json:"pinyin"`        //简拼
	PinyinAll     string `json:"pinyinAll"`     //全拼
	Reserved1     int64  `json:"reserved1"`     //未知
	Reserved2     int64  `json:"reserved2"`     //未知
	Type          int64  `json:"type"`          //未知
	VerifyFlag    int64  `json:"verifyFlag"`    //未知
	Wxid          string `json:"wxid"`          //wxid
}

type ContactProfile struct {
	Account   string `json:"account"`
	HeadImage string `json:"headImage"`
	Nickname  string `json:"nickname"`
	V3        string `json:"v3"`
	Wxid      string `json:"wxid"`
}

type SendAppletMsgReq struct {
	Wxid       string `json:"wxid"`       //	接收人wxid
	WaidConcat string `json:"waidConcat"` //	app的wxid与回调信息之类绑定的拼接字符串，伪造的数据可以随意
	AppletWxid string `json:"appletWxid"` //	app的wxid
	JsonParam  string `json:"jsonParam"`  //	相关参数
	HeadImgUrl string `json:"headImgUrl"` //	头像url
	MainImg    string `json:"mainImg"`    //	主图的本地路径,需要在小程序的临时目录下
	IndexPage  string `json:"indexPage"`  //	小程序的跳转页面
}

type PrivateMsgV2 struct {
	Content            string  `json:"content"`
	CreateTime         int     `json:"createTime"`
	DisplayFullContent string  `json:"displayFullContent"`
	FromUser           string  `json:"fromUser"`
	MsgId              int64   `json:"msgId"`
	MsgSequence        int     `json:"msgSequence"`
	Pid                int     `json:"pid"`
	Signature          string  `json:"signature"`
	ToUser             string  `json:"toUser"`
	Type               MsgType `json:"type"`
}

type PrivateMsgV1 struct {
	Content   string  `json:"content"`
	FromGroup string  `json:"fromGroup"`
	FromUser  string  `json:"fromUser"`
	IsSendMsg int     `json:"isSendMsg"`
	MsgId     int64   `json:"msgId"`
	Pid       int     `json:"pid"`
	Sign      string  `json:"sign"`
	Signature string  `json:"signature"`
	Time      string  `json:"time"`
	Timestamp int     `json:"timestamp"`
	Type      MsgType `json:"type"`
}

type WxPrivateMsg struct {
	Content            string  `json:"content"`
	FromGroup          string  `json:"fromGroup"`
	FromUser           string  `json:"fromUser"`
	ToUser             string  `json:"toUser"`
	IsSendMsg          int     `json:"isSendMsg"`
	MsgId              int64   `json:"msgId"`
	Pid                int     `json:"pid"`
	Sign               string  `json:"sign"`
	Signature          string  `json:"signature"`
	Time               string  `json:"time"`
	Timestamp          int     `json:"timestamp"`
	Type               MsgType `json:"type"`
	DisplayFullContent string  `json:"displayFullContent"`
	GroupMemberCount   int     `json:"groupMemberCount"`

	AtWxIds []string `json:"atWxIds"`
}

func (wxPrivateMsg WxPrivateMsg) IsFromGroup() bool {
	return wxPrivateMsg.FromGroup != "" && (wxPrivateMsg.FromGroup != wxPrivateMsg.FromUser || strings.Contains(wxPrivateMsg.FromGroup, "@"))
}
