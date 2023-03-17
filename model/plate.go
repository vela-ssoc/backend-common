package model

import "time"

type Plate struct {
	ID          string    `json:"id"           gorm:"column:id;primaryKey"`
	Title       []byte    `json:"title"        gorm:"column:title"`
	Body        []byte    `json:"body"         gorm:"column:body"`
	NeedTitle   bool      `json:"need_title"   gorm:"column:need_title"`
	NeedBody    bool      `json:"need_body"    gorm:"column:need_body"`
	EscapeTitle bool      `json:"escape_title" gorm:"column:escape_title"`
	EscapeBody  bool      `json:"escape_body"  gorm:"column:escape_body"`
	Desc        string    `json:"desc"         gorm:"column:desc"`
	CreatedAt   time.Time `json:"created_at"   gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at"   gorm:"column:updated_at"`
}

func (Plate) TableName() string {
	return "plate"
}
