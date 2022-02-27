package yunxin

import (
	"os"
	"testing"
)

var testAccID = os.Getenv("YUNXIN_TEST_ACCID")

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
		Accid: testAccID,
		Token: `notoken`,
	})

	assertNil(t, err)
	assertNotNil(t, resp)
	assert(t, resp.IsSuccess(), resp.RawBody, testAccID)
}

func TestIMReferToken(t *testing.T) {
	im := &IM{Client: newTestClient(t)}
	resp, err := im.RefreshToken(&RefreshTokenParam{
		Accid: testAccID,
	})

	assertNil(t, err)
	assertNotNil(t, resp)
	assert(t, resp.IsSuccess(), resp.RawBody, testAccID)
}

func TestBlockAndUnBlockUser(t *testing.T) {
	im := &IM{Client: newTestClient(t)}
	resp, err := im.BlockUser(&BlockUserParam{Accid: testAccID})
	assertNil(t, err)
	assertNotNil(t, resp)
	assert(t, resp.IsSuccess(), resp.RawBody, testAccID)

	resp2, err := im.UnBlockUser(&UnBlockUserParam{Accid: testAccID})
	assertNil(t, err)
	assertNotNil(t, resp2)
	assert(t, resp2.IsSuccess(), resp2.RawBody, testAccID)
}
