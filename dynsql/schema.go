package dynsql

type Schema struct {
	Filters columnSchemas `json:"filters"`
	Groups  nameSchemas   `json:"groups"`
	Orders  nameSchemas   `json:"orders"`
}

type nameSchema struct {
	Col  string `json:"col"`
	Name string `json:"name"`
}

type nameSchemas []*nameSchema

type columnSchema struct {
	Col       string            `json:"col"`
	Name      string            `json:"name"`
	Type      string            `json:"type"`
	Operators []*operatorSchema `json:"operators"`
	Enums     []*enumItem       `json:"enums"`
}

type columnSchemas []*columnSchema
