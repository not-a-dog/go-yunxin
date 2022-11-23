# go-yunxin

https://dev.yunxin.163.com/ server side API sdk

https://pkg.go.dev/github.com/not-a-dog/go-yunxin

## Usage

```go

client := yunxin.NewClient("", "")
user, err := client.GetUserInfos(ctx, &yunxin.GetUserInfosParam{Accids: []string{"111111"}})
```


## Dev

### 添加接口

1. 从 yunxin 文档中找到相关接口 XXX
2. `path.go` 中定义 PathXXX
3. `xxx.go` 中定义好 XXXParam XXXResponse
4. go generate
5. go build ./...

### 测试

```bash
WY_APP_KEY=xxx WY_APP_SECRET=yyy go run example/chatroom/main.go
```
