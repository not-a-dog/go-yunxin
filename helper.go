package yunxin

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

func (c *Client) SignACCIDToken(accID string, TTLSecond int) string {
	currentTime := time.Now().UnixMilli()
	text := fmt.Sprintf("%s%s%d%d%s",
		c.appKey,
		accID,
		currentTime,
		TTLSecond,
		c.appSecret,
	)
	h := sha1.New()
	_, _ = h.Write([]byte(text))
	signature := hex.EncodeToString(h.Sum(nil))
	payload := map[string]any{
		"signature": signature,
		"curTime":   currentTime,
		"ttl":       TTLSecond,
	}
	data, _ := json.Marshal(payload)
	return base64.StdEncoding.EncodeToString(data)
}
