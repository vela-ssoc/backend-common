package model

import "time"

// Certificate 证书模块
type Certificate struct {
	ID          int64     `json:"id"           gorm:"column:id;primaryKey"`  // ID
	Name        string    `json:"name"         gorm:"column:name"`           // 名字
	Certificate []byte    `json:"certificate"  gorm:"column:certificate"`    // 证书
	PrivateKey  []byte    `json:"private_key"  gorm:"column:private_key"`    // 私钥
	Version     int       `json:"version"      gorm:"column:version"`        // 证书版本
	IssCountry  string    `json:"iss_country"  gorm:"column:iss_country"`    // 颁发者国家
	IssOrg      string    `json:"iss_org"      gorm:"column:iss_org"`        // 颁发者组织
	IssCN       string    `json:"iss_cn"       gorm:"column:iss_cn"`         // 颁发者 Common Name
	IssOrgUnit  string    `json:"iss_org_unit" gorm:"column:iss_org_unit"`   // 颁发者组织单位
	SubCountry  string    `json:"sub_country"  gorm:"column:sub_country"`    // 主题国家
	SubOrg      string    `json:"sub_org"      gorm:"column:sub_org"`        // 主题组织
	SubProvince string    `json:"sub_province" gorm:"column:sub_province"`   // 主题省份
	SubCN       string    `json:"sub_cn"       gorm:"column:sub_cn"`         // 主题 Common Name
	DNSNames    []string  `json:"dns_names"    gorm:"column:dns_names;json"` // 证书 DNS Name
	NotBefore   time.Time `json:"not_before"   gorm:"column:not_before"`     // 证书生效时间
	NotAfter    time.Time `json:"not_after"    gorm:"column:not_after"`      // 证书过期时间
	CreatedAt   time.Time `json:"created_at"   gorm:"column:created_at"`     // 创建时间
	UpdatedAt   time.Time `json:"updated_at"   gorm:"column:updated_at"`     // 更新时间
}

// TableName implemented schema.Tabler
func (Certificate) TableName() string {
	return "certificate"
}
