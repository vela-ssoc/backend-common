package model

import "time"

// SBOMProject 项目表
type SBOMProject struct {
	ID            int64     `json:"id,string"        gorm:"column:id;primaryKey"`  // 项目 ID
	MinionID      int64     `json:"minion_id,string" gorm:"column:minion_id"`      // 节点 ID
	Inet          string    `json:"inet"             gorm:"column:inet"`           // 节点 IP
	Filepath      string    `json:"filepath"         gorm:"column:filepath"`       // 文件路径
	SHA1          string    `json:"sha1"             gorm:"column:sha1"`           // 文件哈希
	Size          int       `json:"size"             gorm:"column:size"`           // 文件大小
	ComponentNum  int       `json:"component_num"    gorm:"column:component_num"`  // 包含组件个数
	PID           int       `json:"pid"              gorm:"column:pid"`            // 进程 PID
	Exe           string    `json:"exe"              gorm:"column:exe"`            // 执行路径
	Username      string    `json:"username"         gorm:"column:username"`       // 执行用户
	ModifyAt      time.Time `json:"modify_at"        gorm:"column:modify_at"`      // 文件最近修改时间
	CriticalNum   int       `json:"critical_num"     gorm:"column:critical_num"`   // 紧急漏洞个数
	CriticalScore CVSSScore `json:"critical_score"   gorm:"column:critical_score"` // 紧急漏洞总分
	HighNum       int       `json:"high_num"         gorm:"column:high_num"`       // 高危漏洞个数
	HighScore     CVSSScore `json:"high_score"       gorm:"column:high_score"`     // 高危漏洞总分
	MediumNum     int       `json:"medium_num"       gorm:"column:medium_num"`     // 中危漏洞个数
	MediumScore   CVSSScore `json:"medium_score"     gorm:"column:medium_score"`   // 中危漏洞总分
	LowNum        int       `json:"low_num"          gorm:"column:low_num"`        // 低危漏洞个数
	LowScore      CVSSScore `json:"low_score"        gorm:"column:low_score"`      // 低危漏洞总分
	TotalNum      int       `json:"total_num"        gorm:"column:total_num"`      // 漏洞总数
	TotalScore    CVSSScore `json:"total_score"      gorm:"column:total_score"`    // 漏洞总分
	Nonce         int64     `json:"-"                gorm:"column:nonce"`          // 同步批次 ID
	CreatedAt     time.Time `json:"created_at"       gorm:"column:created_at"`     // 创建时间
	UpdatedAt     time.Time `json:"updated_at"       gorm:"column:updated_at"`     // 修改时间
}

// TableName implement schema.Tabler
func (SBOMProject) TableName() string {
	return "sbom_project"
}
