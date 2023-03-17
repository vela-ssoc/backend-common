package model

// Elastic es 代理配置
type Elastic struct {
	ID       int64  `json:"id,string" gorm:"column:id;primaryKey"` // ID
	Host     string `json:"host"      gorm:"column:host"`          // es 地址
	Username string `json:"username"  gorm:"column:username"`      // es 用户名
	Password string `json:"password"  gorm:"column:password"`      // es 密码
	Desc     string `json:"desc"      gorm:"column:desc"`          // 简介
	Enable   bool   `json:"enable"    gorm:"column:enable"`        // 是否启用，最多只能有一个启用
}

// TableName implement gorm schema.Tabler
func (Elastic) TableName() string {
	return "elastic"
}
