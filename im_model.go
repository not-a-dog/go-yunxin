package yunxin

type BasicResponese struct {
	RawBody string `json:"-"`
	Code    int    `json:"code"`
}

func (r *BasicResponese) IsSuccess() bool {
	return r.Code == 200
}

func (r *BasicResponese) SetRawBody(raw []byte) {
	r.RawBody = string(raw)
}
