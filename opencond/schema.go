package opencond

type pair struct {
	Key  string `json:"key"`  // 值
	Desc string `json:"desc"` // 描述
}

type Pairs []*pair

type schema struct {
	Key       string `json:"key"`
	Desc      string `json:"desc"`
	Type      string `json:"type"`
	Operators Pairs  `json:"operators"`
	Enum      bool   `json:"enum"`
	Enums     Pairs  `json:"enums,omitempty"`
}

type Schemas []*schema
