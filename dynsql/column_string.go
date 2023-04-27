package dynsql

type StringColumnBuilder interface {
	Operators(ops []Operator) StringColumnBuilder
	Enums(StringEnumBuilder) StringColumnBuilder
	Build() Column
}

func StringColumn(col, name string) StringColumnBuilder {
	return &stringColumnBuilder{
		col:  col,
		name: name,
	}
}

type stringColumnBuilder struct {
	col  string
	name string
	ops  []Operator
	eb   StringEnumBuilder
}

func (sc *stringColumnBuilder) Operators(ops []Operator) StringColumnBuilder {
	sc.ops = ops
	return sc
}

func (sc *stringColumnBuilder) Enums(eb StringEnumBuilder) StringColumnBuilder {
	sc.eb = eb
	return sc
}

func (sc *stringColumnBuilder) Build() Column {
	if len(sc.ops) == 0 {
		sc.ops = []Operator{Eq, Ne, Gt, Lt, Gte, Lte, In, NotIn, Like, NotLike}
	}
	opMap := make(map[string]Operator, len(sc.ops))
	for _, op := range sc.ops {
		val := op.value()
		opMap[val] = op
	}

	tc := &tableColumn{
		tp:     typeString,
		column: sc.col,
		name:   sc.name,
		ops:    sc.ops,
		opMap:  opMap,
	}
	if sc.eb != nil {
		enum := sc.eb.build()
		tc.enum = enum
		tc.enumMap = enum.toMap()
	}

	return tc
}
