package dynsql

import (
	"strings"

	"gorm.io/gorm/clause"
)

type Column interface {
	inter(op, val string) (clause.Expression, error)
	columnName() string
	columnSchema() *columnSchema
	nameSchema() *nameSchema
}

type tableColumn struct {
	tp      columnTyper
	column  string
	name    string
	ops     []Operator
	opMap   map[string]Operator
	enum    Enums
	enumMap map[string]struct{}
}

func (tc *tableColumn) columnName() string {
	return tc.column
}

func (tc *tableColumn) inter(op, val string) (clause.Expression, error) {
	opr, ok := tc.opMap[op]
	if !ok {
		return nil, &Error{name: "运算符", value: op}
	}
	if opr == In || opr == NotIn {
		return tc.splitBatch(opr, val)
	}

	if !tc.passEnum(val) {
		return nil, &Error{name: "枚举值", value: val}
	}

	value, err := tc.tp.cast(val)
	if err != nil || value == nil {
		return nil, err
	}

	exp := opr.expr(tc.column, value)

	return exp, nil
}

func (tc *tableColumn) splitBatch(opr Operator, val string) (clause.Expression, error) {
	sn := strings.Split(val, "|")

	anis := make([]any, 0, len(sn))
	for _, s := range sn {
		if s == "" || strings.TrimSpace(s) == "" {
			continue
		}
		if !tc.passEnum(s) {
			return nil, &Error{name: "枚举值", value: val}
		}
		value, err := tc.tp.cast(s)
		if err != nil {
			return nil, err
		}
		anis = append(anis, value)
	}
	if len(anis) == 0 {
		return nil, nil
	}

	exp := opr.expr(tc.column, anis...)

	return exp, nil
}

func (tc *tableColumn) passEnum(str string) bool {
	if tc.enumMap == nil {
		return true
	}
	_, ok := tc.enumMap[str]

	return ok
}

func (tc *tableColumn) columnSchema() *columnSchema {
	opSchemas := make([]*operatorSchema, 0, len(tc.ops))
	for _, op := range tc.ops {
		opSchemas = append(opSchemas, op.schema())
	}

	items := tc.enum.items
	if items == nil {
		items = []*enumItem{}
	}

	return &columnSchema{
		Col:       tc.column,
		Name:      tc.name,
		Type:      tc.tp.name(),
		Operators: opSchemas,
		Enums:     items,
	}
}

func (tc *tableColumn) nameSchema() *nameSchema {
	return &nameSchema{
		Col:  tc.column,
		Name: tc.name,
	}
}
