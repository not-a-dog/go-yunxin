package yunxin

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

type UserBlockParam struct {
	Accid         string `schema:"accid,required"`
	Needkick      bool   `schema:"needkick"`
	KickNotifyExt string `schema:"kickNotifyExt"`
}

func (UserBlockParam) GetPath() string {
	return PathUserBlock
}

type UserBlockResponse struct {
	BasicResponese
}

type UserUnBlockParam struct {
	Accid string `schema:"accid,required"`
}

func (UserUnBlockParam) GetPath() string {
	return PathUserUnBlock
}

type UserUnBlockResponse struct {
	BasicResponese
}

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

type UpdateUserInfoResponse struct {
	BasicResponese
}

type GetUserInfosParam struct {
	Accids StringSlice `schema:"accids,required"`
}

func (GetUserInfosParam) GetPath() string {
	return PathGetUserInfos
}

type GetUserInfosResponse struct {
	BasicResponese
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

type UserSetDonnopResponse struct {
	BasicResponese
}

type UserMuteParam struct {
	Accid string `schema:"accid,required"`
	Mute  bool   `schema:"mute,required"`
}

func (UserMuteParam) GetPath() string {
	return PathUserMute
}

type UserMuteResponse struct {
	BasicResponese
}

func (im *IM) CreateUser(param *CreateUserParam) (*CreateUserResponse, error) {
	var result CreateUserResponse
	err := im.PostFormAs(param, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (im *IM) UpdateUser(param *UpdateUserParam) (*UpdateUserResponse, error) {
	var result UpdateUserResponse
	err := im.PostFormAs(param, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (im *IM) RefreshToken(param *RefreshTokenParam) (*RefreshTokenResponse, error) {
	var result RefreshTokenResponse
	err := im.PostFormAs(param, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (im *IM) UserBlock(param *UserBlockParam) (*UserBlockResponse, error) {
	var result UserBlockResponse
	err := im.PostFormAs(param, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (im *IM) UserUnBlock(param *UserUnBlockParam) (*UserUnBlockResponse, error) {
	var result UserUnBlockResponse
	err := im.PostFormAs(param, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (im *IM) UpdateUserInfo(param *UpdateUserInfoParam) (*UpdateUserInfoResponse, error) {
	var result UpdateUserInfoResponse
	err := im.PostFormAs(param, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (im *IM) GetUserInfos(param *GetUserInfosParam) (*GetUserInfosResponse, error) {
	var result GetUserInfosResponse
	err := im.PostFormAs(param, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (im *IM) UserSetDonnop(param *UserSetDonnopParam) (*UserSetDonnopResponse, error) {
	var result UserSetDonnopResponse
	err := im.PostFormAs(param, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (im *IM) UserMute(param *UserMuteParam) (*UserMuteResponse, error) {
	var result UserMuteResponse
	err := im.PostFormAs(param, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
