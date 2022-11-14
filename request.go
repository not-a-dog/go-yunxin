package yunxin

import (
	"context"
	"net/http"
)

type WithRawBody interface {
	SetRawBody(raw []byte)
}

type Param interface {
	GetPath() string
}

type RawBodyModel struct {
	RawBody string `json:"-"`
}

func (r *RawBodyModel) SetRawBody(raw []byte) {
	r.RawBody = string(raw)
}

type BasicResponse struct {
	RawBodyModel
	Code int    `json:"code"`
	Desc string `json:"desc,omitempty"` // Error description
}

func (r *BasicResponse) IsSuccess() bool {
	return r.Code == http.StatusOK
}

func Request[T any](client *Client, ctx context.Context, param Param) (*T, error) {
	var r = new(T)
	path := param.GetPath()
	resp, err := client.PostForm(ctx, path, param)
	if err != nil {
		return nil, err
	}
	err = client.JSONResponse(resp, &r)
	return r, nil
}
