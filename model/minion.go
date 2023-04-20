package model

import (
	"database/sql"
	"time"
)

type MinionStatus uint8

const (
	// MSInactive 未激活
	MSInactive MinionStatus = iota + 1
	// MSOffline 离线
	MSOffline
	// MSOnline 在线
	MSOnline
	// MSDelete 已删除
	MSDelete
)

func (ms MinionStatus) String() string {
	switch ms {
	case MSInactive:
		return "未激活"
	case MSOffline:
		return "离线"
	case MSOnline:
		return "在线"
	case MSDelete:
		return "已删除"
	default:
		return "未知"
	}
}

// Minion 节点表
type Minion struct {
	ID         int64        `json:"id,string"        gorm:"column:id;primaryKey"` // 节点 ID
	Inet       string       `json:"inet"             gorm:"column:inet"`          // 节点 IPv4
	Inet6      string       `json:"inet6"            gorm:"column:inet6"`         // 节点 IPv6
	MAC        string       `json:"mac"              gorm:"column:mac"`           // 节点 MAC 地址
	Goos       string       `json:"goos"             gorm:"column:goos"`          // 节点操作系统 runtime.GOOS
	Arch       string       `json:"arch"             gorm:"column:arch"`          // 节点架构 runtime.GOARCH
	Edition    string       `json:"edition"          gorm:"column:edition"`       // 节点当前运行的版本
	Status     MinionStatus `json:"status"           gorm:"column:status"`        // 1-未激活 2-已激活(离线) 3-在线 4-已删除
	Uptime     sql.NullTime `json:"uptime"           gorm:"column:uptime"`        // 上线时间
	BrokerID   int64        `json:"broker_id,string" gorm:"column:broker_id"`     // 上线所在 broker 节点 ID
	BrokerName string       `json:"broker_name"      gorm:"column:broker_name"`   // broker 节点名字
	OrgPath    string       `json:"org_path"         gorm:"column:org_path"`      // 部门路径
	Identity   string       `json:"identity"         gorm:"column:identity"`      // 堡垒机用户
	Category   string       `json:"category"         gorm:"column:category"`      // 部门信息
	OpDuty     string       `json:"op_duty"          gorm:"column:op_duty"`       // 运维负责人
	Comment    string       `json:"comment"          gorm:"column:comment"`       // 说明
	IBu        string       `json:"ibu"              gorm:"column:ibu"`           // 部门
	IDC        string       `json:"idc"              gorm:"column:idc"`           // IDC机房
	CreatedAt  time.Time    `json:"created_at"       gorm:"column:created_at"`    // 创建时间
	UpdatedAt  time.Time    `json:"updated_at"       gorm:"column:updated_at"`    // 修改时间
}

// TableName implement gorm schema.Tabler
func (Minion) TableName() string {
	return "minion"
}

func (m Minion) Invalid() bool {
	return m.Status != MSOffline && m.Status != MSOnline
}

// Minions []*Minion
type Minions []*Minion

// BrokerMap 整理为 key: brokerID; value: minionIDs
func (ms Minions) BrokerMap() map[int64][]int64 {
	ret := make(map[int64][]int64, 16)
	for _, m := range ms {
		minionIDs := ret[m.BrokerID]
		if minionIDs == nil {
			ss := make([]int64, 0, 32)
			ss = append(ss, m.ID)
			ret[m.BrokerID] = ss
		} else {
			ret[m.BrokerID] = append(minionIDs, m.ID)
		}
	}
	return ret
}

//import (
//	"database/sql"
//	"time"
//)
//
//// Minion 节点信息表
//type Minion struct {
//	ID         int64        `json:"id,string"        gorm:"column:id;primaryKey"` // ID
//	Name       string       `json:"name"             gorm:"column:name"`          // 节点名字
//	Inet       string       `json:"inet"             gorm:"column:inet"`          // IPv4
//	Inet6      string       `json:"inet6"            gorm:"column:inet6"`         // IPv6
//	Status     MinionStatus `json:"status"           gorm:"column:status"`        // 节点状态
//	MAC        string       `json:"mac"              gorm:"column:mac"`           // MAC
//	Goos       string       `json:"goos"             gorm:"column:goos"`          // Goos
//	Arch       string       `json:"arch"             gorm:"column:arch"`          // Arch
//	Semver     string       `json:"semver"           gorm:"column:semver"`        // 版本号
//	CPU        int          `json:"cpu"              gorm:"column:cpu"`           // CPU 核心数
//	PID        int          `json:"pid"              gorm:"column:pid"`           // 进程 PID
//	Username   string       `json:"username"         gorm:"column:username"`      // 运行 agent 程序的 用户
//	Hostname   string       `json:"hostname"         gorm:"column:hostname"`      // 主机名
//	Workdir    string       `json:"workdir"          gorm:"column:workdir"`       // 工作目录
//	Executable string       `json:"executable"       gorm:"column:executable"`    // 执行路径
//	PingedAt   sql.NullTime `json:"pinged_at"        gorm:"column:pinged_at"`     // 最近一次 ping 的时间
//	JoinedAt   sql.NullTime `json:"joined_at"        gorm:"column:joined_at"`     // 最近一次加入（连接）时间
//	BrokerID   int64        `json:"broker_id,string" gorm:"column:broker_id"`     // 接入的代理节点 ID
//	BrokerName string       `json:"broker_name"      gorm:"column:broker_name"`   // 接入的代理节点名字
//	CreatedAt  time.Time    `json:"created_at"       gorm:"column:created_at"`    // 创建时间
//	UpdatedAt  time.Time    `json:"updated_at"       gorm:"column:updated_at"`    // 更新时间
//}
//
//// TableName gorm table name
//func (Minion) TableName() string {
//	return "minion"
//}
//
//type MinionStatus uint8
//
//const (
//	MinionInactive MinionStatus = iota // 未激活
//	MinionOffline                      // 离线
//	MinionOnline                       // 在线
//	MinionRemove                       // 移除
//)
//
//// Minions []*Minion
//type Minions []*Minion
//
//// BrokerMap 整理为 key: brokerID; value: minionIDs
//func (ms Minions) BrokerMap() map[int64][]int64 {
//	ret := make(map[int64][]int64, 16)
//	for _, m := range ms {
//		minionIDs := ret[m.BrokerID]
//		if minionIDs == nil {
//			ss := make([]int64, 0, 32)
//			ss = append(ss, m.ID)
//			ret[m.BrokerID] = ss
//		} else {
//			ret[m.BrokerID] = append(minionIDs, m.ID)
//		}
//	}
//	return ret
//}
