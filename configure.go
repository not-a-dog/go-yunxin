package yunxin

import "time"

type Configure struct {
	Timeout time.Duration
	Host    string
}

const DefaultHost = "https://api.netease.im/nimserver"
