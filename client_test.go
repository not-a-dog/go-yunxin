package yunxin

import (
	"net/http"
	"testing"
)

type PostData struct {
	Name  string `schema:"name"`
	Accid string `schema:"accid"`
}

func assert(t *testing.T, ok bool, msg ...interface{}) {
	t.Helper()
	if !ok {
		t.Error(`assertion failed`, msg)
	}
}

func assertNil(t *testing.T, value interface{}) {
	t.Helper()
	assert(t, value == nil, "value should be nil, but got ", value)
}

func assertNotNil(t *testing.T, value interface{}) {
	t.Helper()
	assert(t, value != nil, "value should not be nil ", value)
}

func TestClientForm(t *testing.T) {
	request, err := http.NewRequest("POST", "https://httpbin.org/anything", nil)
	assertNil(t, err)
	assertNotNil(t, request)

	client := NewClient("abc", "def", &Configure{})

	data := &PostData{Name: "yunxin", Accid: "123"}
	err = client.AddFormBody(request, data)
	assertNil(t, err)

	response, err := http.DefaultClient.Do(request)
	assertNil(t, err)
	assertNotNil(t, response)

	var r BasicResponese
	err = client.DecodeResponse(response, &r)
	assertNil(t, err)
	assert(t, r.Code == 0, r.RawBody)
}

func TestClientSignToken(t *testing.T) {
	client := NewClient("abc", "def", &Configure{})
	token := client.SignToken("accid", 600)
	assert(t, len(token) == 120, token)
}
