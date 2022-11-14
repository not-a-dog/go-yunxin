package yunxin

const DefaultHost = "https://api.netease.im/nimserver"

const (
	HTTPAppKey   = "AppKey"
	HTTPNonce    = "Nonce"
	HTTPCurTime  = "CurTime"
	HTTPCheckSum = "CheckSum"

	HTTPContentType     = "Content-Type"
	HTTPXFormURLEncoded = "application/x-www-form-urlencoded;charset=utf-8"
)

const (
	PathUserCreate       = `/user/create.action`
	PathUserUpdate       = `/user/update.action`
	PathUserRefreshToken = `/user/refreshToken.action`
	PathUserBlock        = `/user/block.action`
	PathUserUnBlock      = `/user/unblock.action`
	PathUpdateUserInfo   = `/user/updateUinfo.action`
	PathGetUserInfos     = `/user/getUinfos.action`
	PathUserSetDonnop    = `/user/setDonnop.action`
	PathUserMute         = `/user/mute.action`

	PathFriendAdd    = `/friend/add.action`
	PathFriendUpdate = `/friend/update.action`
	PathFriendDelete = `/friend/delete.action`
	PathFriendGet    = `/friend/get.action`

	PathUserSetSpecialRelation   = `/user/setSpecialRelation.action`
	PathUserListBlackAndMuteList = `/user/listBlackAndMuteList.action`

	PathMsgSendMsg = `/msg/sendMsg.action`

	PathChatRoomCreate = `/chatroom/create.action`
)
