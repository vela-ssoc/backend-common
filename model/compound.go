package model

import "time"

// Compound 配置组合映射表
type Compound struct {
	ID         int64     `json:"id,string"         gorm:"column:id;primaryKey"`   // ID
	Name       string    `json:"name"              gorm:"column:name"`            // 名字
	Desc       string    `json:"desc"              gorm:"column:desc"`            // 说明
	Substances []int64   `json:"substances"        gorm:"column:substances;json"` // 配置项
	Version    int64     `json:"version"           gorm:"column:version"`         // 乐观锁
	CreatedID  int64     `json:"created_id,string" gorm:"column:created_id"`      // 创建人 ID
	UpdatedID  int64     `json:"updated_id,string" gorm:"column:updated_id"`      // 最后更新者 ID
	CreatedAt  time.Time `json:"created_at"        gorm:"column:created_at"`      // 创建时间
	UpdatedAt  time.Time `json:"updated_at"        gorm:"column:updated_at"`      // 最后更新时间
}

// TableName implement gorm schema.Tabler
func (Compound) TableName() string {
	return "compound"
}

type Compounds []*Compound

// SubstanceIDs 获取所有的 substance id
func (cs Compounds) SubstanceIDs() []int64 {
	hm := make(map[int64]struct{}, 32)
	ret := make([]int64, 0, 32)

	for _, c := range cs {
		for _, subID := range c.Substances {
			if _, exist := hm[subID]; exist {
				continue
			}
			hm[subID] = struct{}{}
			ret = append(ret, subID)
		}
	}

	return ret
}
