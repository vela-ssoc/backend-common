package model

import "time"

type MinionListen struct {
	ID        int64     `json:"id,string"        gorm:"column:id;primaryKey"`
	MinionID  int64     `json:"minion_id,string" gorm:"column:minion_id"`
	Inet      string    `json:"inet"             gorm:"column:inet"`
	RecordID  string    `json:"record_id"        gorm:"column:record_id"`
	PID       uint32    `json:"pid"              gorm:"column:pid"`
	FD        int       `json:"fd"               gorm:"column:fd"`
	Family    uint8     `json:"family"           gorm:"column:family"`
	Protocol  uint8     `json:"protocol"         gorm:"column:protocol"`
	LocalIP   string    `json:"local_ip"         gorm:"column:local_ip"`
	LocalPort int       `json:"local_port"       gorm:"column:local_port"`
	Path      string    `json:"path"             gorm:"column:path"`
	Process   string    `json:"process"          gorm:"column:process"`
	Username  string    `json:"username"         gorm:"column:username"`
	CreatedAt time.Time `json:"created_at"       gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at"       gorm:"column:updated_at"`
}

func (MinionListen) TableName() string {
	return "minion_listen"
}
