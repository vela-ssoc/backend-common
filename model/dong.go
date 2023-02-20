package model

import "time"

// Dong 咚咚服务号配置
type Dong struct {
	ID        int64     `json:"id,string"  gorm:"column:id;primaryKey"` // ID
	Name      string    `json:"name"       gorm:"column:name"`          // 名字
	Addr      string    `json:"addr"       gorm:"column:addr"`          // 咚咚服务器
	Host      string    `json:"host"       gorm:"column:host"`          // HTTP Header Host
	Account   string    `json:"account"    gorm:"column:account"`       // 咚咚 Account
	Token     string    `json:"token"      gorm:"column:token"`         // 咚咚 Token
	Enable    bool      `json:"enable"     gorm:"column:enable"`        // 是否启用
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`    // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`    // 更新时间
}

// TableName implemented gorm schema.Tabler
func (Dong) TableName() string {
	return "dong"
}
