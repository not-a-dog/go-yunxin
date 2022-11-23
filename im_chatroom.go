package yunxin

type ChatRoomCreateParam struct {
	Creator          string           `schema:"creator,required"`
	Name             string           `schema:"name,required"`
	Announcement     string           `schema:"announcement"`
	BroadcastURL     string           `schema:"broadcasturl"`
	Ext              string           `schema:"ext"`
	QueueLevel       QueueLevel       `schema:"queuelevel"`
	Bid              string           `schema:"bid"` // TODO json string
	DelayClosePolicy DelayClosePolicy `schema:"delayClosePolicy"`
	DelaySeconds     int64            `schema:"delaySeconds"`
}

type QueueLevel int

const (
	QueueLevelAny       QueueLevel = 0
	QueueLevelOnlyAdmin QueueLevel = 1
)

type DelayClosePolicy int

const (
	DelayClosePolicyNever         DelayClosePolicy = 0 // 不开启定时关闭
	DelayClosePolicyByTime        DelayClosePolicy = 1 // 固定时间关闭
	DelayClosePolicyOnEmptyByTime DelayClosePolicy = 2 // 空闲关闭（等聊天室中没有用户后固定时间关闭）
)

type ChatRoomCreateResponse struct {
	BasicResponse
	ChatRoom *ChatRoomInfo `json:"chatroom"`
}

type ChatRoomInfo struct {
	RoomID          int64             `json:"roomid"`
	Valid           bool              `json:"valid"`
	Announcement    string            `json:"announcement"`
	Name            string            `json:"name"`
	Broadcasturl    string            `json:"broadcasturl"`
	Ext             string            `json:"ext"`
	QueueLevel      QueueLevel        `json:"queuelevel"`
	Muted           bool              `json:"muted"`
	Creator         string            `json:"creator"`
	DelayInfo       ChatroomDelayInfo `json:"delayInfo"`
	Onlineusercount *int              `json:"onlineusercount,omitempty"`
}

type ChatroomDelayInfo struct {
	DelayCloseEnable bool `json:"delayCloseEnable"`
	DelayClosePolicy bool `json:"delayClosePolicy"`
	DelaySeconds     bool `json:"delaySeconds"`
	Status           bool `json:"status"`
}

type ChatroomGetParam struct {
	RoomID              int64 `schema:"roomid"`
	NeedOnlineUserCount bool  `schema:"needOnlineUserCount"`
}

type ChatroomGetResponse struct {
	BasicResponse
	ChatRoom *ChatRoomInfo `json:"chatroom"`
}

type ChatRoomRequestAddrParam struct {
	RoomID     int64      `schema:"roomid,required"`
	AccID      string     `schema:"accid,required"`
	ClientType ClientType `schema:"clienttype"`
	ClientIP   string     `schema:"clientip"`
}

type ClientType int

const (
	ClientTypeWeblink    = 1
	ClientTypeCommonlink = 2
	ClientTypeWechatlink = 3
)

type ChatRoomRequestAddrResponse struct {
	BasicResponse
	Addr []string `json:"addr"`
}

type ChatRoomSendMsgParam struct {
	RoomID                         int64   `schema:"roomid,required"`
	MsgID                          string  `schema:"msgId,required"`
	FromAccID                      string  `schema:"fromAccid,required"`
	MsgType                        MsgType `schema:"msgType"`
	SubType                        int     `schema:"subType"`
	ResendFlag                     int     `schema:"resendFlag"`
	Attach                         string  `schema:"attach"`
	Ext                            string  `schema:"ext"`
	SkipHistory                    int     `schema:"skipHistory"`
	AbandonRatio                   int     `schema:"abandonRatio"`
	HighPriority                   bool    `schema:"highPriority"`
	NeedHighPriorityMsgResend      bool    `schema:"needHighPriorityMsgResend"`
	UseYidun                       int     `schema:"useYidun"`
	YidunAntiCheating              string  `schema:"yidunAntiCheating"`
	YidunAntiSpamExt               string  `schema:"yidunAntiSpamExt"`
	Bid                            string  `schema:"bid"`
	Antispam                       string  `schema:"antispam"`
	NotifyTargetTags               string  `schema:"notifyTargetTags"`
	AntispamCustom                 string  `schema:"antispamCustom"`
	Env                            string  `schema:"env"`
	ChatMsgPriority                int     `schema:"chatMsgPriority"`
	ForbiddenIfHighPriorityMsgFreq int     `schema:"forbiddenIfHighPriorityMsgFreq"`
	LocX                           float64 `schema:"locX"`
	LocY                           float64 `schema:"locY"`
	LocZ                           float64 `schema:"locZ"`
}

type ChatRoomSendMsgResponse struct {
	BasicResponse
	Desc *ChatRoomMsg `json:"desc"`
}

type ChatRoomMsg struct {
	Time             string `json:"time"`
	FromAccount      string `json:"fromAccount"`
	FromNick         string `json:"fromNick"`
	FromAvator       string `json:"fromAvator"`
	MsgIDClient      string `json:"msgid_client"`
	FromClientType   string `json:"fromClientType"`
	Attach           string `json:"attach"`
	RoomID           string `json:"roomId"`
	Type             string `json:"type"`
	Ext              string `json:"ext"`
	HighPriorityFlag int    `json:"highPriorityFlag"`
	MsgAbandonFlag   int    `json:"msgAbandonFlag"`
}
