package dynsql

type IntColumnBuilder interface {
	Operators(ops []Operator) IntColumnBuilder
	Enums(IntEnumBuilder) IntColumnBuilder
	Build() Column
}

func IntColumn(col, name string) IntColumnBuilder {
	return &intColumnBuilder{
		col:  col,
		name: name,
	}
}

type intColumnBuilder struct {
	col  string
	name string
	ops  []Operator
	eb   IntEnumBuilder
}

func (ic *intColumnBuilder) Operators(ops []Operator) IntColumnBuilder {
	ic.ops = ops
	return ic
}

func (ic *intColumnBuilder) Enums(eb IntEnumBuilder) IntColumnBuilder {
	ic.eb = eb
	return ic
}

func (ic *intColumnBuilder) Build() Column {
	if len(ic.ops) == 0 {
		ic.ops = []Operator{Eq, Ne, Gt, Lt, Gte, Lte, In, NotIn}
	}
	opMap := make(map[string]Operator, len(ic.ops))
	for _, op := range ic.ops {
		val := op.value()
		opMap[val] = op
	}

	tc := &tableColumn{
		tp:     typeInt,
		column: ic.col,
		name:   ic.name,
		ops:    ic.ops,
		opMap:  opMap,
	}
	if ic.eb != nil {
		enum := ic.eb.build()
		tc.enum = enum
		tc.enumMap = enum.toMap()
	} else {
		tc.enum = Enums{items: []*enumItem{}}
		tc.enumMap = map[string]struct{}{}
	}

	return tc
}
