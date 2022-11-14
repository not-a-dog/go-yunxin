package yunxin

import "context"

type CreateUserParam struct {
	Accid  string `schema:"accid,required"`
	Name   string `schema:"name"`
	Icon   string `schema:"icon"`
	Token  string `schema:"token"`
	Sign   string `schema:"sign"` // 用户签名
	Email  string `schema:"email"`
	Birth  string `schema:"birth"`
	Mobile string `schema:"mobile"`
	Gender int    `schema:"gender"`
	Ex     string `schema:"ex"`
	Bid    string `schema:"bid"`
}

func (p CreateUserParam) GetPath() string {
	return PathUserCreate
}

type CreateUserResponse struct {
	BasicResponse
	Info *CreateUserInfo `json:"info"`
}

type CreateUserInfo struct {
	Accid string `json:"accid"`
	Token string `json:"token"`
	Name  string `json:"name"`
}

type UpdateUserParam struct {
	Accid string `schema:"accid,required"`
	Token string `schema:"token"`
}

func (p UpdateUserParam) GetPath() string {
	return PathUserUpdate
}

type UpdateUserResponse BasicResponse

type RefreshTokenParam struct {
	Accid string `schema:"accid,required"`
}

func (RefreshTokenParam) GetPath() string {
	return PathUserRefreshToken
}

type RefreshTokenResponse struct {
	BasicResponse
	Info *RefreshTokenInfo `json:"info"`
}

type RefreshTokenInfo struct {
	Accid string `json:"accid"`
	Token string `json:"token"`
}

type UserBlockParam struct {
	Accid         string `schema:"accid,required"`
	Needkick      bool   `schema:"needkick"`
	KickNotifyExt string `schema:"kickNotifyExt"`
}

func (UserBlockParam) GetPath() string {
	return PathUserBlock
}

type UserBlockResponse BasicResponse

type UserUnBlockParam struct {
	Accid string `schema:"accid,required"`
}

func (UserUnBlockParam) GetPath() string {
	return PathUserUnBlock
}

type UserUnBlockResponse BasicResponse

type UpdateUserInfoParam struct {
	Accid  string `schema:"accid,required"`
	Name   string `schema:"name"`
	Icon   string `schema:"icon"`
	Sign   string `schema:"sign"` // 用户签名
	Email  string `schema:"email"`
	Birth  string `schema:"birth"`
	Mobile string `schema:"mobile"`
	Gender int    `schema:"gender"`
	Ex     string `schema:"ex"`
	Bid    string `schema:"bid"` // 反垃圾业务ID，JSON字符串，{"textbid":"","picbid":""}
}

func (UpdateUserInfoParam) GetPath() string {
	return PathUpdateUserInfo
}

type UpdateUserInfoResponse BasicResponse

type GetUserInfosParam struct {
	Accids StringSlice `schema:"accids,required"`
}

func (GetUserInfosParam) GetPath() string {
	return PathGetUserInfos
}

type GetUserInfosResponse struct {
	BasicResponse
	Uinfos []*UserInfo `json:"uinfos"`
}

type UserInfo struct {
	Accid  string `json:"accid"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	Sign   string `json:"sign"` // 用户签名
	Email  string `json:"email"`
	Birth  string `json:"birth"`
	Mobile string `json:"mobile"`
	Ex     string `json:"ex"`
	Gender int    `json:"gender"`
	Valid  bool   `json:"valid"`
	Mute   bool   `json:"mute"`
}

type UserSetDonnopParam struct {
	Accid      string `schema:"accid,required"`
	DonnopOpen bool   `schema:"donnopOpen,required"`
}

func (UserSetDonnopParam) GetPath() string {
	return PathUserSetDonnop
}

type UserSetDonnopResponse BasicResponse

type UserMuteParam struct {
	Accid string `schema:"accid,required"`
	Mute  bool   `schema:"mute,required"`
}

func (UserMuteParam) GetPath() string {
	return PathUserMute
}

type UserMuteResponse BasicResponse

func (c *Client) CreateUser(ctx context.Context, param *CreateUserParam) (*CreateUserResponse, error) {
	return Request[CreateUserResponse](c, ctx, param)
}

func (c *Client) UpdateUser(ctx context.Context, param *UpdateUserParam) (*UpdateUserResponse, error) {
	return Request[UpdateUserResponse](c, ctx, param)
}

func (c *Client) RefreshToken(ctx context.Context, param *RefreshTokenParam) (*RefreshTokenResponse, error) {
	return Request[RefreshTokenResponse](c, ctx, param)
}

func (c *Client) UserBlock(ctx context.Context, param *UserBlockParam) (*UserBlockResponse, error) {
	return Request[UserBlockResponse](c, ctx, param)
}

func (c *Client) UserUnBlock(ctx context.Context, param *UserUnBlockParam) (*UserUnBlockResponse, error) {
	return Request[UserUnBlockResponse](c, ctx, param)
}

func (c *Client) UpdateUserInfo(ctx context.Context, param *UpdateUserInfoParam) (*UpdateUserInfoResponse, error) {
	return Request[UpdateUserInfoResponse](c, ctx, param)
}

func (c *Client) GetUserInfos(ctx context.Context, param *GetUserInfosParam) (*GetUserInfosResponse, error) {
	return Request[GetUserInfosResponse](c, ctx, param)
}

func (c *Client) UserSetDonnop(ctx context.Context, param *UserSetDonnopParam) (*UserSetDonnopResponse, error) {
	return Request[UserSetDonnopResponse](c, ctx, param)
}

func (c *Client) UserMute(ctx context.Context, param *UserMuteParam) (*UserMuteResponse, error) {
	return Request[UserMuteResponse](c, ctx, param)
}
