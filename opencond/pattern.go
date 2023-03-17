package opencond

import "fmt"

type patternBuilder struct {
	key       string
	column    string
	desc      string
	datatype  datatype
	operators []operator
	enums     Pairs
}

func Builder(key, column, desc string, dt datatype) *patternBuilder {
	return &patternBuilder{key: key, column: column, desc: desc, datatype: dt}
}

func (sb *patternBuilder) Operators(ops ...operator) *patternBuilder {
	sb.operators = ops
	return sb
}

func (sb *patternBuilder) Enums(enums Pairs) *patternBuilder {
	sb.enums = enums
	if len(sb.operators) == 0 { // 枚举类型默认操作符号
		sb.operators = []operator{OpEq, OpNe, OpIn, OpNotIn}
	}
	return sb
}

func (sb *patternBuilder) Build() *pattern {
	operators := sb.operators
	if len(operators) == 0 {
		operators = sb.datatype.operators()
	}
	if sb.datatype.key() == TypeBool.key() && len(sb.enums) == 0 {
		sb.enums = Pairs{{Key: "true", Desc: "是"}, {Key: "false", Desc: "否"}}
	}

	// 关系运算符去重
	size := len(operators)
	ops := make(Pairs, 0, size)
	opm := make(map[string]operator, size)
	for _, op := range operators {
		key := op.String()
		if _, exist := opm[key]; exist {
			continue
		}
		opm[key] = op
		ops = append(ops, &pair{Key: key, Desc: op.desc()})
	}

	// 处理 enums
	ehm := make(map[string]*pair, 8)
	enumKeys := make([]string, 0, 8)
	var enums Pairs
	enum := len(sb.enums) > 0
	if enum {
		for _, p := range sb.enums {
			key := p.Key
			if _, exist := ehm[key]; exist {
				continue
			}
			ehm[key] = p
			enums = append(enums, &pair{Key: key, Desc: p.Desc})
			enumKeys = append(enumKeys, key)
		}
	}

	shm := &schema{Key: sb.key, Desc: sb.desc, Type: sb.datatype.key(), Operators: ops, Enum: enum, Enums: enums}

	return &pattern{
		key:       sb.key,
		column:    sb.column,
		desc:      sb.desc,
		datatype:  sb.datatype,
		operators: opm,
		enum:      enum,
		enumKeys:  enumKeys,
		enums:     ehm,
		schema:    shm,
	}
}

type pattern struct {
	key       string
	column    string
	desc      string
	datatype  datatype
	operators map[string]operator
	enum      bool
	enumKeys  []string
	enums     map[string]*pair
	schema    *schema
}

func (pt pattern) inEnums(str string) error {
	if !pt.enum {
		return nil
	}
	if _, exist := pt.enums[str]; exist {
		return nil
	}

	return fmt.Errorf("%s必须是%v中的一个", pt.key, pt.enumKeys)
}

type Patterns []*pattern

func (pts Patterns) State() State {
	length := len(pts)
	patterns := make(map[string]*pattern, length)
	schemas := make([]*schema, 0, length)

	for _, pt := range pts {
		key := pt.key
		if _, exist := patterns[key]; exist {
			continue
		}

		patterns[key] = pt
		schemas = append(schemas, pt.schema)
	}

	return State{patterns: patterns, schemas: schemas}
}

type Pattern struct {
	fields []*pattern
	groups []*pair
}

func (pt *Pattern) Field(pts ...*pattern) *Pattern {
	for _, p := range pts {
		if p != nil {
			pt.fields = append(pt.fields, p)
		}
	}

	return pt
}

func (pt *Pattern) Group(prs ...*pair) *Pattern {
	for _, p := range prs {
		if p != nil {
			pt.groups = append(pt.groups, p)
		}
	}

	return pt
}

func (pt *Pattern) State() State {
	fsz := len(pt.fields)
	fieldMap := make(map[string]*pattern, fsz)
	fields := make([]*schema, 0, fsz)

	for _, p := range pt.fields {
		key := p.key
		if _, exist := fieldMap[key]; exist {
			continue
		}

		fieldMap[key] = p
		fields = append(fields, p.schema)
	}

	gsz := len(pt.groups)
	groupMap := make(map[string]*pair, gsz)
	groups := make([]*pair, 0, gsz)
	for _, p := range pt.groups {
		key := p.Key
		if _, exist := groupMap[key]; exist {
			continue
		}

		groupMap[key] = p
		groups = append(groups, p)
	}

	shm := Schema{Fields: fields, Groups: groups}

	return State{
		fields: fieldMap,
		groups: groupMap,
		schema: shm,
	}
}
