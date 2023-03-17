package model

import "time"

// MinionProcess minion 节点收集上来的进程列表
type MinionProcess struct {
	ID           int64     `json:"id,string"        gorm:"column:id;primaryKey"`
	MinionID     int64     `json:"minion_id,string" gorm:"column:minion_id"`
	Inet         string    `json:"inet"             gorm:"column:inet"`
	Name         string    `json:"name"             gorm:"column:name"`
	State        string    `json:"state"            gorm:"column:state"`
	Pid          int       `json:"pid"              gorm:"column:pid"`
	Ppid         int       `json:"ppid"             gorm:"column:ppid"`
	Pgid         uint32    `json:"pgid"             gorm:"column:pgid"`
	Cmdline      string    `json:"cmdline"          gorm:"column:cmdline"`
	Username     string    `json:"username"         gorm:"column:username"`
	Cwd          string    `json:"cwd"              gorm:"column:cwd"`
	Executable   string    `json:"executable"       gorm:"column:executable"`
	Args         []string  `json:"args"             gorm:"column:args;json"`
	UserTicks    uint64    `json:"user_ticks"       gorm:"column:user_ticks"`
	TotalPct     float64   `json:"total_pct"        gorm:"column:total_pct"`
	TotalNormPct float64   `json:"total_norm_pct"   gorm:"column:total_norm_pct"`
	SystemTicks  uint64    `json:"system_ticks"     gorm:"column:system_ticks"`
	TotalTicks   uint64    `json:"total_ticks"      gorm:"column:total_ticks"`
	StartTime    string    `json:"start_time"       gorm:"column:start_time"`
	MemSize      uint64    `json:"mem_size"         gorm:"column:mem_size"`
	RssBytes     uint64    `json:"rss_bytes"        gorm:"column:rss_bytes"`
	RssPct       float64   `json:"rss_pct"          gorm:"column:rss_pct"`
	Share        uint64    `json:"share"            gorm:"column:share"`
	Checksum     string    `json:"checksum"         gorm:"column:checksum"`
	CreatedTime  time.Time `json:"created_time"     gorm:"column:created_time"`
	ModifiedAt   time.Time `json:"modified_at"      gorm:"column:modified_at"`
	CreatedAt    time.Time `json:"created_at"       gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at"       gorm:"column:updated_at"`
}

// TableName implement gorm schema.Tabler
func (MinionProcess) TableName() string {
	return "minion_process"
}
