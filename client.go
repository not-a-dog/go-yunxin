package yunxin

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
	HTTPClient  *http.Client
	*Configure
}

func NewClient(appKey, appSecert string, cfg *Configure) *Client {
	encoder := schema.NewEncoder()
	LoadCustomTypes(encoder)
	return &Client{
		appKey:      appKey,
		appSecret:   appSecert,
		FormEncoder: encoder,
		Configure:   cfg,
		HTTPClient:  http.DefaultClient,
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

func (c *Client) URL(path string) string {
	if c.Host == "" {
		return DefaultHost + path
	}
	return c.Host + path
}

func (c *Client) do(r *http.Request) (*http.Response, error) {
	c.AddCheckSum(r)
	ctx := r.Context()
	// add timeout ? TODO: use HTTPClient.Timeout ?
	if c.Timeout > 0 {
		ctx, cancel := context.WithTimeout(ctx, c.Timeout)
		defer cancel()
		r = r.WithContext(ctx)
	}

	return c.HTTPClient.Do(r)
}

func (c *Client) PostForm(path string, value interface{}) (*http.Response, error) {
	r, err := http.NewRequest("POST", c.URL(path), nil)
	if err != nil {
		return nil, err
	}
	err = c.AddFormBody(r, value)
	if err != nil {
		return nil, err
	}
	return c.do(r)
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

func (c *Client) DecodeResponse(resp *http.Response, outPtr Response) error {
	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	outPtr.SetRawBody(raw)
	return json.Unmarshal(raw, &outPtr)
}

type Param interface {
	GetPath() string
}

type Response interface {
	SetRawBody(raw []byte)
}

func (c *Client) PostFormAs(param Param, outPtr Response) error {
	resp, err := c.PostForm(param.GetPath(), param)
	if err != nil {
		return err
	}
	return c.DecodeResponse(resp, outPtr)
}

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

type BasicResponese struct {
	RawBody string `json:"-"`
	Code    int    `json:"code"`
	Desc    string `json:"desc,omitempty"`
}

func (r *BasicResponese) IsSuccess() bool {
	return r.Code == 200
}

func (r *BasicResponese) SetRawBody(raw []byte) {
	r.RawBody = string(raw)
}
