package model

import "time"

// Domain 域名解析信息
type Domain struct {
	ID        int64     `json:"id,string"  gorm:"column:id;primaryKey"` // 主键
	Record    string    `json:"record"     gorm:"column:record"`        // 域名
	Type      string    `json:"type"       gorm:"column:type"`          // 解析类型
	Addr      string    `json:"addr"       gorm:"column:addr"`          // 解析地址
	Origin    string    `json:"origin"     gorm:"column:origin"`        // 来源
	ISP       string    `json:"isp"        gorm:"column:isp"`           // ISP
	Comment   string    `json:"comment"    gorm:"column:comment"`       // 备注
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`    // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`    // 更新时间
}

// TableName implement gorm schema.Tabler
func (Domain) TableName() string {
	return "domain"
}
