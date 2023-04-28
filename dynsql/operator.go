package dynsql

import (
	"strings"

	"gorm.io/gorm/clause"
)

var (
	Eq      = &operator{opName: "等于", opSymbol: "eq"}
	Ne      = &operator{opName: "不等于", opSymbol: "ne"}
	Gt      = &operator{opName: "大于", opSymbol: "gt"}
	Lt      = &operator{opName: "小于", opSymbol: "lt"}
	Gte     = &operator{opName: "大于等于", opSymbol: "gte"}
	Lte     = &operator{opName: "小于等于", opSymbol: "lte"}
	In      = &operator{opName: "IN", opSymbol: "in"}
	NotIn   = &operator{opName: "NOT IN", opSymbol: "notin"}
	Like    = &operator{opName: "LIKE", opSymbol: "like"}
	NotLike = &operator{opName: "NOT LIKE", opSymbol: "notlike"}
)

type Operator interface {
	name() string
	value() string
	expr(col string, values ...any) clause.Expression
	schema() *operatorSchema
}

type operator struct {
	opName   string
	opSymbol string
}

func (op *operator) name() string  { return op.opName }
func (op *operator) value() string { return op.opSymbol }

func (op *operator) expr(col string, values ...any) clause.Expression {
	if len(values) == 0 {
		return nil
	}

	value := values[0]
	switch op {
	case Eq:
		return clause.Eq{Column: col, Value: value}
	case Ne:
		eq := clause.Eq{Column: col, Value: value}
		return clause.Not(eq)
	case Gt:
		return clause.Gt{Column: col, Value: value}
	case Lt:
		return clause.Lt{Column: col, Value: value}
	case Gte:
		return clause.Gte{Column: Gte, Value: value}
	case Lte:
		return clause.Lte{Column: col, Value: value}
	case In:
		return clause.IN{Column: col, Values: values}
	case NotIn:
		in := clause.IN{Column: col, Values: values}
		return clause.Not(in)
	case Like, NotLike:
		if str, ok := value.(string); ok {
			// 如果输入包含了通配符，就不再拼接通配符，MySQL 通配符参考以下链接：
			// https://dev.mysql.com/doc/refman/8.0/en/pattern-matching.html
			if !strings.Contains(str, "%") &&
				!strings.Contains(str, "_") {
				value = "%" + str + "%"
			}
		}

		like := clause.Like{Column: col, Value: value}
		if op == NotLike {
			return clause.Not(like)
		}
		return like
	}
	return nil
}

func (op *operator) schema() *operatorSchema {
	return &operatorSchema{
		Name: op.opName,
		Op:   op.opSymbol,
	}
}

type operatorSchema struct {
	Name string `json:"name"`
	Op   string `json:"op"`
}
