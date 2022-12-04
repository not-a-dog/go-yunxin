package yunxin

type MsgType int

const (
	MsgTypeText   MsgType = 0
	MsgTypePic    MsgType = 1
	MsgTypeAudio  MsgType = 2
	MsgTypeVideo  MsgType = 3
	MsgTypeLoc    MsgType = 4
	MsgTypeFile   MsgType = 6
	MsgTypeTips   MsgType = 10
	MsgTypeCustom MsgType = 100
)

type MsgDelMsgParam struct {
	DeleteMsgID string        `schema:"deleteMsgid,required"`
	TimeTag     int64         `schema:"timetag,required"`
	Type        DeleteMsgType `schema:"type,required"`
	From        string        `schema:"from,required"`
	To          string        `schema:"to,required"`
}

type DeleteMsgType int

const (
	DeleteMsgTypeSingle DeleteMsgType = 7
	DeleteMsgTypeTeam   DeleteMsgType = 8
)

type MsgDelMsgResponse struct {
	BasicResponse
}
