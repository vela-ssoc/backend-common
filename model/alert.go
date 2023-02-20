package model

import "time"

// Alert 自动化运维平台告警接口
type Alert struct {
	ID        int64     `json:"id,string"   gorm:"column:id;primaryKey"` // ID
	Name      string    `json:"name"        gorm:"column:name"`          // 名字
	Origin    string    `json:"origin"      gorm:"column:origin"`        // 去告警中心提前申请
	Addr      string    `json:"addr"        gorm:"column:addr"`          // 咚咚服务器
	Host      string    `json:"host"        gorm:"column:host"`          // HTTP Header Host
	Enable    bool      `json:"enable"      gorm:"column:enable"`        // 是否启用
	CreatedAt time.Time `json:"created_at"  gorm:"column:created_at"`    // 创建时间
	UpdatedAt time.Time `json:"updated_at"  gorm:"column:updated_at"`    // 更新时间
}

// TableName implemented gorm schema.Tabler
func (Alert) TableName() string {
	return "alert"
}
