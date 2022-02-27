package yunxin

import (
	"os"
	"testing"
)

func newTestClient(t *testing.T) *Client {
	t.Helper()
	key := os.Getenv(`YUNXIN_APP_KEY`)
	secert := os.Getenv(`YUNXIN_APP_SECRET`)
	if key == `` || secert == `` {
		t.Error(`YUNXIN_APP_KEY or YUNXIN_APP_SECRET is not set`)
	}
	return NewClient(key, secert, &Configure{})
}

func TestYuxinUpdateUser(t *testing.T) {
	im := &IM{Client: newTestClient(t)}
	resp, err := im.UpdateUser(&UpdateUserParam{
		Accid: os.Getenv(`YUNXIN_TEST_ACCID`),
		Token: `notoken`,
	})

	assertNil(t, err)
	assertNotNil(t, resp)
	assert(t, resp.IsSuccess(), resp.RawBody, os.Getenv(`YUNXIN_TEST_ACCID`))
}

func TestIMReferToken(t *testing.T) {
	im := &IM{Client: newTestClient(t)}
	resp, err := im.RefreshToken(&RefreshTokenParam{
		Accid: os.Getenv(`YUNXIN_TEST_ACCID`),
	})

	assertNil(t, err)
	assertNotNil(t, resp)
	assert(t, resp.IsSuccess(), resp.RawBody, os.Getenv(`YUNXIN_TEST_ACCID`))
}
