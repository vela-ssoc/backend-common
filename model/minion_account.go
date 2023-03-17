package model

import "time"

type MinionAccount struct {
	ID          int64     `json:"id,string"        gorm:"column:id;primaryKey"`
	MinionID    int64     `json:"minion_id,string" gorm:"column:minion_id"`
	Inet        string    `json:"inet"             gorm:"column:inet"`
	Name        string    `json:"name"             gorm:"column:name"`
	LoginName   string    `json:"login_name"       gorm:"column:login_name"`
	UID         string    `json:"uid"              gorm:"column:uid"`
	GID         string    `json:"gid"              gorm:"column:gid"`
	HomeDir     string    `json:"home_dir"         gorm:"column:home_dir"`
	Description string    `json:"description"      gorm:"column:description"`
	Status      string    `json:"status"           gorm:"column:status"`
	Raw         string    `json:"raw"              gorm:"column:raw"`
	CreatedAt   time.Time `json:"created_at"       gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at"       gorm:"column:updated_at"`
}

// TableName implement gorm schema.Tabler
func (MinionAccount) TableName() string {
	return "minion_account"
}
