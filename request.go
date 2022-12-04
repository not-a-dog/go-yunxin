package yunxin

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/schema"
)

var defaultEncoder = schema.NewEncoder()

func RegisterEncoder(value interface{}, encoder func(reflect.Value) string) {
	defaultEncoder.RegisterEncoder(value, encoder)
}

func JSONFormEncode(v reflect.Value) string {
	value := v.Interface()
	data, _ := json.Marshal(value)
	return string(data)
}

func init() {
	RegisterEncoder(StringSlice{}, JSONFormEncode)
}

type StringSlice []string
type MillsSecond int64

func NewMillsSecond(t time.Time) MillsSecond {
	return MillsSecond(t.UnixMilli())
}

func (t MillsSecond) ToTime() time.Time {
	return time.UnixMilli(int64(t))
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

var YunxinError = errors.New("request yunxin failed")

type BasicResponse struct {
	RawBodyModel
	Code int    `json:"code"`
	Desc string `json:"desc,omitempty"` // Error description
}

func (r *BasicResponse) IsSuccess() bool {
	return r.Code == http.StatusOK
}

func (r *BasicResponse) AsError() error {
	if r.Code == http.StatusOK {
		return nil
	}
	return YunxinError
	//return fmt.Errorf("%w code %d desc %s rawBody:%s", YunxinError, r.Code, r.Desc, r.RawBody)
}

func Request[T any](client *Client, ctx context.Context, param Param) (*T, error) {
	var r = new(T)
	path := param.GetPath()
	if client.debug {
		client.logger.Info("request to ", path, " with param ", param)
	}
	resp, err := client.PostForm(ctx, path, param)
	if err != nil {
		if client.debug {
			client.logger.Error("request to ", path, " got err ", err)
		}
		return nil, err
	}
	err = client.JSONResponse(resp, r)
	if client.debug {
		client.logger.Error("request to ", path, " got json resp ", r, " and err ", err)
	}
	return r, err
}
