package model

import (
	"database/sql/driver"
	"time"
)

type Risk struct {
	ID       int64  `json:"id,string"        gorm:"column:id;primaryKey"` // 数据ID
	MinionID int64  `json:"minion_id,string" gorm:"column:minion_id"`     // 节点ID
	Inet     string `json:"inet"             gorm:"column:inet"`          // 节点 IPv4

	// RiskType 风险类型
	// ["暴力破解", "病毒事件", "弱口令", "数据爬虫", "蜜罐应用", "web 攻击", "监控事件", "登录事件"]
	RiskType   string    `json:"risk_type"   gorm:"column:risk_type"`
	Level      RiskLevel `json:"level"       gorm:"column:level"`       // 风险级别
	Payload    string    `json:"payload"     gorm:"column:payload"`     // 攻击载荷
	Subject    string    `json:"subject"     gorm:"column:subject"`     // 风险事件主题
	LocalIP    string    `json:"local_ip"    gorm:"column:local_ip"`    // 本地 IP
	LocalPort  int       `json:"local_port"  gorm:"column:local_port"`  // 本地端口
	RemoteIP   string    `json:"remote_ip"   gorm:"column:remote_ip"`   // 远程 IP
	RemotePort int       `json:"remote_port" gorm:"column:remote_port"` // 远程端口
	FromCode   string    `json:"from_code"   gorm:"column:from_code"`   // 来源模块
	Region     string    `json:"region"      gorm:"column:region"`      // IP 归属地
	Reference  string    `json:"reference"   gorm:"column:reference"`   // 参考引用
	SendAlert  bool      `json:"send_alert"  gorm:"column:send_alert"`  // 是否发送告警
	Secret     string    `json:"-"           gorm:"column:secret"`      // 查询密文
	// Status     RiskStatus `json:"status"      gorm:"column:status"`      // 状态: 1-未处理 2-已处理 3-忽略
	OccurAt   time.Time `json:"occur_at"    gorm:"column:occur_at"`   // 风险产生的时间
	CreatedAt time.Time `json:"created_at"  gorm:"column:created_at"` // 入库保存时间
}

// TableName implemented gorm schema.Tabler
func (Risk) TableName() string {
	return "risk"
}

// RiskLevel 风险级别
// 用 int 表示是为了方便比较：level > RiskHigh
type RiskLevel uint8

func (lvl RiskLevel) Value() (driver.Value, error) {
	str := riskIntMap[lvl]
	return str, nil
}

func (lvl *RiskLevel) Scan(raw any) error {
	switch dat := raw.(type) {
	case string:
		*lvl = riskFmtMap[dat]
	case []byte:
		*lvl = riskFmtMap[string(dat)]
	}
	return nil
}

func (lvl RiskLevel) MarshalText() ([]byte, error) {
	str := riskIntMap[lvl]
	return []byte(str), nil
}

func (lvl *RiskLevel) UnmarshalText(dat []byte) error {
	*lvl = riskFmtMap[string(dat)]
	return nil
}

func (lvl RiskLevel) String() string {
	str := riskIntMap[lvl]
	return str
}

const (
	RiskLow RiskLevel = iota
	RiskMiddle
	RiskHigh
	RiskCritical
)

var (
	riskFmtMap = map[string]RiskLevel{
		"低危": RiskLow,
		"中危": RiskMiddle,
		"高危": RiskHigh,
		"紧急": RiskCritical,
	}

	riskIntMap = map[RiskLevel]string{
		RiskLow:      "低危",
		RiskMiddle:   "中危",
		RiskHigh:     "高危",
		RiskCritical: "紧急",
	}
)
