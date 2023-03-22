package dynsql

import (
	"encoding/json"
	"gorm.io/gorm/clause"
	"net/url"
)

type Conditions []*Condition

type Condition struct {
	Col string `json:"col"`
	Op  string `json:"op"`
	Val string `json:"val"`
}

func (cs *Conditions) UnmarshalBind(raw string) error {
	data, err := url.QueryUnescape(raw)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(data), cs)
}

type Actuator interface {
	Inter() (clause.Expression, error)
}

type environment struct {
	sep     string
	ops     map[string]Operator
	columns map[string]*parsedColumn
}

func (env *environment) Inter(conds Conditions) (clause.Expression, error) {
	if len(conds) == 0 {
		return nil, nil
	}

	//for _, cond := range conds {
	//	// col, op := cond.Col, cond.Op
	//	col := cond.Col
	//	op := env.matchOp(col, cond.Op)
	//	if op == nil {
	//		return nil, errors.New("")
	//	}
	//
	//	val := strings.TrimSpace(cond.Val)
	//	values := []string{val}
	//	if op.split() {
	//		values = strings.Split(val, env.sep)
	//	}
	//
	//}

	return nil, nil
}

func (*environment) expr(col string, values []string) {

}

func (env *environment) matchOp(col, op string) Operator {
	ret, ok := env.ops[op]
	if !ok {
		return nil
	}

	pc, exist := env.columns[col]
	if !exist {
		return nil
	}

	if _, ok = pc.allowOps[op]; !ok {
		return nil
	}

	return ret
}

type parsedColumn struct {
	allowOps map[string]struct{}
}
