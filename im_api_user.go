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

func (im *IM) RefreshToken(param *RefreshTokenParam) (*RefreshTokenResponse, error) {
	var result RefreshTokenResponse
	err := im.PostFormAs(param, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (im *IM) BlockUser(param *BlockUserParam) (*BlockUserResponse, error) {
	var result BlockUserResponse
	err := im.PostFormAs(param, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (im *IM) UnBlockUser(param *UnBlockUserParam) (*UnBlockUserResponse, error) {
	var result UnBlockUserResponse
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
