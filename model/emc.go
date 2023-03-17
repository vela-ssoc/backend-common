package model

// Emc 咚咚服务号配置 emc 是咚咚软件名
type Emc struct {
	ID      int64  `json:"id,string" gorm:"column:id;primaryKey"` // ID
	Name    string `json:"name"      gorm:"column:name"`          // 名字
	Host    string `json:"host"      gorm:"column:host"`          // 咚咚服务器
	Account string `json:"account"   gorm:"column:account"`       // 咚咚 Account
	Token   string `json:"token"     gorm:"column:token"`         // 咚咚 Token
	Enable  bool   `json:"enable"    gorm:"column:enable"`        // 是否启用
}

// TableName implement gorm schema.Tabler
func (Emc) TableName() string {
	return "emc"
}
