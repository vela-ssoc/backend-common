package model

import (
	"database/sql"
	"time"
)

// BrokerV2 节点信息表
type BrokerV2 struct {
	ID         int64        `json:"id"         gorm:"column:id;primaryKey"` // broker 节点 ID
	Name       string       `json:"name"       gorm:"column:name"`          // broker 节点名字
	Inet       string       `json:"inet"       gorm:"column:inet"`          // 出口 IP，一般为 IPv4
	MAC        string       `json:"mac"        gorm:"column:mac"`           // 出口 IP 所在的网卡的 MAC 地址
	Goos       string       `json:"goos"       gorm:"column:goos"`          // runtime.GOOS 节点操作系统类型
	Arch       string       `json:"arch"       gorm:"column:arch"`          // runtime.GOARCH 节点 CPU 架构
	CPU        int          `json:"cpu"        gorm:"column:cpu"`           // CPU 核心数
	Semver     Semver       `json:"semver"     gorm:"column:semver"`        // 节点版本
	Status     bool         `json:"status"     gorm:"column:status"`        // 状态 true-在线 false-离线
	Secret     string       `json:"secret"     gorm:"column:secret"`        // 连接认证密钥
	PID        int          `json:"pid"        gorm:"column:pid"`           // os.Getpid
	Username   string       `json:"username"   gorm:"column:username"`      // user.Current 运行 broker 程序的系统用户名
	Hostname   string       `json:"hostname"   gorm:"column:hostname"`      // os.Hostname 节点主机名
	Workdir    string       `json:"workdir"    gorm:"column:workdir"`       // os.Getwd
	Executable string       `json:"executable" gorm:"column:executable"`    // os.Executable
	PingedAt   sql.NullTime `json:"pinged_at"  gorm:"column:pinged_at"`     // 最近一次心跳时间
	JoinedAt   sql.NullTime `json:"joined_at"  gorm:"column:joined_at"`     // 最近一次上线时间
	CreatedAt  time.Time    `json:"created_at" gorm:"column:created_at"`    // 创建时间
	UpdatedAt  time.Time    `json:"updated_at" gorm:"column:updated_at"`    // 最近一次更新时间
}

// TableName implemented gorm schema.Tabler
func (BrokerV2) TableName() string {
	return "broker"
}

// Broker 代理节点表
type Broker struct {
	ID          int64     `json:"id,string"    gorm:"column:id;primaryKey"` // broker 节点 ID
	Name        string    `json:"name"         gorm:"column:name"`          // 名字
	Servername  string    `json:"servername"   gorm:"column:servername"`    // servername minion 节点 TLS 认证用
	LAN         []string  `json:"lan"          gorm:"column:lan;json"`      // 内网地址
	VIP         []string  `json:"vip"          gorm:"column:vip;json"`      // 外网地址
	Status      bool      `json:"status"       gorm:"column:status"`        // 状态
	Secret      string    `json:"secret"       gorm:"column:secret"`        // 随机密钥防止恶意攻击
	Bind        string    `json:"bind"         gorm:"column:bind"`          // 服务监听地址
	HeartbeatAt time.Time `json:"heartbeat_at" gorm:"column:heartbeat_at"`  // 最近一次心跳时间
	CreatedAt   time.Time `json:"created_at"   gorm:"column:created_at"`    // 创建时间
	UpdatedAt   time.Time `json:"updated_at"   gorm:"column:updated_at"`    // 更新时间
}

// TableName implement gorm schema.Tabler
func (Broker) TableName() string {
	return "broker"
}
