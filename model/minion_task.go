package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// MinionTask 配置运行
type MinionTask struct {
	ID          int64       `json:"id,string"           gorm:"column:id;primaryKey"`
	SubstanceID int64       `json:"substance_id,string" gorm:"column:substance_id"`
	MinionID    int64       `json:"minion_id,string"    gorm:"column:minion_id"`
	Inet        string      `json:"inet"                gorm:"column:inet"`
	Dialect     bool        `json:"dialect"             gorm:"column:dialect"`
	Name        string      `json:"name"                gorm:"column:name"`    // 配置名称
	Link        string      `json:"link"                gorm:"column:link"`    // 外链
	Status      string      `json:"status"              gorm:"column:status"`  // 运行状态
	Hash        string      `json:"hash"                gorm:"column:hash"`    // hash
	From        string      `json:"from"                gorm:"column:from"`    // 来源
	Uptime      time.Time   `json:"uptime"              gorm:"column:uptime"`  // 启动时间
	Failed      bool        `json:"failed"              gorm:"column:failed"`  // 是否失败
	Cause       string      `json:"cause"               gorm:"column:cause"`   // 如果发生失败，失败的原因
	Runners     TaskRunners `json:"runners"             gorm:"column:runners"` // task 内部服务
	CreatedAt   time.Time   `json:"created_at"          gorm:"column:created_at"`
}

// TableName implement gorm schema.Tabler
func (MinionTask) TableName() string {
	return "minion_task"
}

type TaskRunners []*TaskRunner

// Scan implement sql.Scanner
func (trs *TaskRunners) Scan(v any) error {
	return json.Unmarshal(v.([]byte), trs)
}

// Value implement driver.Valuer
func (trs TaskRunners) Value() (driver.Value, error) {
	return json.Marshal(trs)
}
