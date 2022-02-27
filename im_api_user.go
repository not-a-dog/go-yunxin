package yunxin

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
