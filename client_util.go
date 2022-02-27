package yunxin

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

func (c *Client) SignToken(accid string, ttl int) string {
	currentTime := time.Now().UnixMilli()
	text := fmt.Sprintf("%s%s%d%d%s",
		c.appKey,
		accid,
		currentTime,
		ttl,
		c.appSecret,
	)
	h := sha1.New()
	_, _ = h.Write([]byte(text))
	signature := hex.EncodeToString(h.Sum(nil))
	payload := map[string]interface{}{
		"signature": signature,
		"curTime":   currentTime,
		"ttl":       ttl,
	}
	data, _ := json.Marshal(payload)
	return base64.StdEncoding.EncodeToString(data)
}
