package yunxin

import "encoding/json"

const (
	CreateUserURL = `/user/create.action`
)

func (im *IM) CreateUser(param *CreateUserParam) (*CreateUserResponse, error) {
	resp, err := im.PostForm(CreateUserURL, param)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var result CreateUserResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
