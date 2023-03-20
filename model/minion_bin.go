package model

import "time"

// MinionBin minion 节点二进制发行版本记录表
type MinionBin struct {
	ID         int64     `json:"id,string"  gorm:"column:id;primaryKey"`
	FileID     int64     `json:"-"          gorm:"column:file_id"`
	Goos       string    `json:"goos"       gorm:"column:goos"`
	Arch       string    `json:"arch"       gorm:"column:arch"`
	Name       string    `json:"name"       gorm:"column:name"`
	Size       int64     `json:"size"       gorm:"column:size"`
	Hash       string    `json:"hash"       gorm:"column:hash"`
	Semver     Semver    `json:"semver"     gorm:"column:semver"`
	Changelog  string    `json:"changelog"  gorm:"column:changelog"`
	Weight     int64     `json:"-"          gorm:"column:weight"`
	Deprecated bool      `json:"deprecated" gorm:"column:deprecated"` // 是否已过期
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// TableName implement gorm schema.Tabler
func (MinionBin) TableName() string {
	return "minion_bin"
}
