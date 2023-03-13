package model

import "time"

// Oplog 用户操作日志
type Oplog struct {
	ID         int64         `json:"id,string"      gorm:"column:id;primaryKey"` // 数据库 ID
	UserID     int64         `json:"user_id,string" gorm:"column:user_id"`       // 用户 ID
	Username   string        `json:"username"       gorm:"column:username"`      // 用户名
	Nickname   string        `json:"nickname"       gorm:"column:nickname"`      // 操作用户昵称
	Name       string        `json:"name"           gorm:"column:name"`          // 路由名字
	ClientAddr string        `json:"client_addr"    gorm:"column:client_addr"`   // 客户端地址
	DirectAddr string        `json:"direct_addr"    gorm:"column:direct_addr"`   // 直连地址（可能是反向代理地址）
	Method     string        `json:"method"         gorm:"column:method"`        // method
	Path       string        `json:"path"           gorm:"column:path"`          // 请求路径
	Query      string        `json:"query"          gorm:"column:query"`         // query 参数
	Length     int64         `json:"length"         gorm:"column:length"`        // body 长度
	Content    []byte        `json:"content"        gorm:"column:content"`       // body 内容
	Failed     bool          `json:"failed"         gorm:"column:failed"`        // 是否出错
	Cause      string        `json:"cause"          gorm:"column:cause"`         // 如果操作出现错误，此处为错误原因
	RequestAt  time.Time     `json:"request_at"     gorm:"column:request_at"`    // 请求时间
	Elapsed    time.Duration `json:"elapsed"        gorm:"column:elapsed"`       // 操作耗时
	CreatedAt  time.Time     `json:"created_at"     gorm:"column:created_at"`    // 入库创建时间
}

// TableName implement gorm schema.Tabler
func (Oplog) TableName() string {
	return "oplog"
}
