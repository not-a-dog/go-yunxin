package yunxin

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/schema"
)

const DefaultHost = "https://api.netease.im/nimserver"

const (
	HTTPAppKey   = "AppKey"
	HTTPNonce    = "Nonce"
	HTTPCurTime  = "CurTime"
	HTTPCheckSum = "CheckSum"

	HTTPContentType     = "Content-Type"
	HTTPXFormURLEncoded = "application/x-www-form-urlencoded;charset=utf-8"
)

type Client struct {
	appKey      string
	appSecret   string
	formEncoder *schema.Encoder
	host        string
	http        *http.Client
}

func NewClient(appKey, appSecret string, opts ...Option) *Client {
	c := &Client{
		appKey:      appKey,
		appSecret:   appSecret,
		formEncoder: defaultEncoder,
		http:        http.DefaultClient,
		host:        DefaultHost,
	}

	for _, opt := range opts {
		opt(c)
	}
	return c
}

type Option func(c *Client)

func WithEncoder(encoder *schema.Encoder) Option {
	return func(c *Client) {
		c.formEncoder = encoder
	}
}

func WithHTTPClient(h *http.Client) Option {
	return func(c *Client) {
		c.http = h
	}
}

func WithHost(host string) Option {
	return func(c *Client) {
		c.host = host
	}
}

func (c *Client) AddCheckSum(r *http.Request) {
	curTime := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := strconv.FormatInt(rand.Int63(), 10)

	h := sha1.New()
	_, _ = h.Write([]byte(c.appSecret + nonce + curTime))
	checkSum := hex.EncodeToString(h.Sum(nil))

	r.Header.Set(HTTPAppKey, c.appKey)
	r.Header.Set(HTTPNonce, nonce)
	r.Header.Set(HTTPCurTime, curTime)
	r.Header.Set(HTTPCheckSum, checkSum)
}

func (c *Client) URL(path string) string {
	return c.host + path
}

func (c *Client) do(ctx context.Context, r *http.Request) (*http.Response, error) {
	c.AddCheckSum(r)
	r = r.WithContext(ctx)
	return c.http.Do(r)
}

func (c *Client) PostForm(ctx context.Context, path string, value any) (*http.Response, error) {
	r, err := http.NewRequest("POST", c.URL(path), nil)
	if err != nil {
		return nil, err
	}
	err = c.AddFormBody(r, value)
	if err != nil {
		return nil, err
	}
	return c.do(ctx, r)
}

func (c *Client) AddFormBody(r *http.Request, v any) error {
	form := url.Values{}
	err := c.formEncoder.Encode(v, form)
	if err != nil {
		return err
	}
	r.Body = io.NopCloser(strings.NewReader(form.Encode()))
	r.Header.Set(HTTPContentType, HTTPXFormURLEncoded)
	return nil
}

func (c *Client) JSONResponse(resp *http.Response, outPtr any) error {
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if r, ok := outPtr.(WithRawBody); ok {
		r.SetRawBody(raw)
	}
	return json.Unmarshal(raw, outPtr)
}
