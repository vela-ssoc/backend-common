package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64          `json:"id"         gorm:"column:id;primaryKey"`
	Username  string         `json:"username"   gorm:"column:username"`
	Nickname  string         `json:"nickname"   gorm:"column:nickname"`
	Password  string         `json:"password"   gorm:"column:password"`
	Dong      string         `json:"dong"       gorm:"column:dong"`
	Enable    bool           `json:"enable"     gorm:"column:enable"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "user"
}
