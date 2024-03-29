package yunxin

import (
	"net/url"
	"testing"
)

func TestStringSliceEncoder(t *testing.T) {
	var query = struct {
		Names StringSlice `schema:"names"`
	}{
		Names: StringSlice{"x", "y"},
	}

	form := url.Values{}
	err := defaultEncoder.Encode(query, form)
	if err != nil {
		t.Fatal(err)
	}

	data := form.Encode()
	data, err = url.QueryUnescape(data)
	if err != nil {
		t.Fatal(err)
	}

	if data != `names=["x","y"]` {
		t.Fatal(data)
	}
}

func TestBasicResponse(t *testing.T) {
	var r = BasicResponse{Code: 100, Desc: ""}
	if r.IsSuccess() != false {
		t.Fail()
	}

	err := r.AsError()
	if _, ok := err.(*YunxinError); !ok {
		t.Error(err)
	}
}

type unkonwnResponse struct {
	BasicResponse
}

func TestUnkonwnResponse(t *testing.T) {
	var r unkonwnResponse
	err := r.AsError()
	if _, ok := err.(*YunxinError); !ok {
		t.Error(err)
	}
}
