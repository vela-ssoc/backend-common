package model

import "time"

// SBOMVuln 漏洞库
type SBOMVuln struct {
	ID          int64     `json:"id,string"   gorm:"column:id;primaryKey"`   // ID
	VulnID      string    `json:"vuln_id"     gorm:"column:vuln_id"`         // 漏洞编号
	PURL        string    `json:"purl"        gorm:"column:purl"`            // PURL
	Title       string    `json:"title"       gorm:"column:title"`           // 漏洞标题
	Description string    `json:"description" gorm:"column:description"`     // 漏洞简介
	Score       CVSSScore `json:"score"       gorm:"column:score"`           // 漏洞分数
	Level       CVSSLevel `json:"level"       gorm:"column:level"`           // 漏洞级别
	Vector      string    `json:"vector"      gorm:"column:vector"`          // CVSS Vector
	CVE         string    `json:"cve"         gorm:"column:cve"`             // CVE
	CWE         string    `json:"cwe"         gorm:"column:cwe"`             // CWE
	Reference   string    `json:"reference"   gorm:"column:reference"`       // reference
	References  []string  `json:"references"  gorm:"column:references;json"` // External References
	Nonce       int64     `json:"-"           gorm:"column:nonce"`           // 批次 ID
	CreatedAt   time.Time `json:"created_at"  gorm:"column:created_at"`      // 创建时间
	UpdatedAt   time.Time `json:"updated_at"  gorm:"column:updated_at"`      // 更新时间
}

// TableName implement gorm schema.Tabler
func (SBOMVuln) TableName() string {
	return "sbom_vuln"
}
