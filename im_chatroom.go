package yunxin

type ChatRoomCreateParam struct {
	Creator          string           `schema:"creator,required"`
	Name             string           `schema:"name,required"`
	Announcement     string           `schema:"announcement"`
	Broadcasturl     string           `schema:"broadcasturl"`
	Ext              string           `schema:"ext"`
	QueueLevel       QueueLevel       `schema:"queuelevel"`
	Bid              string           `schema:"bid"` // TODO json string
	DelayClosePolicy DelayClosePolicy `schema:"delayClosePolicy"`
	DelaySeconds     int64            `schema:"delaySeconds"`
}

func (ChatRoomCreateParam) GetPath() string {
	return PathChatRoomCreate
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
	BasicResponese
	ChatRoomCreateParam *ChatRoomInfo `json:"chatroom"`
}

type ChatRoomInfo struct {
	Roomid       string            `json:"roomid"`
	Valid        bool              `json:"valid"`
	Announcement string            `json:"announcement"`
	Name         string            `json:"name"`
	Broadcasturl string            `json:"broadcasturl"`
	Ext          string            `json:"ext"`
	QueueLevel   QueueLevel        `json:"queuelevel"`
	Muted        bool              `json:"muted"`
	Creator      string            `json:"creator"`
	DelayInfo    ChatroomDelayInfo `json:"delayInfo"`
}

type ChatroomDelayInfo struct {
	DelayCloseEnable bool `json:"delayCloseEnable"`
	DelayClosePolicy bool `json:"delayClosePolicy"`
	DelaySeconds     bool `json:"delaySeconds"`
	Status           bool `json:"status"`
}

func (im *IM) ChatRoomCreate(param *ChatRoomCreateParam) (*ChatRoomCreateResponse, error) {
	var result ChatRoomCreateResponse
	err := im.PostFormAs(param, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
