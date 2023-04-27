package dynsql

type TimeColumnBuilder interface {
	Operators(ops []Operator) TimeColumnBuilder
	Build() Column
}

func TimeColumn(col, name string) TimeColumnBuilder {
	return &timeColumnBuilder{
		col:  col,
		name: name,
	}
}

type timeColumnBuilder struct {
	col  string
	name string
	ops  []Operator
}

func (tc *timeColumnBuilder) Operators(ops []Operator) TimeColumnBuilder {
	tc.ops = ops
	return tc
}

func (tc *timeColumnBuilder) Build() Column {
	if len(tc.ops) == 0 {
		tc.ops = []Operator{Eq, Ne, Gt, Lt, Gte, Lte}
	}
	opMap := make(map[string]Operator, len(tc.ops))
	for _, op := range tc.ops {
		val := op.value()
		opMap[val] = op
	}

	col := &tableColumn{
		tp:     typeTime,
		column: tc.col,
		name:   tc.name,
		ops:    tc.ops,
		opMap:  opMap,
	}

	return col
}
