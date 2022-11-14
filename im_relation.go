package yunxin

import "context"

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

func (FriendAddParam) GetPath() string {
	return PathFriendAdd
}

type FriendAddResponse BasicResponse

type FriendUpdateParam struct {
	AccID    string `schema:"accid,required"`
	FAccID   string `schema:"faccid,required"`
	Alias    string `schema:"alias"`
	Ex       string `schema:"ex"`
	ServerEx string `schema:"serverex"`
}

func (FriendUpdateParam) GetPath() string {
	return PathFriendUpdate
}

type FriendUpdateResponse BasicResponse

type FriendDeleteParam struct {
	AccID         string `schema:"accid,required"`
	FAccID        string `schema:"faccid,required"`
	IsDeleteAlias bool   `schema:"isDeleteAlias"`
}

func (FriendDeleteParam) GetPath() string {
	return PathFriendDelete
}

type FriendDeleteResponse BasicResponse

type FriendGetParam struct {
	AccID      string `schema:"accid,required"`
	UpdateTime int64  `schema:"updatetime,required"` // 毫秒
}

func (FriendGetParam) GetPath() string {
	return PathFriendGet
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

func (c *Client) FriendAdd(ctx context.Context, param *FriendAddParam) (*FriendAddResponse, error) {
	return Request[FriendAddResponse](c, ctx, param)
}

func (c *Client) FriendUpdate(ctx context.Context, param *FriendUpdateParam) (*FriendUpdateResponse, error) {
	return Request[FriendUpdateResponse](c, ctx, param)
}

func (c *Client) FriendDelete(ctx context.Context, param *FriendDeleteParam) (*FriendDeleteResponse, error) {
	return Request[FriendDeleteResponse](c, ctx, param)
}

func (c *Client) FriendGet(ctx context.Context, param *FriendGetParam) (*FriendGetResponse, error) {
	return Request[FriendGetResponse](c, ctx, param)
}

type UserSetSpecialRelationParam struct {
	Accid        string       `schema:"accid,required"`
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

func (UserSetSpecialRelationParam) GetPath() string {
	return PathUserSetSpecialRelation
}

type UserSetSpecialRelationResponse BasicResponse

type UserListBlackAndMuteListParam struct {
	Accid string `schema:"accid,required"`
}

func (UserListBlackAndMuteListParam) GetPath() string {
	return PathUserListBlackAndMuteList
}

type UserListBlackAndMuteListResponse struct {
	BasicResponse
	BlackList []string `json:"blacklist"`
	MuteList  []string `json:"mutelist"`
}

func (c *Client) UserSetSpecialRelation(ctx context.Context, param *UserSetSpecialRelationParam) (*UserSetSpecialRelationResponse, error) {
	return Request[UserSetSpecialRelationResponse](c, ctx, param)
}

func (c *Client) UserListBlackAndMuteList(ctx context.Context, param *UserListBlackAndMuteListParam) (*UserListBlackAndMuteListResponse, error) {
	return Request[UserListBlackAndMuteListResponse](c, ctx, param)
}
