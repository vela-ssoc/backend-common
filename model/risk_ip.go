package model

import "time"

type RiskIP struct {
	ID        int64     `json:"id,string"  gorm:"column:id;primaryKey"`
	IP        string    `json:"ip"         gorm:"column:ip"`
	Kind      string    `json:"kind"       gorm:"column:kind"`
	Origin    string    `json:"origin"     gorm:"column:origin"`
	BeforeAt  time.Time `json:"before_at"  gorm:"column:before_at"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// TableName implement gorm schema.Tabler
func (RiskIP) TableName() string {
	return "risk_ip"
}
