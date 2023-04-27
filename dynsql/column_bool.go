package dynsql

type BoolColumnBuilder interface {
	Operators(ops []Operator) BoolColumnBuilder
	Enums(BoolEnumBuilder) BoolColumnBuilder
	Build() Column
}

func BoolColumn(col, name string) BoolColumnBuilder {
	return &boolColumnBuilder{
		col:  col,
		name: name,
	}
}

type boolColumnBuilder struct {
	col  string
	name string
	ops  []Operator
	eb   BoolEnumBuilder
}

func (bc *boolColumnBuilder) Operators(ops []Operator) BoolColumnBuilder {
	bc.ops = ops
	return bc
}

func (bc *boolColumnBuilder) Enums(eb BoolEnumBuilder) BoolColumnBuilder {
	bc.eb = eb
	return bc
}

func (bc *boolColumnBuilder) Build() Column {
	if len(bc.ops) == 0 {
		bc.ops = []Operator{Eq, Ne}
	}
	opMap := make(map[string]Operator, len(bc.ops))
	for _, op := range bc.ops {
		val := op.value()
		opMap[val] = op
	}

	tc := &tableColumn{
		tp:     typeBool,
		column: bc.col,
		name:   bc.name,
		ops:    bc.ops,
		opMap:  opMap,
	}
	if bc.eb != nil {
		enum := bc.eb.build()
		tc.enum = enum
		tc.enumMap = enum.toMap()
	}

	return tc
}
