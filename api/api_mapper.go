package api

type ApiUrl int
type ApiVersion int

const (
	ApiVersionV1 ApiVersion = 1
	ApiVersionV2 ApiVersion = 2
)

const (
	CheckLoginUrl = iota
	UserInfoUrl
	SendTextMsgUrl
	SendAtTextUrl
	SendImagesMsgUrl
	SendFileMsgUrl
	SendCustomEmotionUrl
	SendAppletUrl
	SendPatMsgUrl
	SearchFriend
	FriendRequest
	ConfirmFriendRequest
	OcrUrl
	ForwardMsgUrl
	HookSyncMsgUrl
	UnHookSyncMsgUrl
	HookAudioMsgUrl
	UnHookAudioMsgUrl
	GetContactListUrl
	GetContactProfileUrl
	GetContactNicknameUrl

	//转账相关
	ConfirmTransfer

	// db操作
	GetDBInfoUrl
	ExecSqlUrl

	// 公众号相关
	ForwardPublicMsgByMsgIdUrl
	ForwardPublicMsgUrl

	// 群操作
	GetChatRoomDetailInfoUrl
	GetMemberFromChatRoomUrl
	GetNicknameFromChatRoomUrl
	AddMemberToChatRoomUrl
	ModifyNicknameUrl
	DelMemberFromChatRoomUrl
	TopMsgUrl
	RemoveTopMsgUrl
	InviteMemberToChatRoomUrl

	// 朋友圈相关,朋友圈首页,前置条件需先调用hook消息接口成功,具体内容会在hook消息里返回，格式如下：
	/**
	  {
		  "data":[
			  {
				  "content": "",
				  "createTime': 1691125287,
				  "senderId': "",
				  "snsId': 123,
				  "xml':""
			  }
		  ]
	  }
	*/
	GetSNSFirstPageUrl
	GetSNSNextPageUrl

	// 收藏相关
	AddFavFromMsgUrl
	AddFavFromImageUrl

	// 附件相关
	DownloadAttachUrl
	DecodeImageUrl
	GetVoiceByMsgIdUrl
)

