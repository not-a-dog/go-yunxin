package yunxin

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
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
	return &Client{
		appKey:      appKey,
		appSecret:   appSecert,
		FormEncoder: schema.NewEncoder(),
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

func (c *Client) URL(path string) string {
	if c.Host == "" {
		return DefaultHost + path
	}
	return c.Host + path
}

func (c *Client) Post(path string, body io.Reader) (*http.Response, error) {
	r, err := http.NewRequest("POST", c.URL(path), body)
	if err != nil {
		return nil, err
	}
	return c.do(r)
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

func (c *Client) do(r *http.Request) (*http.Response, error) {
	c.AddCheckSum(r)
	ctx := r.Context()
	// add timeout ? TODO: check HTTPClient.Timeout
	if c.Timeout > 0 {
		ctx, cancel := context.WithTimeout(ctx, c.Timeout)
		defer cancel()
		r = r.WithContext(ctx)
	}

	return c.HTTPClient.Do(r)
}

type Response interface {
	SetRawBody(raw []byte)
}

func (c *Client) PostFormAs(path string, value interface{}, outPtr Response) error {
	resp, err := c.PostForm(PathCreateUser, value)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	raw, _ := ioutil.ReadAll(resp.Body)
	outPtr.SetRawBody(raw)
	err = json.Unmarshal(raw, &outPtr)
	if err != nil {
		return err
	}
	return nil
}
