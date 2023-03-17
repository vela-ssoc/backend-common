package model

import "time"

// Store 一些 KV 类型的数据存储
type Store struct {
	ID        string    `json:"id"         gorm:"column:id;primaryKey"` // 数据 ID，就是 key
	Value     []byte    `json:"value"      gorm:"column:value"`         // 数据的值
	Desc      string    `json:"desc"       gorm:"column:desc"`          // 说明
	Version   int64     `json:"version"    gorm:"column:version"`       // 乐观锁
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`    // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`    // 更新时间
}

// TableName implement gorm schema.Tabler
func (Store) TableName() string {
	return "store"
}
