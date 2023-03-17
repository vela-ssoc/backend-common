package model

import "time"

type PassIP struct {
	ID        int64     `json:"id,string"  gorm:"column:id;primaryKey"`
	IP        string    `json:"ip"         gorm:"column:ip"`
	Kind      string    `json:"kind"       gorm:"column:kind"`
	BeforeAt  time.Time `json:"before_at"  gorm:"column:before_at"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (PassIP) TableName() string {
	return "pass_ip"
}
