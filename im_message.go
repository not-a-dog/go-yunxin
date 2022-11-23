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
