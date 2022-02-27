package yunxin

type FriendAddParam struct {
	Accid    string        `schema:"accid,required"`
	Faccid   string        `schema:"faccid,required"`
	Type     FriendAddType `schema:"type,required"`
	Msg      string        `schema:"msg"`
	ServerEx string        `schema:"serverex"`
}

type FriendAddType int

// 1直接加好友，2请求加好友，3同意加好友，4拒绝加好友
const (
	FriendAddDirect  FriendAddType = 1
	FriendAddRequest FriendAddType = 2
	FriendAddApprove FriendAddType = 3
	FriendAddReject  FriendAddType = 4
)

func (p FriendAddParam) GetPath() string {
	return PathFriendAdd
}

type FriendAddResponse struct {
	BasicResponese
}

type FriendUpdateParam struct {
	Accid    string `schema:"accid,required"`
	Faccid   string `schema:"faccid,required"`
	Alias    string `schema:"alias"`
	Ex       string `schema:"ex"`
	ServerEx string `schema:"serverex"`
}

func (FriendUpdateParam) GetPath() string {
	return PathFriendUpdate
}

type FriendUpdateResponse struct {
	BasicResponese
}

type FriendDeleteParam struct {
	Accid         string `schema:"accid,required"`
	Faccid        string `schema:"faccid,required"`
	IsDeleteAlias bool   `schema:"isDeleteAlias"`
}

func (FriendDeleteParam) GetPath() string {
	return PathFriendDelete
}

type FriendDeleteResponse struct {
	BasicResponese
}

type FriendGetParam struct {
	Accid      string `schema:"accid,required"`
	Updatetime int64  `schema:"updatetime,required"` // 毫秒
}

func (FriendGetParam) GetPath() string {
	return PathFriendGet
}

type FriendGetResponse struct {
	BasicResponese
	Friends []*FriendInfo `json:"friends"`
}

type FriendInfo struct {
	Faccid      string `json:"faccid"`
	CreateTime  int64  `json:"createtime"`
	Bidirection bool   `json:"bidirection"`
	Alias       string `json:"alias"`
}

func (im *IM) FriendAdd(param FriendAddParam) (*FriendAddResponse, error) {
	var result FriendAddResponse
	err := im.PostFormAs(&param, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (im *IM) FriendUpdate(param FriendUpdateParam) (*FriendUpdateResponse, error) {
	var result FriendUpdateResponse
	err := im.PostFormAs(&param, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (im *IM) FriendDelete(param FriendDeleteParam) (*FriendDeleteResponse, error) {
	var result FriendDeleteResponse
	err := im.PostFormAs(&param, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (im *IM) FriendGet(param FriendGetParam) (*FriendGetResponse, error) {
	var result FriendGetResponse
	err := im.PostFormAs(&param, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
