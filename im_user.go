package yunxin

type UserCreateParam struct {
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

type UserCreateResponse struct {
	BasicResponse
	Info *CreateUserInfo `json:"info"`
}

type CreateUserInfo struct {
	Accid string `json:"accid"`
	Token string `json:"token"`
	Name  string `json:"name"`
}

type UserUpdateParam struct {
	Accid string `schema:"accid,required"`
	Token string `schema:"token"`
}

type UserUpdateResponse BasicResponse

type UserRefreshTokenParam struct {
	Accid string `schema:"accid,required"`
}

type UserRefreshTokenResponse struct {
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

type UserBlockResponse BasicResponse

type UserUnBlockParam struct {
	Accid string `schema:"accid,required"`
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

type UpdateUserInfoResponse BasicResponse

type GetUserInfosParam struct {
	Accids StringSlice `schema:"accids,required"`
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

type UserSetDonnopResponse BasicResponse

type UserMuteParam struct {
	Accid string `schema:"accid,required"`
	Mute  bool   `schema:"mute,required"`
}

type UserMuteResponse BasicResponse
