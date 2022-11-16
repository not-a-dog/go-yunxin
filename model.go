package yunxin

type MsgSendMsgParamTODO struct {
	From               string          `schema:"from,required"`
	Ope                OpeType         `schema:"ope,required"`
	To                 string          `schema:"to,required"`
	Type               MsgType         `schema:"type,required"`
	Body               string          `schema:"body,required"` // JSON 格式
	MsgDesc            string          `schema:"msgDesc"`
	Option             *SendMsgOption  `schema:"option"`
	Pushcontent        string          `schema:"pushcontent"`
	Payload            string          `schema:"payload"` // JSON 格式
	Ext                string          `schema:"ext"`
	Forcepushlist      StringSlice     `schema:"forcepushlist"`
	Forcepushcontent   string          `schema:"forcepushcontent"`
	Forcepushall       bool            `schema:"forcepushall"`
	Bid                string          `schema:"bid"`
	Antispam           bool            `schema:"antispam"`
	AntispamCustom     *AntispamCustom `schema:"antispamCustom"`
	UseYidun           *int            `schema:"useYidun"`          // https://doc.yunxin.163.com/docs/TM5MzM5Njk/DEwMTE3NzQ?platformId=60353
	YidunAntiCheating  string          `schema:"yidunAntiCheating"` // JSON
	YidunAntiSpamExt   string          `schema:"yidunAntiSpamExt"`  // JSON
	MarkRead           *int            `schema:"markRead"`
	CheckFriend        bool            `schema:"checkFriend"`
	SubType            int             `schema:"subType,omitempty"`
	MsgSenderNoSense   int             `schema:"msgSenderNoSense"`
	MsgReceiverNoSense int             `schema:"msgReceiverNoSense"`
	Env                string          `schema:"env"`
}

type OpeType int

// 0：点对点个人消息，1：群消息（高级群）
const (
	OpeTypeUser  OpeType = 0
	OpeTypeGroup OpeType = 1
)

type MsgType int

const (
	MsgTypeText     MsgType = 0   // 0 表示文本
	MsgTypePic      MsgType = 1   // 1 表示图片
	MsgTypeAudio    MsgType = 2   // 2 表示语音
	MsgTypeVideo    MsgType = 3   // 3 表示视频
	MsgTypeLocation MsgType = 4   // 4 表示地理位置信息
	MsgTypeFile     MsgType = 6   // 6 表示文件
	MsgTypeTips     MsgType = 10  // 10 表示提示消息
	MsgTypeCustom   MsgType = 100 // 100 表示自定义消息
)

type AntispamCustom struct {
	Type AntispamCustomType `json:"type"`
	Data string             `json:"data"` // 文本内容or图片地址
}

type AntispamCustomType int

// 1：文本，2：图片
const (
	AntispamCustomTypeText AntispamCustomType = 1
	AntispamCustomTypePic  AntispamCustomType = 2
)

type SendMsgOption struct {
	// 1. 该消息是否需要漫游，默认true（需要app开通漫游消息功能）；
	Roam *bool `json:"roam,omitempty"`
	// 2. 该消息是否存云端历史，默认true
	History *bool `json:"history,omitempty"`
	// 3. 该消息是否需要发送方多端同步，默认true；
	Sendersync *bool `json:"sendersync,omitempty"`
	// 4. 该消息是否需要APNS推送或安卓系统通知栏推送，默认true；
	Push *bool `json:"push,omitempty"`
	// 5. 该消息是否需要抄送第三方；默认true (需要app开通消息抄送功能);
	Route *bool `json:"route,omitempty"`
	// 6. 该消息是否需要计入到未读计数中，默认true;
	Badge *bool `json:"badge,omitempty"`
	// 7. 推送文案是否需要带上昵称，不设置该参数时默认true;
	NeedPushNick *bool `json:"needPushNick,omitempty"`
	// 8. 是否需要存离线消息，不设置该参数时默认true;
	Persistent *bool `json:"persistent,omitempty"`
	// 9. 是否将本消息更新到会话列表服务里本会话的lastmsg，默认true。
	SessionUpdate *bool `json:"sessionUpdate,omitempty"`
}

const (
	MarkReadTypeNoNeed int = 0
	MarkReadTypeNeed   int = 1
)

/*
checkFriend	boolean	否	是否为好友关系才发送消息，默认否
注：使用该参数需要先开通功能服务
subType	int	否	自定义消息子类型，大于0
msgSenderNoSense	int	否	发送方是否无感知。0-有感知，1-无感知。若无感知，则消息发送者无该消息的多端、漫游、历史记录等。
msgReceiverNoSense	int	否	接受方是否无感知。0-有感知，1-无感知。若无感知，则消息接收者者无该消息的多端、漫游、历史记录等
env	String	否	所属环境，根据env可以配置不同的抄送地址
*/
