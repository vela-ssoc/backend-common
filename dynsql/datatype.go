package dynsql

import (
	"strconv"
	"time"
)

var (
	typeString columnTyper = &stringColumnType{}
	typeInt    columnTyper = &intColumnType{}
	typeBool   columnTyper = &boolColumnType{}
	typeTime   columnTyper = &timeColumnType{}
)

type columnTyper interface {
	name() string
	cast(string) (any, error)
}

type stringColumnType struct{}

func (s *stringColumnType) name() string                 { return "string" }
func (s *stringColumnType) cast(str string) (any, error) { return str, nil }

type intColumnType struct{}

func (i *intColumnType) name() string                 { return "int" }
func (i *intColumnType) cast(str string) (any, error) { return strconv.ParseInt(str, 10, 64) }

type boolColumnType struct{}

func (b *boolColumnType) name() string                 { return "bool" }
func (b *boolColumnType) cast(str string) (any, error) { return strconv.ParseBool(str) }

type timeColumnType struct{}

func (t *timeColumnType) name() string                 { return "time" }
func (t *timeColumnType) cast(val string) (any, error) { return time.Parse(time.RFC3339, val) }
