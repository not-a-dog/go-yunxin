package yunxin

import (
	"net/http"
	"testing"
)

type PostData struct {
	Name  string `schema:"name"`
	Accid string `schema:"accid"`
}

func assertNil(t *testing.T, value interface{}) {
	t.Helper()
	if value != nil {
		t.Errorf("value should be nil, but got %v", value)
	}
}

func assertNotNil(t *testing.T, value interface{}) {
	t.Helper()
	if value == nil {
		t.Errorf("value should not be nil, but got %v", value)
	}
}

func TestClientForm(t *testing.T) {
	request, err := http.NewRequest("POST", "https://httpbin.org/anything", nil)
	assertNil(t, err)
	assertNotNil(t, request)

	client := NewClient("abc", "def")

	data := &PostData{Name: "yunxin", Accid: "123"}
	err = client.AddFormBody(request, data)
	assertNil(t, err)

	response, err := http.DefaultClient.Do(request)
	assertNil(t, err)
	assertNotNil(t, response)
}
