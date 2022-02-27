package yunxin

const (
	PathUserCreate = `/user/create.action`
	PathUserUpdate = `/user/update.action`
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
	Info *CreateUserInfo
}

type CreateUserInfo struct {
	Accid string `json:"accid"`
	Token string `json:"token"`
	Name  string `json:"name"`
}

type UpdateUserParam struct {
	Accid string `json:"accid"`
	Token string `json:"token"`
}

func (p UpdateUserParam) GetPath() string {
	return PathUserUpdate
}

type UpdateUserResponse struct {
	BasicResponese
}
