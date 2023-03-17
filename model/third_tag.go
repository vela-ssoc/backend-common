package model

// ThirdTag 三方文件标签表
type ThirdTag struct {
	ID      int64  `json:"id,string"       gorm:"column:id;primaryKey"` // 数据库 ID，对于业务没有意义
	Tag     string `json:"tag"             gorm:"column:tag"`           // 标签
	ThirdID int64  `json:"third_id,string" gorm:"column:third_id"`      // minion 节点 ID
}

// TableName implement gorm schema.Tabler
func (ThirdTag) TableName() string {
	return "third_tag"
}

type ThirdTags []*ThirdTag

func (ts ThirdTags) Map() map[int64][]string {
	ret := make(map[int64][]string, 16)
	for _, tt := range ts {
		ss := ret[tt.ThirdID]
		ret[tt.ThirdID] = append(ss, tt.Tag)
	}
	return ret
}
