package model

import "time"

// SysInfo 节点上报的状态统计
type SysInfo struct {
	ID            int64     `json:"-"              gorm:"column:id;primaryKey"` // 节点 ID
	Release       string    `json:"release"        gorm:"column:release"`       // 发行版
	CPUCore       int       `json:"cpu_core"       gorm:"column:cpu_core"`      // CPU 核数
	MemTotal      int       `json:"mem_total"      gorm:"column:mem_total"`     // 内存总大小
	MemFree       int       `json:"mem_free"       gorm:"column:mem_free"`      // 空闲内存
	SwapTotal     int       `json:"swap_total"     gorm:"column:swap_total"`    // 交换分区总大小
	SwapFree      int       `json:"swap_free"      gorm:"column:swap_free"`     // 空闲交换分区
	HostID        string    `json:"host_id"        gorm:"column:host_id"`
	Family        string    `json:"family"         gorm:"column:family"`
	Uptime        int64     `json:"uptime"         gorm:"column:uptime"`
	BootAt        int64     `json:"boot_at"        gorm:"column:boot_at"`
	Virtual       string    `json:"virtual"        gorm:"column:virtual_sys"` // MySQL 不允许 virtual 作为字段
	VirtualRole   string    `json:"virtual_role"   gorm:"column:virtual_role"`
	ProcNumber    int       `json:"proc_number"    gorm:"column:proc_number"`
	Hostname      string    `json:"hostname"       gorm:"column:hostname"`
	CPUModel      string    `json:"cpu_model"      gorm:"column:cpu_model"`
	AgentTotal    int       `json:"agent_total"    gorm:"column:agent_total"`
	AgentAlloc    int       `json:"agent_alloc"    gorm:"column:agent_alloc"`
	KernelVersion string    `json:"kernel_version" gorm:"column:kernel_version"`
	UpdatedAt     time.Time `json:"-"              gorm:"column:updated_at"` // 更新时间
}

func (SysInfo) TableName() string {
	return "sys_info"
}
