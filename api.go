package yunxin

import "context"

//go:generate go run gen/main.go

func (ChatRoomCreateParam) GetPath() string {
	return PathChatRoomCreate
}

func (c *Client) ChatRoomCreate(ctx context.Context, param *ChatRoomCreateParam) (*ChatRoomCreateResponse, error) {
	return Request[ChatRoomCreateResponse](c, ctx, param)
}

func (FriendAddParam) GetPath() string {
	return PathFriendAdd
}

func (c *Client) FriendAdd(ctx context.Context, param *FriendAddParam) (*FriendAddResponse, error) {
	return Request[FriendAddResponse](c, ctx, param)
}

func (FriendDeleteParam) GetPath() string {
	return PathFriendDelete
}

func (c *Client) FriendDelete(ctx context.Context, param *FriendDeleteParam) (*FriendDeleteResponse, error) {
	return Request[FriendDeleteResponse](c, ctx, param)
}

func (FriendGetParam) GetPath() string {
	return PathFriendGet
}

func (c *Client) FriendGet(ctx context.Context, param *FriendGetParam) (*FriendGetResponse, error) {
	return Request[FriendGetResponse](c, ctx, param)
}

func (FriendUpdateParam) GetPath() string {
	return PathFriendUpdate
}

func (c *Client) FriendUpdate(ctx context.Context, param *FriendUpdateParam) (*FriendUpdateResponse, error) {
	return Request[FriendUpdateResponse](c, ctx, param)
}

func (GetUserInfosParam) GetPath() string {
	return PathGetUserInfos
}

func (c *Client) GetUserInfos(ctx context.Context, param *GetUserInfosParam) (*GetUserInfosResponse, error) {
	return Request[GetUserInfosResponse](c, ctx, param)
}

func (UpdateUserInfoParam) GetPath() string {
	return PathUpdateUserInfo
}

func (c *Client) UpdateUserInfo(ctx context.Context, param *UpdateUserInfoParam) (*UpdateUserInfoResponse, error) {
	return Request[UpdateUserInfoResponse](c, ctx, param)
}

func (UserBlockParam) GetPath() string {
	return PathUserBlock
}

func (c *Client) UserBlock(ctx context.Context, param *UserBlockParam) (*UserBlockResponse, error) {
	return Request[UserBlockResponse](c, ctx, param)
}

func (UserCreateParam) GetPath() string {
	return PathUserCreate
}

func (c *Client) UserCreate(ctx context.Context, param *UserCreateParam) (*UserCreateResponse, error) {
	return Request[UserCreateResponse](c, ctx, param)
}

func (UserListBlackAndMuteListParam) GetPath() string {
	return PathUserListBlackAndMuteList
}

func (c *Client) UserListBlackAndMuteList(ctx context.Context, param *UserListBlackAndMuteListParam) (*UserListBlackAndMuteListResponse, error) {
	return Request[UserListBlackAndMuteListResponse](c, ctx, param)
}

func (UserMuteParam) GetPath() string {
	return PathUserMute
}

func (c *Client) UserMute(ctx context.Context, param *UserMuteParam) (*UserMuteResponse, error) {
	return Request[UserMuteResponse](c, ctx, param)
}

func (UserRefreshTokenParam) GetPath() string {
	return PathUserRefreshToken
}

func (c *Client) UserRefreshToken(ctx context.Context, param *UserRefreshTokenParam) (*UserRefreshTokenResponse, error) {
	return Request[UserRefreshTokenResponse](c, ctx, param)
}

func (UserSetDonnopParam) GetPath() string {
	return PathUserSetDonnop
}

func (c *Client) UserSetDonnop(ctx context.Context, param *UserSetDonnopParam) (*UserSetDonnopResponse, error) {
	return Request[UserSetDonnopResponse](c, ctx, param)
}

func (UserSetSpecialRelationParam) GetPath() string {
	return PathUserSetSpecialRelation
}

func (c *Client) UserSetSpecialRelation(ctx context.Context, param *UserSetSpecialRelationParam) (*UserSetSpecialRelationResponse, error) {
	return Request[UserSetSpecialRelationResponse](c, ctx, param)
}

func (UserUnBlockParam) GetPath() string {
	return PathUserUnBlock
}

func (c *Client) UserUnBlock(ctx context.Context, param *UserUnBlockParam) (*UserUnBlockResponse, error) {
	return Request[UserUnBlockResponse](c, ctx, param)
}

func (UserUpdateParam) GetPath() string {
	return PathUserUpdate
}

func (c *Client) UserUpdate(ctx context.Context, param *UserUpdateParam) (*UserUpdateResponse, error) {
	return Request[UserUpdateResponse](c, ctx, param)
}