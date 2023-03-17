package model

import "time"

type MinionLogon struct {
	ID        int64     `json:"id,string"        gorm:"column:id;primaryKey"`
	MinionID  int64     `json:"minion_id,string" gorm:"column:minion_id"`
	Inet      string    `json:"inet"             gorm:"column:inet"`
	User      string    `json:"user"             gorm:"column:user"`
	Addr      string    `json:"addr"             gorm:"column:addr"`
	Msg       string    `json:"msg"              gorm:"column:msg"`
	Type      string    `json:"type"             gorm:"column:type"`
	PID       int       `json:"pid"              gorm:"column:pid"`
	Device    string    `json:"device"           gorm:"column:device"`
	Process   string    `json:"process"          gorm:"column:process"`
	LogonAt   time.Time `json:"logon_at"         gorm:"column:logon_at"`
	Ignore    bool      `json:"ignore"           gorm:"column:ignore"`
	CreatedAt time.Time `json:"created_at"       gorm:"column:created_at"`
}

func (MinionLogon) TableName() string {
	return "minion_logon"
}
