package opencond

import (
	"strconv"
	"time"
)

var (
	TypeString   = &stringType{}
	TypeInt      = &intType{}
	TypeBool     = &boolType{}
	TypeDatetime = &datetimeType{}
)

type datatype interface {
	// key 类型的名字，不要和其它类型重复
	key() string

	// operators 该类型的默认可用操作符
	operators() []operator

	// cast 将 string 转为实现的类型
	cast(string) (any, error)
}

// stringType string
type stringType struct{}

func (*stringType) key() string {
	return "string"
}

func (*stringType) operators() []operator {
	return []operator{OpEq, OpNe, OpIn, OpNotIn, OpLike, OpNotLike}
}

func (st *stringType) cast(str string) (any, error) {
	return str, nil
}

// intType int
type intType struct{}

func (*intType) key() string {
	return "int"
}

func (*intType) operators() []operator {
	return []operator{OpEq, OpNe, OpGt, OpLt, OpGte, OpLte, OpIn, OpNotIn}
}

func (*intType) cast(str string) (any, error) {
	return strconv.Atoi(str)
}

// intType bool
type boolType struct{}

func (*boolType) key() string {
	return "bool"
}

func (*boolType) operators() []operator {
	return []operator{OpEq, OpNe, OpIn, OpNotIn}
}

func (bt *boolType) cast(str string) (any, error) {
	return strconv.ParseBool(str)
}

// datetimeType datetime
type datetimeType struct{}

func (*datetimeType) key() string {
	return "datetime"
}

func (*datetimeType) operators() []operator {
	return []operator{OpEq, OpNe, OpGt, OpLt, OpGte, OpLte, OpIn, OpNotIn}
}

func (*datetimeType) cast(str string) (any, error) {
	tm := &time.Time{}
	err := tm.UnmarshalText([]byte(str))
	return tm, err
}
