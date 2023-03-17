package model

import "time"

// VIP Virtual IP Address
type VIP struct {
	ID          int64     `json:"id,string"    gorm:"column:id;primaryKey"` // ID
	VirtualIP   string    `json:"virtual_ip"   gorm:"column:virtual_ip"`    // 虚拟 IP
	VirtualPort int       `json:"virtual_port" gorm:"column:virtual_port"`  // 虚拟 IP 开放的端口
	VirtualAddr string    `json:"virtual_addr" gorm:"column:virtual_addr"`  // 虚拟地址 virtual_ip + : + virtual_port
	Enable      bool      `json:"enable"       gorm:"column:enable"`        // 是否开启
	IDC         string    `json:"idc"          gorm:"column:idc"`           // IDC
	MemberIP    string    `json:"member_ip"    gorm:"column:member_ip"`     // 内部成员 IP
	MemberPort  int       `json:"member_port"  gorm:"column:member_port"`   // 成员 端口
	Status      string    `json:"status"       gorm:"column:status"`        // 状态
	Priority    int       `json:"priority"     gorm:"column:priority"`      // 优先级
	BizBranch   string    `json:"biz_branch"   gorm:"column:biz_branch"`    // 业务分支
	BizDept     string    `json:"biz_dept"     gorm:"column:biz_dept"`      // 业务部门
	BizType     string    `json:"biz_type"     gorm:"column:biz_type"`      // 业务类型
	CreatedAt   time.Time `json:"created_at"   gorm:"column:created_at"`    // 创建时间
	UpdatedAt   time.Time `json:"updated_at"   gorm:"column:updated_at"`    // 更新事件
}

// TableName implement gorm schema.Tabler
func (VIP) TableName() string {
	return "vip"
}

type VIPs []*VIP

func (vips VIPs) Mapping() VIPMembers {
	hm := make(map[string]*vipMapping, 32)
	ret := make(VIPMembers, 0, 32)
	for _, row := range vips {
		addr := row.VirtualAddr
		if mapping := hm[addr]; mapping != nil {
			mb := &memberAddr{MemberIP: row.MemberIP, MemberPort: row.MemberPort}
			mapping.Members = append(mapping.Members, mb)
			continue
		}

		mbs := []*memberAddr{{MemberIP: row.MemberIP, MemberPort: row.MemberPort}}
		mapping := &vipMapping{
			VirtualIP:   row.VirtualIP,
			VirtualPort: row.VirtualPort,
			VirtualAddr: row.VirtualAddr,
			Enable:      row.Enable,
			IDC:         row.IDC,
			Status:      row.Status,
			Priority:    row.Priority,
			BizBranch:   row.BizBranch,
			BizDept:     row.BizDept,
			BizType:     row.BizType,
			Members:     mbs,
		}
		ret = append(ret, mapping)
		hm[addr] = mapping
	}

	return ret
}

type memberAddr struct {
	MemberIP   string `json:"member_ip"`
	MemberPort int    `json:"member_port"`
}

type VIPMembers []*vipMapping

type vipMapping struct {
	VirtualIP   string        `json:"virtual_ip"`   // 虚拟 IP
	VirtualPort int           `json:"virtual_port"` // 虚拟 IP 开放的端口
	VirtualAddr string        `json:"virtual_addr"` // 虚拟地址 virtual_ip + : + virtual_port
	Enable      bool          `json:"enable"`       // 是否开启
	IDC         string        `json:"idc"`          // IDC
	Status      string        `json:"status"`       // 状态
	Priority    int           `json:"priority"`     // 优先级
	BizBranch   string        `json:"biz_branch"`   // 业务分支
	BizDept     string        `json:"biz_dept"`     // 业务部门
	BizType     string        `json:"biz_type"`     // 业务类型
	Members     []*memberAddr `json:"members"`      // 映射内网 IP
}
