package model

import "time"

// JobCode 下发任务的代码
type JobCode struct {
	ID        int64     `json:"id,string"  gorm:"column:id;primaryKey"` // ID
	Name      string    `json:"name"       gorm:"column:name"`          // 名称
	Icon      []byte    `json:"icon"       gorm:"column:icon"`          // 图标
	Chunk     []byte    `json:"chunk"      gorm:"column:chunk"`         // 代码
	Desc      string    `json:"desc"       gorm:"column:desc"`          // 说明
	Hash      string    `json:"hash"       gorm:"column:hash"`          // 哈希
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`    // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`    // 更新时间
}

// TableName implement gorm schema.Tabler
func (JobCode) TableName() string {
	return "job_code"
}
