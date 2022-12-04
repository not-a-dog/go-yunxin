package yunxin

type FriendAddParam struct {
	AccID    string        `schema:"accid,required"`
	FAccID   string        `schema:"faccid,required"`
	Type     FriendAddType `schema:"type,required"`
	Msg      string        `schema:"msg"`
	ServerEx string        `schema:"serverex"`
}

type FriendAddType int

// 1直接加好友，2请求加好友，3同意加好友，4拒绝加好友
const (
	FriendAddDirect  FriendAddType = 1
	FriendAddRequest FriendAddType = 2
	FriendAddApprove FriendAddType = 3
	FriendAddReject  FriendAddType = 4
)

type FriendAddResponse struct {
	BasicResponse
}

type FriendUpdateParam struct {
	AccID    string `schema:"accid,required"`
	FAccID   string `schema:"faccid,required"`
	Alias    string `schema:"alias"`
	Ex       string `schema:"ex"`
	ServerEx string `schema:"serverex"`
}

type FriendUpdateResponse struct {
	BasicResponse
}

type FriendDeleteParam struct {
	AccID         string `schema:"accid,required"`
	FAccID        string `schema:"faccid,required"`
	IsDeleteAlias bool   `schema:"isDeleteAlias"`
}

type FriendDeleteResponse struct {
	BasicResponse
}

type FriendGetParam struct {
	AccID      string `schema:"accid,required"`
	UpdateTime int64  `schema:"updatetime,required"` // 毫秒
}

type FriendGetResponse struct {
	BasicResponse
	Friends []*FriendInfo `json:"friends"`
}

type FriendInfo struct {
	FAccID      string `json:"faccid"`
	CreateTime  int64  `json:"createtime"`
	BiDirection bool   `json:"bidirection"`
	Alias       string `json:"alias"`
}

type UserSetSpecialRelationParam struct {
	AccID        string       `schema:"accid,required"`
	TargetAcc    string       `schema:"targetAcc,required"`
	RelationType RelationType `schema:"relationType,required"`
	Value        int          `schema:"value,required"` // 0:取消，1:加入
}

type RelationType int

// 1:黑名单操作，2:静音列表操作
const (
	RelationTypeBlack RelationType = 1
	RelationTypeMute  RelationType = 2
)

type UserSetSpecialRelationResponse struct {
	BasicResponse
}

type UserListBlackAndMuteListParam struct {
	AccID string `schema:"accid,required"`
}

type UserListBlackAndMuteListResponse struct {
	BasicResponse
	BlackList []string `json:"blacklist"`
	MuteList  []string `json:"mutelist"`
}
