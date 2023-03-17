package model

import "time"

// SBOMMinion 节点维度的漏洞统计
type SBOMMinion struct {
	ID            int64     `json:"id,string"      gorm:"column:id;primaryKey"`  // 节点 ID
	Inet          string    `json:"inet"           gorm:"column:inet"`           // 节点 IP
	CriticalNum   int       `json:"critical_num"   gorm:"column:critical_num"`   // 紧急漏洞个数
	CriticalScore CVSSScore `json:"critical_score" gorm:"column:critical_score"` // 紧急漏洞总分
	HighNum       int       `json:"high_num"       gorm:"column:high_num"`       // 高危漏洞个数
	HighScore     CVSSScore `json:"high_score"     gorm:"column:high_score"`     // 高危漏洞总分
	MediumNum     int       `json:"medium_num"     gorm:"column:medium_num"`     // 中危漏洞个数
	MediumScore   CVSSScore `json:"medium_score"   gorm:"column:medium_score"`   // 中危漏洞总分
	LowNum        int       `json:"low_num"        gorm:"column:low_num"`        // 低危漏洞个数
	LowScore      CVSSScore `json:"low_score"      gorm:"column:low_score"`      // 低危漏洞总分
	TotalNum      int       `json:"total_num"      gorm:"column:total_num"`      // 漏洞总数
	TotalScore    CVSSScore `json:"total_score"    gorm:"column:total_score"`    // 漏洞总分
	Nonce         int64     `json:"-"              gorm:"column:nonce"`          // 同步批次 ID
	UpdatedAt     time.Time `json:"updated_at"     gorm:"column:updated_at"`     // 最近一更新时间
}

// TableName implement schema.Tabler
func (SBOMMinion) TableName() string {
	return "sbom_minion"
}

type CVSSLevel uint8

const (
	CVSSNone     CVSSLevel = iota // 0.0
	CVSSLow                       // 0.1-3.9
	CVSSMedium                    // 4.0-6.9
	CVSSHigh                      // 7.0-8.9
	CVSSCritical                  // 9.0-10.0
)

type CVSSScore float64

func (s CVSSScore) Level() CVSSLevel {
	if s >= 9 {
		return CVSSCritical
	}
	if s >= 7 {
		return CVSSHigh
	}
	if s >= 4 {
		return CVSSMedium
	}
	if s > 0 {
		return CVSSLow
	}

	return CVSSNone
}
