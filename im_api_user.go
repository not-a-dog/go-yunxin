package yunxin

const (
	PathCreateUser = `/user/create.action`
)

func (im *IM) CreateUser(param *CreateUserParam) (*CreateUserResponse, error) {
	var result CreateUserResponse
	err := im.PostFormAs(PathCreateUser, param, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
