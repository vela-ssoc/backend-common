package dynsql

type BoolEnumBuilder interface {
	True(name string) BoolEnumBuilder
	False(name string) BoolEnumBuilder
	build() Enums
}

func BoolEnum() BoolEnumBuilder {
	return &boolEnumBuilder{}
}

type boolEnumBuilder struct {
	trueName  string
	falseName string
}

func (e *boolEnumBuilder) True(name string) BoolEnumBuilder {
	e.trueName = name
	return e
}

func (e *boolEnumBuilder) False(name string) BoolEnumBuilder {
	e.falseName = name
	return e
}

func (e *boolEnumBuilder) build() Enums {
	return Enums{
		items: []*enumItem{
			{Val: true, Name: e.trueName},
			{Val: false, Name: e.falseName},
		},
	}
}
