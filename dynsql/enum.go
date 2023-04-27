package dynsql

import "strconv"

type enumItem struct {
	Val  any    `json:"val"`
	Name string `json:"name"`
}

type Enums struct {
	items []*enumItem
}

func (es Enums) toMap() map[string]struct{} {
	size := len(es.items)
	if size == 0 {
		return nil
	}

	ret := make(map[string]struct{}, size)
	for _, e := range es.items {
		switch elem := e.Val.(type) {
		case string:
			ret[elem] = struct{}{}
		case int:
			str := strconv.FormatInt(int64(elem), 10)
			ret[str] = struct{}{}
		case bool:
			str := strconv.FormatBool(elem)
			ret[str] = struct{}{}
		default: // 枚举仅支持 string int bool 三种类型
			return nil
		}
	}

	return ret
}
