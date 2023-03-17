package model

import "time"

// Effect 配置/组合配置 发布到节点的关系表
type Effect struct {
	ID       int64  `json:"id,string"        gorm:"column:id;primaryKey"` // 表 ID
	SubmitID int64  `json:"submit_id,string" gorm:"column:submit_id"`     // 同一次提交的 SubmitID 一样
	Name     string `json:"name"             gorm:"column:name"`          // 同一次提交的 Name 一样
	Tag      string `json:"tag"              gorm:"column:tag"`           // minion 节点的标签

	// 发布的 ID。与 Compound 结合使用：
	// 如果 Compound 为 true 则该字段存放的是 Substance.ID。代表发布的是单个配置
	// 如果 Compound 为 false 则该字段存放的是 Compound.ID。代表发布的是配置组合
	EffectID  int64     `json:"effect_id,string"  gorm:"column:effect_id"`      // 与 Compound 结合使用
	Compound  bool      `json:"compound"          gorm:"column:compound"`       // 是否是组合
	Enable    bool      `json:"enable"            gorm:"column:enable"`         // 同一次提交的 Enable 一样
	Version   int64     `json:"version"           gorm:"column:version"`        // 同一次提交的 Version 一样
	Exclusion []string  `json:"exclusion"         gorm:"column:exclusion;json"` // 排除的节点 IPv4
	CreatedID int64     `json:"created_id,string" gorm:"column:created_id"`     // 创建者 ID
	UpdatedID int64     `json:"updated_id,string" gorm:"column:updated_id"`     // 更新者 ID
	CreatedAt time.Time `json:"created_at"        gorm:"column:created_at"`     // 创建时间
	UpdatedAt time.Time `json:"updated_at"        gorm:"column:updated_at"`     // 更新时间
}

// TableName implement gorm schema.Tabler
func (Effect) TableName() string {
	return "effect"
}

type Effects []*Effect

func (es Effects) Exclusion(ip string) ([]int64, []int64) {
	chm, shm := make(map[int64]struct{}, 32), make(map[int64]struct{}, 32)
	comIDs, subIDs := make([]int64, 0, 32), make([]int64, 0, 32)

	for _, eff := range es {
		if es.exclusion(eff.Exclusion, ip) {
			continue
		}

		effID, com := eff.EffectID, eff.Compound
		if com {
			if _, exist := chm[effID]; exist {
				continue
			}
			chm[effID] = struct{}{}
			comIDs = append(comIDs, effID)
			continue
		}

		if _, exist := shm[effID]; exist {
			continue
		}
		shm[effID] = struct{}{}
		subIDs = append(subIDs, effID)
	}

	return comIDs, subIDs
}

func (Effects) exclusion(exs []string, ip string) bool {
	for _, ex := range exs {
		if ex == ip {
			return true
		}
	}

	return false
}
