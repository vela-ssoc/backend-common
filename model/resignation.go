package model

import "time"

// Resignation 离职员工表
type Resignation struct {
	ID        string    `json:"id"         gorm:"column:id;primaryKey"` // 员工工号
	Name      string    `json:"name"       gorm:"column:name"`          // 员工名字
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`    // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`    // 更新时间
}

// TableName implemented gorm schema.Tabler
func (Resignation) TableName() string {
	return "resignation"
}
