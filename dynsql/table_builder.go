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
	for i := fsz - 1; i > -1; i-- {
		c := tb.filters[i]
		cn := c.columnName()
		if _, exist := filterMap[cn]; !exist {
			filterMap[cn] = c
			filters = append(filters, c.columnSchema())
		}
	}
	for i := osz - 1; i > -1; i-- {
		o := tb.orders[i]
		cn := o.columnName()
		if _, exist := orderMap[cn]; !exist {
			orderMap[cn] = struct{}{}
			orders = append(orders, o.nameSchema())
		}
	}
	for i := gsz - 1; i > -1; i-- {
		o := tb.groups[i]
		cn := o.columnName()
		if _, exist := groupMap[cn]; !exist {
			groupMap[cn] = struct{}{}
			groups = append(groups, o.nameSchema())
		}
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
