package yunxin

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/schema"
)

type Client struct {
	appKey      string
	appSecret   string
	FormEncoder *schema.Encoder
}

func NewClient(appKey, appSecert string) *Client {
	return &Client{
		appKey:      appKey,
		appSecret:   appSecert,
		FormEncoder: schema.NewEncoder(),
	}
}

func (c *Client) AddCheckSum(r *http.Request) {
	curTime := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := strconv.FormatInt(rand.Int63(), 10)

	h := sha1.New()
	_, _ = h.Write([]byte(c.appSecret + nonce + curTime))
	checkSum := hex.EncodeToString(h.Sum(nil))

	r.Header.Set("AppKey", c.appKey)
	r.Header.Set("Nonce", nonce)
	r.Header.Set("CurTime", curTime)
	r.Header.Set("CheckSum", checkSum)
}

func (c *Client) AddFormBody(r *http.Request, v interface{}) error {
	form := url.Values{}
	err := c.FormEncoder.Encode(v, form)
	if err != nil {
		return err
	}
	r.Body = io.NopCloser(strings.NewReader(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	return nil
}
