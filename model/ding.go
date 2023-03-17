package model

import "time"

type Ding struct {
	ID        string    `json:"id"         gorm:"column:id;primaryKey"` // ID 帐号
	Code      string    `json:"code"       gorm:"column:code"`          // 验证码
	Tries     int       `json:"tries"      gorm:"column:tries"`         // 尝试次数
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`    // 创建时间
}

// TableName implement gorm schema.Tabler
func (Ding) TableName() string {
	return "ding"
}
