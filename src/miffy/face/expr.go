package face

import (
	"fmt"
)

//Expr 接口
type Expr interface {
	Eval(env Env) float64
}

//Var 变量
type Var string

//Env 结构
type Env map[Var]float64

//Eval 方法
func (v Var) Eval(env Env) float64 {
	return env[v]
}

type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

type unary struct {
	op rune
	x  Expr
}

func (u *unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		fmt.Printf("%T\n", u.x)
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsuported unary operator: %q", u.op))
}

//RunExpr 主方法
func RunExpr() {
	fmt.Println("run expr")
	var t unary
	t = struct {
		op rune
		x  Expr
	}{'-', Var("du")}
	fmt.Println(t.Eval(Env{"du": 100}))
}
