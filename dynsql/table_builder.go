package dynsql

type TableBuilder interface {
	Filters(...Column) TableBuilder
	Groups(...Column) TableBuilder
	Orders(...Column) TableBuilder
	Build() Table
}

func Builder() TableBuilder {
	return &tableBuilder{}
}

type tableBuilder struct {
	filters []Column
	orders  []Column
	groups  []Column
}

func (tb *tableBuilder) Filters(cs ...Column) TableBuilder {
	tb.filters = append(tb.filters, cs...)
	return tb
}

func (tb *tableBuilder) Groups(cs ...Column) TableBuilder {
	tb.groups = append(tb.groups, cs...)
	return tb
}

func (tb *tableBuilder) Orders(cs ...Column) TableBuilder {
	tb.orders = append(tb.orders, cs...)
	return tb
}

func (tb *tableBuilder) Build() Table {
	fsz, gsz, osz := len(tb.filters), len(tb.groups), len(tb.orders)
	filterMap := make(map[string]Column, fsz)
	groupMap := make(map[string]struct{}, gsz)
	orderMap := make(map[string]struct{}, osz)
	filters := make(columnSchemas, 0, fsz)
	groups := make(nameSchemas, 0, gsz)
	orders := make(nameSchemas, 0, osz)

	for _, c := range tb.filters {
		cn := c.columnName()
		filterMap[cn] = c
		filters = append(filters, c.columnSchema())
	}
	for _, o := range tb.orders {
		cn := o.columnName()
		orderMap[cn] = struct{}{}
		orders = append(orders, o.nameSchema())
	}
	for _, g := range tb.groups {
		cn := g.columnName()
		groupMap[cn] = struct{}{}
		groups = append(groups, g.nameSchema())
	}

	return &tableEnv{
		filterMap: filterMap,
		groupMap:  groupMap,
		orderMap:  orderMap,
		schema: Schema{
			Filters: filters,
			Groups:  groups,
			Orders:  orders,
		},
	}
}
