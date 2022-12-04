package yunxin

const (
	PathUserCreate         = `/user/create.action`
	PathUserUpdate         = `/user/update.action`
	PathUserRefreshToken   = `/user/refreshToken.action`
	PathUserBlock          = `/user/block.action`
	PathUserUnBlock        = `/user/unblock.action`
	PathUserUpdateUserInfo = `/user/updateUinfo.action`
	PathGetUserInfos       = `/user/getUinfos.action`
	PathUserSetDonnop      = `/user/setDonnop.action`
	PathUserMute           = `/user/mute.action`

	PathFriendAdd    = `/friend/add.action`
	PathFriendUpdate = `/friend/update.action`
	PathFriendDelete = `/friend/delete.action`
	PathFriendGet    = `/friend/get.action`

	PathUserSetSpecialRelation   = `/user/setSpecialRelation.action`
	PathUserListBlackAndMuteList = `/user/listBlackAndMuteList.action`

	PathMsgSendMsg = `/msg/sendMsg.action`
	PathMsgDelMsg  = `/msg/delMsg.action`

	PathChatRoomCreate                  = `/chatroom/create.action`
	PathChatroomGet                     = `/chatroom/get.action`
	PathChatroomGetBatch                = `/chatroom/getBatch.action`
	PathChatRoomRequestAddr             = `/chatroom/requestAddr.action`
	PathChatRoomSendMsg                 = `/chatroom/sendMsg.action`
	PathChatRoomMembersByPage           = `/chatroom/membersByPage.action`
	PathChatRoomSetMemberRole           = `/chatroom/setMemberRole.action`
	PathChatRoomRecall                  = `/chatroom/recall.action`
	PathChatRoomUpdateInOutNotification = `/chatroom/updateInOutNotification.action`
)
