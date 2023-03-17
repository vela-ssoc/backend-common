package opencond

import (
	"fmt"
	"strings"

	"gorm.io/gorm/clause"
)

var (
	OpEq      = &opEq{}
	OpNe      = &opNe{}
	OpGt      = &opGt{}
	OpLt      = &opLt{}
	OpGte     = &opGte{}
	OpLte     = &opLte{}
	OpIn      = &opIn{}
	OpNotIn   = &opNotIn{}
	OpLike    = &opLike{}
	OpNotLike = &opNotLike{}
)

type operator interface {
	fmt.Stringer
	desc() string
	display() string
	interpreter(*pattern, *environment) (clause.Expression, error)
}

type opEq struct{}

func (op *opEq) String() string {
	return "eq"
}

func (op *opEq) desc() string {
	return "等于"
}

func (op *opEq) display() string {
	return "="
}

func (op *opEq) interpreter(ste *pattern, env *environment) (clause.Expression, error) {
	raw := env.Value
	if err := ste.inEnums(raw); err != nil { // 检查是否符合 enum
		return nil, err
	}
	val, err := ste.datatype.cast(raw) // 类型转换
	if err != nil {
		return nil, err
	}

	return clause.Eq{Column: ste.column, Value: val}, nil
}

type opNe struct{}

func (op *opNe) String() string {
	return "ne"
}

func (op *opNe) desc() string {
	return "不等于"
}

func (op *opNe) display() string {
	return "!="
}

func (op *opNe) interpreter(ste *pattern, env *environment) (clause.Expression, error) {
	raw := env.Value
	if err := ste.inEnums(raw); err != nil { // 检查是否符合 enum
		return nil, err
	}
	val, err := ste.datatype.cast(raw) // 类型转换
	if err != nil {
		return nil, err
	}

	return clause.Not(clause.Eq{Column: ste.column, Value: val}), nil
}

type opGt struct{}

func (op *opGt) String() string {
	return "gt"
}

func (op *opGt) desc() string {
	return "大于"
}

func (op *opGt) display() string {
	return ">"
}

func (op *opGt) interpreter(ste *pattern, env *environment) (clause.Expression, error) {
	raw := env.Value
	if err := ste.inEnums(raw); err != nil { // 检查是否符合 enum
		return nil, err
	}
	val, err := ste.datatype.cast(raw) // 类型转换
	if err != nil {
		return nil, err
	}

	return clause.Gt{Column: ste.column, Value: val}, nil
}

type opLt struct{}

func (op *opLt) String() string {
	return "lt"
}

func (op *opLt) desc() string {
	return "小于"
}

func (op *opLt) display() string {
	return "<"
}

func (op *opLt) interpreter(ste *pattern, env *environment) (clause.Expression, error) {
	raw := env.Value
	if err := ste.inEnums(raw); err != nil { // 检查是否符合 enum
		return nil, err
	}
	val, err := ste.datatype.cast(raw) // 类型转换
	if err != nil {
		return nil, err
	}

	return clause.Lt{Column: ste.column, Value: val}, nil
}

type opGte struct{}

func (op *opGte) String() string {
	return "gte"
}

func (op *opGte) desc() string {
	return "大于等于"
}

func (op *opGte) display() string {
	return ">="
}

func (op *opGte) interpreter(ste *pattern, env *environment) (clause.Expression, error) {
	raw := env.Value
	if err := ste.inEnums(raw); err != nil { // 检查是否符合 enum
		return nil, err
	}
	val, err := ste.datatype.cast(raw) // 类型转换
	if err != nil {
		return nil, err
	}

	return clause.Gte{Column: ste.column, Value: val}, nil
}

type opLte struct{}

func (op *opLte) String() string {
	return "lte"
}

func (op *opLte) desc() string {
	return "小于等于"
}

func (op *opLte) display() string {
	return "<="
}

func (op *opLte) interpreter(ste *pattern, env *environment) (clause.Expression, error) {
	raw := env.Value
	if err := ste.inEnums(raw); err != nil { // 检查是否符合 enum
		return nil, err
	}
	val, err := ste.datatype.cast(raw) // 类型转换
	if err != nil {
		return nil, err
	}

	return clause.Lte{Column: ste.column, Value: val}, nil
}

type opIn struct{}

func (op *opIn) String() string {
	return "in"
}

func (op *opIn) desc() string {
	return "IN"
}

func (op *opIn) display() string {
	return "IN"
}

func (op *opIn) interpreter(ste *pattern, env *environment) (clause.Expression, error) {
	raw := env.Value
	split := strings.Split(raw, ",")
	dats := make([]any, 0, len(split))
	for _, str := range split {
		if err := ste.inEnums(str); err != nil { // 检查是否符合 enum
			return nil, err
		}
		val, err := ste.datatype.cast(str) // 类型转换
		if err != nil {
			return nil, err
		}
		dats = append(dats, val)
	}

	return clause.IN{Column: ste.column, Values: dats}, nil
}

type opNotIn struct{}

func (op *opNotIn) String() string {
	return "notin"
}

func (op *opNotIn) desc() string {
	return "NOT IN"
}

func (op *opNotIn) display() string {
	return "IN"
}

func (op *opNotIn) interpreter(ste *pattern, env *environment) (clause.Expression, error) {
	raw := env.Value
	split := strings.Split(raw, ",")
	dats := make([]any, 0, len(split))
	for _, str := range split {
		if err := ste.inEnums(str); err != nil { // 检查是否符合 enum
			return nil, err
		}
		val, err := ste.datatype.cast(str) // 类型转换
		if err != nil {
			return nil, err
		}
		dats = append(dats, val)
	}

	return clause.Not(clause.IN{Column: ste.column, Values: dats}), nil
}

type opLike struct{}

func (op *opLike) String() string {
	return "like"
}

func (op *opLike) desc() string {
	return "LIKE"
}

func (op *opLike) display() string {
	return "LIKE"
}

func (op *opLike) interpreter(ste *pattern, env *environment) (clause.Expression, error) {
	raw := env.Value
	if !strings.ContainsAny(raw, "%_") {
		raw = "%" + raw + "%"
	}

	return clause.Like{Column: ste.column, Value: raw}, nil
}

type opNotLike struct{}

func (op *opNotLike) String() string {
	return "notlike"
}

func (op *opNotLike) desc() string {
	return "NOT LIKE"
}

func (op *opNotLike) display() string {
	return "NOT LIKE"
}

func (op *opNotLike) interpreter(ste *pattern, env *environment) (clause.Expression, error) {
	raw := env.Value
	if !strings.ContainsAny(raw, "%_") {
		raw = "%" + raw + "%"
	}

	return clause.Not(clause.Like{Column: ste.column, Value: raw}), nil
}
