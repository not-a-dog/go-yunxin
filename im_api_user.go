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
