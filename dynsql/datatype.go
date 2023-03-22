package dynsql

import (
	"strconv"
	"time"
)

var (
	TString = &stringType{}
	TInt    = &intType{}
	TFloat  = &floatType{}
	TBool   = &boolType{}
	TTime   = &timeType{}
)

type Typer interface {
	String() string
	cast(string) (any, error)
}

type stringType struct{}

func (stringType) String() string { return "string" }

func (stringType) cast(val string) (any, error) { return val, nil }

type intType struct{}

func (intType) String() string { return "int" }

func (intType) cast(val string) (any, error) { return strconv.ParseInt(val, 10, 64) }

type floatType struct{}

func (floatType) String() string { return "float" }

func (floatType) cast(val string) (any, error) { return strconv.ParseFloat(val, 64) }

type boolType struct{}

func (boolType) String() string { return "bool" }

func (boolType) cast(val string) (any, error) { return strconv.ParseBool(val) }

type timeType struct{}

func (timeType) String() string { return "time" }

func (timeType) cast(val string) (any, error) { return time.Parse(time.RFC3339, val) }
