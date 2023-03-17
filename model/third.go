package model

import "time"

// Third 3rd 文件表
type Third struct {
	ID        int64     `json:"id,string"         gorm:"column:id;primaryKey"` // ID 此处取 File 的 ID
	FileID    int64     `json:"file_id,string"    gorm:"column:file_id"`       // 文件 ID
	Name      string    `json:"name"              gorm:"column:name"`          // 文件名字
	Hash      string    `json:"hash"              gorm:"column:hash"`          // 文件 hash
	Path      string    `json:"path"              gorm:"column:path"`          // 文件发布路径
	Desc      string    `json:"desc"              gorm:"column:desc"`          // 文件简介
	Size      int64     `json:"size"              gorm:"column:size"`          // 文件大小
	Extension string    `json:"extension"         gorm:"column:extension"`     // 扩展名
	CreatedID int64     `json:"created_id,string" gorm:"column:created_id"`    // 创建者 ID
	UpdatedID int64     `json:"updated_id,string" gorm:"column:updated_id"`    // 修改者 ID
	CreatedAt time.Time `json:"created_at"        gorm:"column:created_at"`    // 创建时间
	UpdatedAt time.Time `json:"updated_at"        gorm:"column:updated_at"`    // 修改时间
}

// TableName implement gorm schema.Tabler
func (Third) TableName() string {
	return "third"
}
