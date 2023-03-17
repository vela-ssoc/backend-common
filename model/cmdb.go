package model

import "time"

// Cmdb 对应自动化运维CMDB接口, 接口文档请访问: http://yunwei.eastmoney.com/docs/cmdb_api
// 各个字段类型含义请访问: http://yunwei.eastmoney.com:81/api/v0.1/attributes/citype/8
type Cmdb struct {
	ID               int64     `json:"id,string"          gorm:"column:id;primaryKey"`      // 本系统的节点ID
	Inet             string    `json:"inet"               gorm:"column:inet"`               // 本系统的IPv4
	CmdbID           int       `json:"_id"                gorm:"column:_id"`                // 接口未说明(应该是CMDB的ID)
	Org              int       `json:"org"                gorm:"column:_org"`               // 接口未说明(应该是所属部门, 如: 信息安全)
	OrgPath          string    `json:"org_path"           gorm:"column:_org_path"`          // 接口未说明(应该是所属部门全路径, 如: /运维中心/信息安全)
	Type             int       `json:"type"               gorm:"column:_type"`              // 接口未说明
	AgentNotCheck    string    `json:"agent_not_check"    gorm:"column:agent_not_check"`    // agent巡检白名单
	AgentVersion     int       `json:"agent_version"      gorm:"column:agent_version"`      // 客户端版本
	AppName          string    `json:"appname"            gorm:"column:appname"`            // 应用名
	Area             string    `json:"area"               gorm:"column:area"`               // 地域
	BaoleijiIdentity string    `json:"baoleiji_identity"  gorm:"column:baoleiji_identity"`  // 堡垒机可登陆账号
	BusinessScope    string    `json:"business_scope"     gorm:"column:business_scope"`     // 业务作用域
	Category         string    `json:"category"           gorm:"column:category"`           // 业务类型
	CategoryBranch   string    `json:"category_branch"    gorm:"column:category_branch"`    // 业务分支
	CategoryZone     int       `json:"category_zone"      gorm:"column:category_zone"`      // 分区
	CiType           string    `json:"ci_type"            gorm:"column:ci_type"`            // 接口未说明
	CmcIP            string    `json:"cmc_ip"             gorm:"column:cmc_ip"`             // 移动IP
	CncIP            string    `json:"cnc_ip"             gorm:"column:cnc_ip"`             // 网通IP
	Comment          string    `json:"comment"            gorm:"column:comment"`            // 作用
	CostBu           string    `json:"cost_bu"            gorm:"column:cost_bu"`            // 成本所属事业部
	CPU              string    `json:"cpu"                gorm:"column:cpu"`                // CPU型号
	CPUCount         int       `json:"cpu_count"          gorm:"column:cpu_count"`          // CPU数
	CreatedTime      string    `json:"created_time"       gorm:"column:created_time"`       // CMDB创建时间
	CtcIP            string    `json:"ctc_ip"             gorm:"column:ctc_ip"`             // 电信IP
	Env              string    `json:"env"                gorm:"column:env"`                // 环境
	FloatIP          []string  `json:"float_ip"           gorm:"column:float_ip;json"`      // 浮动IP
	HardDisk         string    `json:"harddisk"           gorm:"column:harddisk"`           // 硬盘信息
	HostIP           string    `json:"host_ip"            gorm:"column:host_ip"`            // 宿主机IP
	Hostname         string    `json:"hostname"           gorm:"column:hostname"`           // 主机名
	IBu              string    `json:"i_bu"               gorm:"column:ibu"`                // 事业部
	IDC              string    `json:"idc"                gorm:"column:idc"`                // IDC
	IPv6             []string  `json:"ipv6"               gorm:"column:ipv6;json"`          // IPv6
	KernelVersion    string    `json:"kernel_version"     gorm:"column:kernel_version"`     // 内核版本
	MinionNotCheck   string    `json:"minion_not_check"   gorm:"column:minion_not_check"`   // minion巡检白名单
	NetOpen          string    `json:"net_open"           gorm:"column:net_open"`           // 网络开放状态: 公网 仅内网
	NicIP            string    `json:"nic_ip"             gorm:"column:nic_ip"`             // 网卡IP
	NicMAC           string    `json:"nic_mac"            gorm:"column:nic_mac"`            // 网卡Mac
	OpDuty           string    `json:"op_duty"            gorm:"column:op_duty"`            // 运维负责人
	OsVersion        string    `json:"os_version"         gorm:"column:os_version"`         // 操作系统版本
	PrivateCloudType string    `json:"private_cloud_type" gorm:"column:private_cloud_type"` // 私有云类型
	PrivateIP        []string  `json:"private_ip"         gorm:"column:private_ip;json"`    // 内网IP
	Rack             string    `json:"rack"               gorm:"column:rack"`               // 机架位置
	RAMSize          string    `json:"ram_size"           gorm:"column:ram_size"`           // 内存大小
	RdDuty           string    `json:"rd_duty"            gorm:"column:rd_duty"`            // 开发负责人
	SecurityInfo     string    `json:"security_info"      gorm:"column:security_info"`      // 安全信息
	SecurityRisk     int       `json:"security_risk"      gorm:"column:security_risk"`      // 安全风险值
	ServerRoom       string    `json:"server_room"        gorm:"column:server_room"`        // 机房
	ServerSN         string    `json:"server_sn"          gorm:"column:server_sn"`          // 宿主机序列号
	SSHPort          int       `json:"ssh_port"           gorm:"column:ssh_port"`           // ssh端口
	Status           string    `json:"status"             gorm:"column:status"`             // 状态
	SysDuty          string    `json:"sys_duty"           gorm:"column:sys_duty"`           // 系统负责人
	Unique           string    `json:"unique"             gorm:"column:unique"`             // 接口未说明
	UUID             string    `json:"uuid"               gorm:"column:uuid"`               // UUID
	VServerType      string    `json:"vserver_type"       gorm:"column:vserver_type"`       // 虚拟机类型
	ZabbixNotCheck   string    `json:"zabbix_not_check"   gorm:"column:zabbix_not_check"`   // zabbix巡检白名单
	CreatedAt        time.Time `json:"created_at"         gorm:"column:created_at"`         // 入库时间
	UpdatedAt        time.Time `json:"updated_at"         gorm:"column:updated_at"`         // 入库时间
}

// TableName implement gorm schema.Tabler
func (Cmdb) TableName() string {
	return "cmdb"
}

type Cmdbs []*Cmdb

func (cs Cmdbs) InetMap() map[string]*Cmdb {
	hm := make(map[string]*Cmdb, len(cs))
	for _, c := range cs {
		for _, inet := range c.PrivateIP {
			hm[inet] = c
		}
	}

	return hm
}
