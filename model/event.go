package model

import (
	"database/sql/driver"
	"time"
)

// Event 事件信息表
type Event struct {
	ID         int64      `json:"id,string"        gorm:"column:id;primaryKey"` // 消息 ID
	MinionID   int64      `json:"minion_id,string" gorm:"column:minion_id"`     // 节点 ID
	Inet       string     `json:"inet"             gorm:"column:inet"`          // 节点 IPv4
	Subject    string     `json:"subject"          gorm:"column:subject"`       // 主题
	RemoteAddr string     `json:"remote_addr"      gorm:"column:remote_addr"`   // 远程地址
	RemotePort int        `json:"remote_port"      gorm:"column:remote_port"`   // 远程端口
	FromCode   string     `json:"from_code"        gorm:"column:from_code"`     // 来源模块
	Typeof     string     `json:"typeof"           gorm:"column:typeof"`        // 模块类型
	User       string     `json:"user"             gorm:"column:user"`          // 用户信息
	Auth       string     `json:"auth"             gorm:"column:auth"`          // 认证信息
	Msg        string     `json:"msg"              gorm:"column:msg"`           // 上报消息
	Error      string     `json:"error"            gorm:"column:error"`         // 错误信息
	Region     string     `json:"region"           gorm:"column:region"`        // IP 定位
	Level      EventLevel `json:"level"            gorm:"column:level"`         // 告警级别
	HaveRead   bool       `json:"have_read"        gorm:"column:have_read"`     // 是否已读确认
	SendAlert  bool       `json:"send_alert"       gorm:"column:send_alert"`    // 是否需要发送告警
	Secret     string     `json:"-"                gorm:"column:secret"`        // 如果告警，生成随机字符串防止恶意遍历
	OccurAt    time.Time  `json:"occur_at"         gorm:"column:occur_at"`      // 事件发生的时间
	CreatedAt  time.Time  `json:"created_at"       gorm:"column:created_at"`    // 创建时间
}

// TableName implemented gorm schema.Tabler
func (Event) TableName() string {
	return "event"
}

// EventLevel 事件级别
// 用 int 表示是为了方便比较：level > EventHigh
type EventLevel uint8

func (lvl EventLevel) Value() (driver.Value, error) {
	str := eventIntMap[lvl]
	return str, nil
}

func (lvl *EventLevel) Scan(raw any) error {
	switch dat := raw.(type) {
	case string:
		*lvl = eventFmtMap[dat]
	case []byte:
		*lvl = eventFmtMap[string(dat)]
	}
	return nil
}

func (lvl EventLevel) MarshalText() ([]byte, error) {
	str := eventIntMap[lvl]
	return []byte(str), nil
}

func (lvl *EventLevel) UnmarshalText(dat []byte) error {
	*lvl = eventFmtMap[string(dat)]
	return nil
}

func (lvl EventLevel) String() string {
	str := eventIntMap[lvl]
	return str
}

const (
	EventNote EventLevel = iota
	EventMinor
	EventMajor
	EventCritical
)

var (
	eventFmtMap = map[string]EventLevel{
		"普通": EventNote,
		"次要": EventMinor,
		"重要": EventMajor,
		"紧急": EventCritical,
	}

	eventIntMap = map[EventLevel]string{
		EventNote:     "普通",
		EventMinor:    "次要",
		EventMajor:    "重要",
		EventCritical: "紧急",
	}
)
