package opencond

import (
	"fmt"

	"gorm.io/gorm/clause"
)

type Schema struct {
	Fields []*schema `json:"fields"`
	Groups []*pair   `json:"groups"`
}

type State struct {
	patterns map[string]*pattern
	schemas  []*schema
	fields   map[string]*pattern
	groups   map[string]*pair
	schema   Schema
}

// Schemas 获取约束
func (st State) Schemas() []*schema {
	return st.schemas
}

// Schema 获取约束
func (st State) Schema() Schema {
	return st.schema
}

// Interpreter 根据传输的环境参数解释出 gorm 表达式
func (st State) Interpreter(envs Environments) (clause.Expression, error) {
	exprs := make([]clause.Expression, 0, len(envs))
	for _, env := range envs {
		key, val := env.Key, env.Value
		if key == "" || val == "" { // value 为空就忽略
			continue
		}

		pt := st.patterns[key]
		if pt == nil {
			return nil, fmt.Errorf("%s不存在", key)
		}

		opt := env.Operator
		op := pt.operators[opt]
		if op == nil {
			return nil, fmt.Errorf("%s不支持%s表达式", key, opt)
		}

		expr, err := op.interpreter(pt, env)
		if err != nil {
			return nil, err
		}

		exprs = append(exprs, expr)
	}
	if len(exprs) == 0 {
		return nil, nil
	}

	return clause.AndConditions{Exprs: exprs}, nil
}

func (st State) Interp(fields Fields, groupBy string) (clause.Expression, error) {
	length := len(fields)
	if length == 0 && groupBy == "" {
		return nil, nil
	}

	exps := make([]clause.Expression, length+1)
	if groupBy != "" {
		if _, ok := st.groups[groupBy]; !ok {
			return nil, fmt.Errorf("%s不允许groupBy", groupBy)
		}

		gb := clause.GroupBy{Columns: []clause.Column{{Name: groupBy}}}
		exps = append(exps, gb)
	}

	for _, f := range fields {
		key := f.Key
		pt := st.fields[key]
		if pt == nil {
			return nil, fmt.Errorf("%s不存在", key)
		}

		opt := f.Operator
		op := pt.operators[opt]
		if op == nil {
			return nil, fmt.Errorf("%s不支持%s表达式", key, opt)
		}

		expr, err := op.interpreter(pt, nil)
		if err != nil {
			return nil, err
		}

		exps = append(exps, expr)
	}

	if len(exps) == 0 {
		return nil, nil
	}

	return clause.AndConditions{Exprs: exps}, nil
}
