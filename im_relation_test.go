package yunxin

import (
	"testing"
	"time"
)

func TestFriendGet(t *testing.T) {
	im := &IM{Client: newTestClient(t)}
	resp, err := im.FriendGet(&FriendGetParam{
		Accid:      testAccID,
		Updatetime: time.Now().UnixMilli(),
	})
	assertNil(t, err)
	assertNotNil(t, resp)
	assert(t, resp.IsSuccess(), resp.RawBody, testAccID)
}
