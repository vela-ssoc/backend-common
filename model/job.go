package model

import (
	"encoding/json"
	"time"
)

// Job 任务表
type Job struct {
	ID         int64           `json:"id,string"        gorm:"column:id;primaryKey"` // ID
	PolicyID   int64           `json:"policy_id,string" gorm:"column:policy_id"`     // 策略 ID
	PolicyName string          `json:"policy_name"      gorm:"column:policy_name"`   // 策略名字
	PolicyDesc string          `json:"policy_desc"      gorm:"column:policy_desc"`   // 策略描述
	CodeID     int64           `json:"code_id,string"   gorm:"column:code_id"`       // 代码 ID
	CodeName   string          `json:"code_name"        gorm:"column:code_name"`     // 代码名字
	CodeIcon   []byte          `json:"code_icon"        gorm:"column:code_icon"`     // 代码图标
	CodeDesc   string          `json:"code_desc"        gorm:"column:code_desc"`     // 代码说明
	CodeHash   string          `json:"code_hash"        gorm:"column:code_hash"`     // 代码哈希
	CodeChunk  []byte          `json:"code_chunk"       gorm:"column:code_chunk"`    // 代码内容
	Timeout    int             `json:"timeout"          gorm:"column:timeout"`       // 执行超时秒数
	Parallel   int             `json:"parallel"         gorm:"column:parallel"`      // 并发下发个数
	Tags       []string        `json:"tags"             gorm:"column:tags;json"`     // 下发的节点标签
	Status     JobStatus       `json:"status"           gorm:"column:status"`        // 状态
	Args       json.RawMessage `json:"args"             gorm:"column:args"`          // 参数
	Total      int             `json:"total"            gorm:"column:total"`         // 总共下发的节点
	Failed     int             `json:"failed"           gorm:"column:failed"`        // 下发成功个数
	Success    int             `json:"success"          gorm:"column:success"`       // 下发失败个数
	Nonce      int64           `json:"-"                gorm:"column:nonce"`         // 随意唯一值
	CreatedAt  time.Time       `json:"created_at"       gorm:"column:created_at"`    // 创建时间
	UpdatedAt  time.Time       `json:"updated_at"       gorm:"column:updated_at"`    // 更新时间
}

// TableName implement gorm schema.Tabler
func (Job) TableName() string {
	return "job"
}

const (
	JobRunning JobStatus = iota
	JobPause
	JobAbort
	JobFinish
)

type JobStatus int8

func (js JobStatus) String() string {
	switch js {
	case JobRunning:
		// 任务创建的初始状态
		return "运行中"
	case JobPause:
		// 运行中的任务可以被暂停
		return "已暂停"
	case JobAbort:
		// <最终状态>
		// 运行中的任务或暂停中的任务可以被人工终止
		// 任务一旦终止将不能重新运行
		return "已终止"
	case JobFinish:
		// <最终状态>
		// 任务被后端执行完成后就会进入已完成状态
		return "已完成"
	default:
		return "状态错误"
	}
}
