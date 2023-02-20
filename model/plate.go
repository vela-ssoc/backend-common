package model

import "time"

// Plate 模板配置表
type Plate struct {
	ID        string    `json:"id"         gorm:"column:id;primaryKey"` // 模板 ID
	Tmpl      []byte    `json:"tmpl"       gorm:"column:tmpl"`          // 模板内容
	Desc      string    `json:"desc"       gorm:"column:desc"`          // 模板说明
	Escape    bool      `json:"escape"     gorm:"column:escape"`        // 是否转义
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`    // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`    // 更新时间
}

// TableName implemented gorm schema.Tabler
func (Plate) TableName() string {
	return "plate"
}
