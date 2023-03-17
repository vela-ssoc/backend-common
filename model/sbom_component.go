package model

import "time"

// SBOMComponent 组件清单
// PURL 标准: https://github.com/package-url/purl-spec
type SBOMComponent struct {
	ID            int64           `json:"id,string"         gorm:"column:id;primaryKey"`  // 组件 ID
	MinionID      int64           `json:"minion_id,string"  gorm:"column:minion_id"`      // 节点 ID
	Inet          string          `json:"inet"              gorm:"column:inet"`           // 节点 IP
	ProjectID     int64           `json:"project_id,string" gorm:"column:project_id"`     // 项目ID
	Filepath      string          `json:"filepath"          gorm:"column:filepath"`       // 文件路径
	SHA1          string          `json:"sha1"              gorm:"column:sha1"`           // 组件的哈希
	Name          string          `json:"name"              gorm:"column:name"`           // 组件名称
	Version       string          `json:"version"           gorm:"column:version"`        // 组件版本
	Language      string          `json:"language"          gorm:"column:language"`       // 组件编程语言
	Licenses      []string        `json:"licenses"          gorm:"column:licenses;json"`  // 组件许可证
	PURL          string          `json:"purl"              gorm:"column:purl"`           // PURL
	CriticalNum   int             `json:"critical_num"      gorm:"column:critical_num"`   // 紧急漏洞个数
	CriticalScore CVSSScore       `json:"critical_score"    gorm:"column:critical_score"` // 紧急漏洞总分
	HighNum       int             `json:"high_num"          gorm:"column:high_num"`       // 高危漏洞个数
	HighScore     CVSSScore       `json:"high_score"        gorm:"column:high_score"`     // 高危漏洞总分
	MediumNum     int             `json:"medium_num"        gorm:"column:medium_num"`     // 中危漏洞个数
	MediumScore   CVSSScore       `json:"medium_score"      gorm:"column:medium_score"`   // 中危漏洞总分
	LowNum        int             `json:"low_num"           gorm:"column:low_num"`        // 低危漏洞个数
	LowScore      CVSSScore       `json:"low_score"         gorm:"column:low_score"`      // 低危漏洞总分
	TotalNum      int             `json:"total_num"         gorm:"column:total_num"`      // 漏洞总数
	TotalScore    CVSSScore       `json:"total_score"       gorm:"column:total_score"`    // 漏洞总分
	Status        ComponentStatus `json:"status"            gorm:"column:status"`         // 处理状态
	Nonce         int64           `json:"-"                 gorm:"column:nonce"`          // 同步批次 ID
	CreatedAt     time.Time       `json:"created_at"        gorm:"column:created_at"`     // 创建时间
	UpdatedAt     time.Time       `json:"updated_at"        gorm:"column:updated_at"`     // 修改时间
}

// TableName implement schema.Tabler
func (SBOMComponent) TableName() string {
	return "sbom_component"
}

type ComponentStatus uint8
