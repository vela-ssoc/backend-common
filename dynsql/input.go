package dynsql

import (
	"encoding/json"
	"net/url"
)

type Input struct {
	Filters []*filter `query:"filters"`
	Group   string    `query:"group"`
	Order   string    `query:"order"`
	Desc    bool      `query:"desc"`
}

func (in Input) empty() bool {
	return len(in.Filters) == 0 && in.Order == ""
}

type filter struct {
	Col string `json:"col" validate:"required,lte=50"`
	Op  string `json:"op"  validate:"oneof=eq ne gt lt gte lte in notin like notlike"`
	Val string `json:"val" validate:"lte=100"`
}

func (f *filter) UnmarshalBind(raw string) error {
	data, err := url.QueryUnescape(raw)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data), f)
}
