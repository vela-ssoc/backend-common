package dynsql

type Column struct {
	Key   string
	Col   string
	Name  string
	Type  Typer
	Enums map[string]string
}

type Enum struct {
	Value string
	Desc  string
}

type EnumInt struct {
	Value int
	Desc  string
}

type Builder interface {
	// Enum 枚举
	Enum(enums []*Enum) *Patter

	// EnumInt 枚举
	EnumInt(enums []*EnumInt) *Patter
}

func BCol(key, col, name string, typ Typer) Builder {
	return nil
}

type columnBuild struct {
	key   string
	col   string
	name  string
	typ   Typer
	ops   []Operator
	enums map[string]string
}

type Patter interface {
	pattern()
}
