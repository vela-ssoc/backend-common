package model

type TagKind int8

const (
	TkLifelong TagKind = iota + 1
	TkManual
	TkMinion
)

func (tk TagKind) Lifelong() bool {
	return tk == TkLifelong
}

func (tk TagKind) String() string {
	switch tk {
	case TkLifelong:
		return "系统永久标签"
	case TkManual:
		return "手动添加标签"
	case TkMinion:
		return "节点上报标签"
	default:
		return "未知类型标签"
	}
}

// MinionTag minion 节点和标签的映射关系
type MinionTag struct {
	ID       int64   `json:"id,string"        gorm:"column:id;primaryKey"` // 数据库 ID，对于业务没有意义
	Tag      string  `json:"tag"              gorm:"column:tag"`           // 标签
	MinionID int64   `json:"minion_id,string" gorm:"column:minion_id"`     // minion 节点 ID
	Kind     TagKind `json:"kind"             gorm:"column:kind"`          // 标签类型
}

// TableName implement gorm schema.Tabler
func (MinionTag) TableName() string {
	return "minion_tag"
}

type MinionTags []*MinionTag

// ToMap map[minionID][]minionTag
func (mts MinionTags) ToMap() map[int64][]string {
	ret := make(map[int64][]string, 16)
	for _, mt := range mts {
		tags := ret[mt.MinionID]
		if tags == nil {
			ret[mt.MinionID] = []string{mt.Tag}
			continue
		}
		ret[mt.MinionID] = append(tags, mt.Tag)
	}
	return ret
}

func (mts MinionTags) Map() map[int64]MinionTags {
	ret := make(map[int64]MinionTags, 16)
	for _, mt := range mts {
		mid := mt.MinionID
		tags := ret[mid]
		ret[mid] = append(tags, mt)
	}
	return ret
}

func (mts MinionTags) Distinct() []string {
	hm := make(map[string]struct{}, 16)
	ret := make([]string, 0, 16)
	for _, mt := range mts {
		tag := mt.Tag
		if _, ok := hm[tag]; !ok {
			hm[tag] = struct{}{}
			ret = append(ret, tag)
		}
	}
	return ret
}

func (mts MinionTags) MinionIDs() []int64 {
	size := len(mts)
	ret := make([]int64, 0, size)
	hm := make(map[int64]struct{}, size)

	for _, mt := range mts {
		id := mt.MinionID
		if _, exist := hm[id]; exist {
			continue
		}
		hm[id] = struct{}{}
		ret = append(ret, id)
	}

	return ret
}

func (mts MinionTags) Equal(tags []string) bool {
	size := len(tags)
	if size != len(mts) {
		return false
	}

	hm := make(map[string]struct{}, size)
	for _, tag := range tags {
		hm[tag] = struct{}{}
	}
	for _, mt := range mts {
		delete(hm, mt.Tag)
	}

	return len(hm) == 0
}

func (mts MinionTags) Merge(minionID int64, fulls []string) (MinionTags, []string) {
	oldMap := make(map[string]*MinionTag, 16)
	for _, mt := range mts {
		oldMap[mt.Tag] = mt
	}

	ret := make(MinionTags, 0, 16)
	for _, tag := range fulls {
		if mt, exist := oldMap[tag]; exist {
			if mt.Kind.Lifelong() {
				continue
			} else {
				ret = append(ret, mt)
				delete(oldMap, tag)
			}
		} else {
			ret = append(ret, &MinionTag{Tag: tag, MinionID: minionID, Kind: TkManual})
			continue
		}
	}

	removes := make([]string, 0, 8)
	for _, mt := range oldMap {
		if mt.Kind.Lifelong() {
			ret = append(ret, mt)
		} else {
			removes = append(removes, mt.Tag)
		}
	}

	return ret, removes
}
