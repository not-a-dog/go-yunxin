package yunxin

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gorilla/schema"
)

var defaultEncoder = schema.NewEncoder()

func init() {
	defaultEncoder.SetAliasTag(`json`)
}

func RegisterEncoder(value interface{}, encoder func(reflect.Value) string) {
	defaultEncoder.RegisterEncoder(value, encoder)
}

func JSONFormEncode(v reflect.Value) string {
	value := v.Interface()
	data, _ := json.Marshal(value)
	return string(data)
}

type StringSlice []string

func init() {
	RegisterEncoder(StringSlice{}, JSONFormEncode)
}

type Param interface {
	GetPath() string
}

type WithRawBody interface {
	SetRawBody(raw []byte)
}

type RawBodyModel struct {
	RawBody string `json:"-"`
}

func (r *RawBodyModel) SetRawBody(raw []byte) {
	r.RawBody = string(raw)
}

var _ WithRawBody = new(RawBodyModel)

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
	return r, err
}
