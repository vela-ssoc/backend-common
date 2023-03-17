package model

// Email 邮件发送配置
type Email struct {
	ID       int64  `json:"id,string" gorm:"column:id;primaryKey"` // ID
	Host     string `json:"host"      gorm:"column:host"`          // 邮箱服务器
	Username string `json:"username"  gorm:"column:username"`      // 邮箱账号
	Password string `json:"password"  gorm:"column:password"`      // 密码
	Enable   bool   `json:"enable"    gorm:"column:enable"`        // 当前使用的
}

// TableName implement gorm schema.Tabler
func (Email) TableName() string {
	return "email"
}
