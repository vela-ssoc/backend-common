package model

import "time"

// PassDNS DNS 白名单
type PassDNS struct {
	ID        int64     `json:"id,string"  gorm:"column:id;primaryKey"` // ID
	Domain    string    `json:"domain"     gorm:"column:domain"`        // 域名
	Kind      string    `json:"kind"       gorm:"column:kind"`          // 类型
	BeforeAt  time.Time `json:"before_at"  gorm:"column:before_at"`     // 有效期
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`    // 更新时间
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`    // 创建时间
}

// TableName implement gorm schema.Tabler
func (PassDNS) TableName() string {
	return "pass_dns"
}
