package dynsql

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Table interface {
	// Schema 规范
	Schema() Schema

	// Inter Intercept
	Inter(Input) (Scope, error)
}

type Error struct {
	name  string
	value string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s %s 不存在", e.name, e.value)
}

type tableEnv struct {
	filterMap map[string]Column
	groupMap  map[string]struct{}
	orderMap  map[string]struct{}
	schema    Schema
}

func (tbl *tableEnv) Schema() Schema {
	return tbl.schema
}

func (tbl *tableEnv) Inter(input Input) (Scope, error) {
	if input.empty() || (len(tbl.orderMap) == 0 && len(tbl.filterMap) == 0) {
		return nil, nil
	}

	filters, group, order := input.Filters, input.Group, input.Order
	items := make([]clause.Expression, 0, len(filters))
	for _, f := range filters {
		col, ok := tbl.filterMap[f.Col]
		if !ok {
			return nil, &Error{name: "列名", value: f.Col}
		}
		item, err := col.inter(f.Op, f.Val)
		if err != nil {
			return nil, err
		}
		if item != nil {
			items = append(items, item)
		}
	}

	ret := new(scope)
	if len(items) != 0 {
		ret.where = clause.And(items...)
	}
	if order != "" && len(tbl.orderMap) != 0 {
		if _, exist := tbl.orderMap[order]; !exist {
			return nil, &Error{name: "排序条件", value: order}
		}
		ret.orderBy = order
		ret.desc = input.Desc
	}
	if group != "" && len(tbl.groupMap) != 0 {
		if _, exist := tbl.groupMap[group]; !exist {
			return nil, &Error{name: "分组条件", value: order}
		}
		ret.groupBy = group
	}

	return ret, nil
}

type Scope interface {
	Scope(*gorm.DB) *gorm.DB
}

type scope struct {
	where   clause.Expression
	groupBy string
	orderBy string
	desc    bool
}

func (sc *scope) Scope(db *gorm.DB) *gorm.DB {
	if w := sc.where; w != nil {
		db.Where(w)
	}
	if g := sc.groupBy; g != "" {
		db.Group(g)
	}
	if o := sc.orderBy; o != "" {
		column := clause.Column{Name: o, Raw: true}
		db.Order(clause.OrderByColumn{Column: column, Desc: sc.desc})
	}

	return db
}
