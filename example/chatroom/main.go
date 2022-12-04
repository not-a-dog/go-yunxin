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
	client := yunxin.NewClient(appKey, appSecret,
		yunxin.WithLogger(yunxin.StdLogger()))

	_, _ = client.ChatRoomUpdateInOutNotification(ctx, &yunxin.ChatRoomUpdateInOutNotificationParam{
		RoomID: 2757903841,
		Close:  true,
	})
}
