package dynsql

import "gorm.io/gorm/clause"

var (
	Eq      = &eqOpera{}
	Gt      = &gtOpera{}
	Lt      = &ltOpera{}
	Ge      = &geOpera{}
	Le      = &leOpera{}
	In      = &inOpera{}
	NotIn   = &notInOpera{}
	Like    = &likeOpera{}
	NotLike = &notLikeOpera{}
)

type Operator interface {
	String() string
	Symbol() string
	split() bool
	expr(string, []string) *clause.Expr
}

type eqOpera struct{}

func (eqOpera) String() string    { return "eq" }
func (eqOpera) SQLString() string { return "=" }
func (eqOpera) split() bool       { return false }

func (eqOpera) expr(tp Typer, col string, values []string) (clause.Expression, error) {
	val, err := tp.cast(values[1])
	if err != nil {
		return nil, err
	}
	exp := clause.Eq{
		Column: col,
		Value:  val,
	}
	return exp, nil
}

type gtOpera struct{}

func (gtOpera) String() string { return "gt" }
func (gtOpera) Symbol() string { return ">" }
func (gtOpera) split() bool    { return false }
func (gtOpera) expr(tp Typer, col string, values []string) (clause.Expression, error) {
	val, err := tp.cast(values[1])
	if err != nil {
		return nil, err
	}
	exp := clause.Neq{
		Column: col,
		Value:  val,
	}
	return exp, nil
}

type ltOpera struct{}

func (ltOpera) String() string { return "lt" }
func (ltOpera) Symbol() string { return "<" }
func (ltOpera) split() bool    { return false }
func (ltOpera) expr(tp Typer, col string, values []string) (clause.Expression, error) {
	val, err := tp.cast(values[1])
	if err != nil {
		return nil, err
	}
	exp := clause.Lt{
		Column: col,
		Value:  val,
	}
	return exp, nil
}

type geOpera struct{}

func (geOpera) String() string { return "ge" }
func (geOpera) Symbol() string { return ">=" }
func (geOpera) split() bool    { return false }

type leOpera struct{}

func (leOpera) String() string { return "le" }
func (leOpera) Symbol() string { return "<=" }
func (leOpera) split() bool    { return false }

type inOpera struct{}

func (inOpera) String() string { return "in" }
func (inOpera) Symbol() string { return "IN" }
func (inOpera) split() bool    { return true }

type notInOpera struct{}

func (notInOpera) String() string { return "notin" }
func (notInOpera) Symbol() string { return "NOT IN" }
func (notInOpera) split() bool    { return true }

type likeOpera struct{}

func (likeOpera) String() string { return "like" }
func (likeOpera) Symbol() string { return "LIKE" }
func (likeOpera) split() bool    { return false }

type notLikeOpera struct{}

func (notLikeOpera) String() string { return "notlike" }
func (notLikeOpera) Symbol() string { return "NOT LIKE" }
func (notLikeOpera) split() bool    { return false }