var (
	urlMapper = map[ApiVersion]map[ApiUrl]string{
		ApiVersionV1: {
			CheckLoginUrl:         "/api/?type=0",
			UserInfoUrl:           "/api/?type=1",
			SendTextMsgUrl:        "/api/?type=2",
			SendAtTextUrl:         "/api/?type=3",
			SendImagesMsgUrl:      "/api/?type=5",
			SendFileMsgUrl:        "/api/?type=6",
			SendCustomEmotionUrl:  "",
			SendAppletUrl:         "",
			SendPatMsgUrl:         "/api/?type=50", // 拍一拍
			SearchFriend:          "/api/?type=19", //通过手机或qq查找微信
			FriendRequest:         "/api/?type=20", //通过wxid添加好友
			ConfirmFriendRequest:  "/api/?type=23", //通过好友申请
			OcrUrl:                "/api/?type=49",
			ForwardMsgUrl:         "/api/?type=40",
			HookSyncMsgUrl:        "/api/?type=9",
			UnHookSyncMsgUrl:      "/api/?type=10",
			HookAudioMsgUrl:       "/api/?type=13",
			UnHookAudioMsgUrl:     "/api/?type=14",
			GetContactListUrl:     "/api/?type=46",
			GetContactNicknameUrl: "/api/?type=55",

			//转账相关
			ConfirmTransfer: "/api/?type=45", //收到转账消息后，自动收款确认。type=49 即是转账消息

			// db操作
			GetDBInfoUrl: "/api/?type=32",
			ExecSqlUrl:   "/api/?type=34",

			// 公众号相关
			ForwardPublicMsgByMsgIdUrl: "",
			ForwardPublicMsgUrl:        "",

			// 群操作
			GetChatRoomDetailInfoUrl:   "/api/?type=47",
			GetMemberFromChatRoomUrl:   "/api/?type=25",
			GetNicknameFromChatRoomUrl: "/api/?type=26",
			AddMemberToChatRoomUrl:     "/api/?type=28",
			ModifyNicknameUrl:          "/api/?type=31",
			DelMemberFromChatRoomUrl:   "/api/?type=27",
			TopMsgUrl:                  "/api/?type=51",
			RemoveTopMsgUrl:            "/api/?type=52",
			InviteMemberToChatRoomUrl:  "",

			// 朋友圈相关,朋友圈首页,前置条件需先调用hook消息接口成功,具体内容会在hook消息里返回，格式如下：
			/**
			  {
				  "data":[
					  {
						  "content": "",
						  "createTime': 1691125287,
						  "senderId': "",
						  "snsId': 123,
						  "xml':""
					  }
				  ]
			  }
			*/
			GetSNSFirstPageUrl: "/api/?type=53",
			GetSNSNextPageUrl:  "/api/?type=54",

			// 收藏相关
			AddFavFromMsgUrl:   "",
			AddFavFromImageUrl: "",

			// 附件相关
			DownloadAttachUrl:  "/api/?type=56",
			DecodeImageUrl:     "/api/?type=48",
			GetVoiceByMsgIdUrl: "/api/?type=57", //根据消息id，获取该语音消息的语音文件，文件为silk3格式，可以自行转换mp3.
		},
		ApiVersionV2: {
			CheckLoginUrl:        "/api/checkLogin",
			UserInfoUrl:          "/api/userInfo",
			SendTextMsgUrl:       "/api/sendTextMsg",
			SendFileMsgUrl:       "/api/sendFileMsg",
			SendCustomEmotionUrl: "/api/sendCustomEmotion",
			SendAppletUrl:        "/api/sendApplet",
			SendPatMsgUrl:        "/api/sendPatMsg",
			OcrUrl:               "/api/ocr",
			SendImagesMsgUrl:     "/api/sendImagesMsg",
			ForwardMsgUrl:        "/api/forwardMsg",
			SendAtTextUrl:        "/api/sendAtText",
			HookSyncMsgUrl:       "/api/hookSyncMsg",
			UnHookSyncMsgUrl:     "/api/unhookSyncMsg",
			GetContactListUrl:    "/api/getContactList",
			GetDBInfoUrl:         "/api/getDBInfo",
			ExecSqlUrl:           "/api/execSql",
			GetContactProfileUrl: "/api/getContactProfile",

			// 公众号相关
			ForwardPublicMsgByMsgIdUrl: "/api/forwardPublicMsgByMsgId",
			ForwardPublicMsgUrl:        "/api/forwardPublicMsg",

			// 群操作
			GetChatRoomDetailInfoUrl:  "/api/getChatRoomDetailInfo",
			GetMemberFromChatRoomUrl:  "/api/getMemberFromChatRoom",
			AddMemberToChatRoomUrl:    "/api/addMemberToChatRoom",
			ModifyNicknameUrl:         "/api/modifyNickname",
			DelMemberFromChatRoomUrl:  "/api/delMemberFromChatRoom",
			TopMsgUrl:                 "/api/topMsg",
			RemoveTopMsgUrl:           "/api/removeTopMsg",
			InviteMemberToChatRoomUrl: "/api/InviteMemberToChatRoom",

			// 朋友圈相关,朋友圈首页,前置条件需先调用hook消息接口成功,具体内容会在hook消息里返回，格式如下：
			/**
			  {
				  "data":[
					  {
						  "content": "",
						  "createTime': 1691125287,
						  "senderId': "",
						  "snsId': 123,
						  "xml':""
					  }
				  ]
			  }
			*/
			GetSNSFirstPageUrl: "/api/getSNSFirstPage",
			GetSNSNextPageUrl:  "/api/getSNSNextPage",

			// 收藏相关
			AddFavFromMsgUrl:   "/api/addFavFromMsg",
			AddFavFromImageUrl: "/api/addFavFromImage",

			// 附件相关
			DownloadAttachUrl:  "/api/downloadAttach",
			DecodeImageUrl:     "/api/decodeImage",
			GetVoiceByMsgIdUrl: "/api/getVoiceByMsgId",
		},
	}
)

func getUrlMapper(apiVersion ApiVersion) map[ApiUrl]string {
	return urlMapper[apiVersion]
}
