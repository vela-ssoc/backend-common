package model

import "time"

type Rotate struct {
	ID        string    `json:"id"         gorm:"column:id"`         // 表名
	DayNum    int       `json:"day_num"    gorm:"column:day_num"`    // 单位：天
	RowNum    int       `json:"row_num"    gorm:"column:row_num"`    // 单位：万
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"` // 更新时间
}

// TableName implement schema.Tabler
func (Rotate) TableName() string {
	return "rotate"
}
