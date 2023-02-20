package model

import (
	"database/sql"
	"time"
)

// Minion 节点信息表
type Minion struct {
	ID         int64        `json:"id,string"  gorm:"column:id;primaryKey"` // ID
	Name       string       `json:"name"       gorm:"column:name"`          // 节点名字
	Inet       string       `json:"inet"       gorm:"column:inet"`          // IPv4
	Inet6      string       `json:"inet6"      gorm:"column:inet6"`         // IPv6
	Status     MinionStatus `json:"status"     gorm:"column:status"`        // 节点状态
	MAC        string       `json:"mac"        gorm:"column:mac"`           // MAC
	Goos       string       `json:"goos"       gorm:"column:goos"`          // Goos
	Arch       string       `json:"arch"       gorm:"column:arch"`          // Arch
	Semver     string       `json:"semver"     gorm:"column:semver"`        // 版本号
	CPU        int          `json:"cpu"        gorm:"column:cpu"`           // CPU 核心数
	PID        int          `json:"pid"        gorm:"column:pid"`           // 进程 PID
	Username   string       `json:"username"   gorm:"column:username"`      // 运行 agent 程序的 用户
	Hostname   string       `json:"hostname"   gorm:"column:hostname"`      // 主机名
	Workdir    string       `json:"workdir"    gorm:"column:workdir"`       // 工作目录
	Executable string       `json:"executable" gorm:"column:executable"`    // 执行路径
	PingedAt   sql.NullTime `json:"pinged_at"  gorm:"column:pinged_at"`     // 最近一次 ping 的时间
	JoinedAt   sql.NullTime `json:"joined_at"  gorm:"column:joined_at"`     // 最近一次加入（连接）时间
	CreatedAt  time.Time    `json:"created_at" gorm:"column:created_at"`    // 创建时间
	UpdatedAt  time.Time    `json:"updated_at" gorm:"column:updated_at"`    // 更新时间
}

// TableName gorm table name
func (Minion) TableName() string {
	return "minion"
}

type MinionStatus uint8

const (
	MinionInactive MinionStatus = iota // 未激活
	MinionOffline                      // 离线
	MinionOnline                       // 在线
	MinionRemove                       // 移除
)
