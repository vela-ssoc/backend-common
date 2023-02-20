package model

import "time"

// Userinfo 用户信息表，用于 session
type Userinfo struct {
	ID       int64     `json:"id,string" gorm:"column:id;primaryKey"` // User.ID
	Username string    `json:"username"  gorm:"column:username"`      // 用户名
	Nickname string    `json:"nickname"  gorm:"column:nickname"`      // 用户昵称
	Token    string    `json:"token"     gorm:"column:token"`         // token
	IssuedAt time.Time `json:"issued_at" gorm:"column:issued_at"`     // token 签发时间
	LastedAt time.Time `json:"lasted_at" gorm:"column:lasted_at"`     // 最后一次续约时间
}

// TableName implemented gorm schema.Tabler
func (Userinfo) TableName() string {
	return "userinfo"
}
