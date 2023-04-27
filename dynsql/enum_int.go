package dynsql

type IntEnumBuilder interface {
	Set(val int, name string) IntEnumBuilder
	build() Enums
}

func IntEnum() IntEnumBuilder {
	return &intEnumBuilder{
		hm:     make(map[int]string, 8),
		orders: make([]int, 0, 10),
	}
}

type intEnumBuilder struct {
	hm     map[int]string
	orders []int
}

func (ib *intEnumBuilder) Set(item int, name string) IntEnumBuilder {
	ib.hm[item] = name
	ib.orders = append(ib.orders, item)
	return ib
}

func (ib *intEnumBuilder) build() Enums {
	enums := make([]*enumItem, 0, len(ib.hm))
	max := len(ib.orders)
	for i := max - 1; i > -1; i-- {
		val := ib.orders[i]
		if name, ok := ib.hm[val]; ok {
			enums = append(enums, &enumItem{Val: val, Name: name})
		}
	}
	return Enums{items: enums}
}
