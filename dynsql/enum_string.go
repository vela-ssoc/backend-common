package dynsql

type StringEnumBuilder interface {
	Set(val, name string) StringEnumBuilder
	Sames(values []string) StringEnumBuilder
	build() Enums
}

func StringEnum() StringEnumBuilder {
	return &stringEnumBuilder{
		hm:     make(map[string]string, 8),
		orders: make([]string, 0, 10),
	}
}

type stringEnumBuilder struct {
	hm     map[string]string
	orders []string
}

func (sb *stringEnumBuilder) Set(item string, name string) StringEnumBuilder {
	sb.hm[item] = name
	sb.orders = append(sb.orders, item)
	return sb
}

func (sb *stringEnumBuilder) Sames(items []string) StringEnumBuilder {
	for _, item := range items {
		sb.hm[item] = item
		sb.orders = append(sb.orders, item)
	}
	return sb
}

func (sb *stringEnumBuilder) build() Enums {
	enums := make([]*enumItem, 0, len(sb.hm))
	for _, od := range sb.orders {
		name := sb.hm[od]
		enums = append(enums, &enumItem{Val: od, Name: name})
	}

	return Enums{items: enums}
}
