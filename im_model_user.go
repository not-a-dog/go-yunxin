package yunxin

const (
	PathUserCreate       = `/user/create.action`
	PathUserUpdate       = `/user/update.action`
	PathUserRefreshToken = `/user/refreshToken.action`
	PathBlockUser        = `/user/block.action`
	PathUnBlockUser      = `/user/unblock.action`
)

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
	BasicResponese
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

type UpdateUserResponse struct {
	BasicResponese
}

type RefreshTokenParam struct {
	Accid string `schema:"accid,required"`
}

func (RefreshTokenParam) GetPath() string {
	return PathUserRefreshToken
}

type RefreshTokenResponse struct {
	BasicResponese
	Info *RefreshTokenInfo `json:"info"`
}

type RefreshTokenInfo struct {
	Accid string `json:"accid"`
	Token string `json:"token"`
}

type BlockUserParam struct {
	Accid         string `schema:"accid,required"`
	Needkick      bool   `schema:"needkick"`
	KickNotifyExt string `schema:"kickNotifyExt"`
}

func (BlockUserParam) GetPath() string {
	return PathBlockUser
}

type BlockUserResponse struct {
	BasicResponese
}

type UnBlockUserParam struct {
	Accid string `schema:"accid,required"`
}

func (UnBlockUserParam) GetPath() string {
	return PathUnBlockUser
}

type UnBlockUserResponse struct {
	BasicResponese
}
