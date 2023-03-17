package model

import (
	"encoding/json"
	"time"
)

// JobPolicy 任务策略
type JobPolicy struct {
	ID        int64           `json:"id,string"         gorm:"column:id;primaryKey"` // ID
	Name      string          `json:"name"              gorm:"column:name"`          // 名字
	Desc      string          `json:"desc"              gorm:"column:desc"`          // 说明
	CodeID    int64           `json:"code_id,string"    gorm:"column:code_id"`       // 代码 ID
	Timeout   int             `json:"timeout"           gorm:"column:timeout"`       // 执行超时时间
	Parallel  int             `json:"parallel"          gorm:"column:parallel"`      // 同时推送个数
	Args      json.RawMessage `json:"args"              gorm:"column:args"`          // 参数
	CreatedID int64           `json:"created_id,string" gorm:"column:created_id"`    // 创建者
	CreatedAt time.Time       `json:"created_at"        gorm:"column:created_at"`    // 创建时间
	UpdatedAt time.Time       `json:"updated_at"        gorm:"column:updated_at"`    // 更新时间
}

// TableName implement gorm schema.Tabler
func (JobPolicy) TableName() string {
	return "job_policy"
}
