package model

import "time"

type RiskFile struct {
	ID        int64     `json:"id,string"  gorm:"column:id;primaryKey"`
	Checksum  string    `json:"checksum"   gorm:"column:checksum"`
	Algorithm string    `json:"algorithm"  gorm:"column:algorithm"`
	Kind      string    `json:"kind"       gorm:"column:kind"`
	Origin    string    `json:"origin"     gorm:"column:origin"`
	Desc      string    `json:"desc"       gorm:"column:desc"`
	BeforeAt  time.Time `json:"before_at"  gorm:"column:before_at"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// TableName implement gorm schema.Tabler
func (RiskFile) TableName() string {
	return "risk_file"
}
