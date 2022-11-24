package main

import (
	"context"
	"log"
	"os"

	"github.com/not-a-dog/go-yunxin"
)

func main() {
	appKey := os.Getenv("WY_APP_KEY")
	appSecret := os.Getenv("WY_APP_SECRET")
	if len(appKey) == 0 || len(appSecret) == 0 {
		log.Fatal("must set env")
	}

	ctx := context.Background()
	client := yunxin.NewClient(appKey, appSecret)

	result, err := client.UserUpdate(ctx, &yunxin.UserUpdateParam{
		AccID: "b845d05e8bf13b3be57e09095c42f577",
		Token: "hello",
	})
	log.Printf("result %+v err %+v \n", result, err)
}
