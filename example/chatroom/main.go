package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

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
	result, err := client.ChatroomGet(ctx, &yunxin.ChatroomGetParam{
		RoomID:              2757903841,
		NeedOnlineUserCount: true,
	})
	log.Printf("result %+v err %+v \n", result, err)

	result2, err := client.ChatRoomRequestAddr(ctx, &yunxin.ChatRoomRequestAddrParam{
		RoomID:     2757903841,
		AccID:      "b845d05e8bf13b3be57e09095c42f577",
		ClientType: yunxin.ClientTypeWeblink,
	})
	log.Printf("result2 %+v err %+v \n", result2, err)

	result3, err := client.ChatRoomSendMsg(ctx, &yunxin.ChatRoomSendMsgParam{
		RoomID:    2757903841,
		FromAccID: "b845d05e8bf13b3be57e09095c42f577",
		MsgID:     "test-chatroom-" + strconv.FormatInt(time.Now().UnixMicro(), 10),
		Attach:    "hello",
		MsgType:   yunxin.MsgTypeText,
		SubType:   1,
	})
	log.Printf("result3 %+v err %+v \n", result3, err)

	result4, err := client.ChatroomGetBatch(ctx, &yunxin.ChatroomGetBatchParam{
		RoomIDs:             []string{"2757903841"},
		NeedOnlineUserCount: true,
	})
	log.Printf("result4 %+v err %+v \n", result4, err)
}
