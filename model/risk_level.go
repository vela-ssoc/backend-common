package model

import "database/sql/driver"

const (
	RLvlCritical RiskLevel = 100000
	RLvlHigh     RiskLevel = 10000
	RLvlMiddle   RiskLevel = 1000
	RLvlLow      RiskLevel = 100
)

var riskLvlSI = map[string]RiskLevel{
	"紧急": RLvlCritical,
	"高危": RLvlHigh,
	"中危": RLvlMiddle,
	"低危": RLvlLow,
}

var riskLvlIS = map[RiskLevel]string{
	RLvlCritical: "紧急",
	RLvlHigh:     "高危",
	RLvlMiddle:   "中危",
	RLvlLow:      "低危",
}

// RiskLevel 风险级别，支持直接比较
type RiskLevel int

func (rl RiskLevel) Value() (driver.Value, error) {
	str := riskLvlIS[rl]
	return []byte(str), nil
}

func (rl *RiskLevel) Scan(src any) error {
	if bs, ok := src.([]byte); ok {
		lv := riskLvlSI[string(bs)]
		*rl = lv
	}
	return nil
}

func (rl *RiskLevel) UnmarshalText(raw []byte) error {
	n := riskLvlSI[string(raw)]
	*rl = n
	return nil
}

func (rl RiskLevel) MarshalText() ([]byte, error) {
	str := riskLvlIS[rl]
	return []byte(str), nil
}

func (rl RiskLevel) String() string {
	return riskLvlIS[rl]
}
