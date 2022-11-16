package yunxin

type FriendAddParam struct {
	Accid    string        `schema:"accid,required"`
	FAccid   string        `schema:"faccid,required"`
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

type FriendAddResponse BasicResponse

type FriendUpdateParam struct {
	Accid     string `schema:"accid,required"`
	FAccid    string `schema:"faccid,required"`
	Alias    string `schema:"alias"`
	Ex       string `schema:"ex"`
	ServerEx string `schema:"serverex"`
}

type FriendUpdateResponse BasicResponse

type FriendDeleteParam struct {
	Accid          string `schema:"accid,required"`
	FAccid         string `schema:"faccid,required"`
	IsDeleteAlias bool   `schema:"isDeleteAlias"`
}

type FriendDeleteResponse BasicResponse

type FriendGetParam struct {
	Accid       string `schema:"accid,required"`
	UpdateTime int64  `schema:"updatetime,required"` // 毫秒
}

type FriendGetResponse struct {
	BasicResponse
	Friends []*FriendInfo `json:"friends"`
}

type FriendInfo struct {
	FAccid       string `json:"faccid"`
	CreateTime  int64  `json:"createtime"`
	BiDirection bool   `json:"bidirection"`
	Alias       string `json:"alias"`
}

type UserSetSpecialRelationParam struct {
	Accid         string       `schema:"accid,required"`
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

type UserSetSpecialRelationResponse BasicResponse

type UserListBlackAndMuteListParam struct {
	Accid string `schema:"accid,required"`
}

type UserListBlackAndMuteListResponse struct {
	BasicResponse
	BlackList []string `json:"blacklist"`
	MuteList  []string `json:"mutelist"`
}
