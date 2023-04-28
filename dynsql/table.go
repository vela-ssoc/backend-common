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
	ret := new(scope)
	if input.empty() || (len(tbl.orderMap) == 0 && len(tbl.filterMap) == 0 && len(tbl.groupMap) == 0) {
		return ret, nil
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
	if len(items) != 0 {
		ret.where = clause.And(items...)
	}

	ret.desc = input.Desc
	if order != "" && len(tbl.orderMap) != 0 {
		if _, exist := tbl.orderMap[order]; !exist {
			return nil, &Error{name: "排序条件", value: order}
		}
		ret.orderBy = order
	}

	if group != "" && len(tbl.groupMap) != 0 {
		if _, exist := tbl.groupMap[group]; !exist {
			return nil, &Error{name: "分组条件", value: group}
		}
		ret.groupBy = group
	}

	return ret, nil
}

type Scope interface {
	Where(*gorm.DB) *gorm.DB
	GroupBy(*gorm.DB) *gorm.DB
	OrderBy(*gorm.DB) *gorm.DB
	GroupColumn() string
}

type scope struct {
	where   clause.Expression
	groupBy string
	orderBy string
	desc    bool
}

func (sc *scope) Where(db *gorm.DB) *gorm.DB {
	if w := sc.where; w != nil {
		return db.Where(w)
	}
	return db
}

func (sc *scope) GroupBy(db *gorm.DB) *gorm.DB {
	if g := sc.groupBy; g != "" {
		return db.Group(g)
	}
	return db
}

func (sc *scope) OrderBy(db *gorm.DB) *gorm.DB {
	if o := sc.orderBy; o != "" {
		column := clause.Column{Name: o, Raw: true}
		db.Order(clause.OrderByColumn{Column: column, Desc: sc.desc})
	}
	return db
}

func (sc *scope) GroupColumn() string {
	return sc.groupBy
}
