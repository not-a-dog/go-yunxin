package yunxin

import (
	"net/url"
	"os"
	"testing"

	"github.com/gorilla/schema"
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
	resp, err := im.UserBlock(&UserBlockParam{Accid: testAccID})
	assertNil(t, err)
	assertNotNil(t, resp)
	assert(t, resp.IsSuccess(), resp.RawBody, testAccID)

	resp2, err := im.UserUnBlock(&UserUnBlockParam{Accid: testAccID})
	assertNil(t, err)
	assertNotNil(t, resp2)
	assert(t, resp2.IsSuccess(), resp2.RawBody, testAccID)
}

func TestUpdateUserInfo(t *testing.T) {
	im := &IM{Client: newTestClient(t)}
	resp, err := im.UpdateUserInfo(&UpdateUserInfoParam{
		Accid: testAccID,
		Name:  `test`,
	})
	assertNil(t, err)
	assertNotNil(t, resp)
	assert(t, resp.IsSuccess(), resp.RawBody, testAccID)
}

func TestUserInfosParam(t *testing.T) {
	encoder := schema.NewEncoder()
	LoadCustomTypes(encoder)
	param := &GetUserInfosParam{
		Accids: StringSlice{"A", "B", "C"},
	}

	form := url.Values{}
	err := encoder.Encode(param, form)
	assertNil(t, err)
	assert(t, form.Get("accids") == `["A","B","C"]`, form["accids"])
}

func TestGetUserInfos(t *testing.T) {
	im := &IM{Client: newTestClient(t)}
	resp, err := im.GetUserInfos(&GetUserInfosParam{
		Accids: []string{testAccID},
	})
	assertNil(t, err)
	assertNotNil(t, resp)
	assert(t, resp.IsSuccess(), resp.RawBody, testAccID)
	assert(t, len(resp.Uinfos) == 1, resp.RawBody)
}

func TestUserSetDonnop(t *testing.T) {
	im := &IM{Client: newTestClient(t)}
	resp, err := im.UserSetDonnop(&UserSetDonnopParam{
		Accid:      testAccID,
		DonnopOpen: false,
	})
	assertNil(t, err)
	assertNotNil(t, resp)
	// assert(t, resp.IsSuccess(), resp.RawBody, testAccID)
}

func TestUserMute(t *testing.T) {
	im := &IM{Client: newTestClient(t)}
	resp, err := im.UserMute(&UserMuteParam{
		Accid: testAccID,
		Mute:  false,
	})
	assertNil(t, err)
	assertNotNil(t, resp)
	assert(t, resp.IsSuccess(), resp.RawBody, testAccID)
}
