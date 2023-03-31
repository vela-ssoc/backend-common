package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type UserDomain uint8

const (
	UdLocal UserDomain = iota + 1
	UdOA
)

func (ud UserDomain) String() string {
	switch ud {
	case UdLocal:
		return "本地帐号"
	case UdOA:
		return "OA帐号"
	default:
		return "未知帐号"
	}
}

// User 用户表
type User struct {
	ID        int64          `json:"id,string"  gorm:"column:id;primaryKey"` // 用户 ID
	Username  string         `json:"username"   gorm:"column:username"`      // 用户名
	Nickname  string         `json:"nickname"   gorm:"column:nickname"`      // 用户昵称
	Password  string         `json:"-"          gorm:"column:password"`      // 密码
	Dong      string         `json:"dong"       gorm:"column:dong"`          // 咚咚号(用于接收通知)
	Enable    bool           `json:"enable"     gorm:"column:enable"`        // 是否启用
	Domain    UserDomain     `json:"domain"     gorm:"column:domain"`        // 帐号归属域，1-本地账户 2-OA账户
	Token     string         `json:"-"          gorm:"column:token"`         // Token
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`    // 创建时间
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`    // 更新时间
	IssueAt   sql.NullTime   `json:"issue_at"   gorm:"column:issue_at"`      // 最近一次 Token 签发时间
	SessionAt sql.NullTime   `json:"session_at" gorm:"column:session_at"`    // session 最近一次活动时间
	DeletedAt gorm.DeletedAt `json:"-"          gorm:"column:deleted_at"`    // 逻辑删除标志
}

// TableName implement gorm schema.Tabler
func (User) TableName() string {
	return "user"
}

func (u User) IsLocal() bool {
	return u.Domain == UdLocal
}

func (u User) IsOA() bool {
	return u.Domain == UdOA
}
