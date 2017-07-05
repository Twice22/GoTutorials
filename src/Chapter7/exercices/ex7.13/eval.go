package eval

import (
	"fmt"
	"math"
)

// NOTE: the package only export type Expr, Var and Env (uppercase first letter)

type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	Eval(env Env) float64
	Check(vars map[Var]bool) error
	String() string // Answer to Q7.13
}

// A Var identifies a variable, e.g., x.
type Var string

// A literal is a numeric constant, e.g., 3.141.
type literal float64

// A unary represents a unary operator expression, e.g., -x.
type unary struct {
	op rune // one of '+', '-'
	x Expr
}

// A binary represents a binary operator expression, e.g., x+y.
type binary struct {
	op rune // one of '+', '-', '*', '/'
	x, y Expr
}

// A call represents a function call expression, e.g., sin(x).
type call struct {
	fn string // one of "pow", "sin", "sqrt"
	args []Expr
}

// To evaluate an expression containing variables, we'll need an env
// that maps variable names to values.
type Env map[Var]float64


// Var satisfies Expr interface
func (v Var) Eval(env Env) float64 {
	return env[v]
}

// Var need to have a String() method to satisfy Expr interface (Answer to Q7.13)
func (v Var) String() string {
	// Var is a string but we still need to convert Var to a string before using print
	return fmt.Sprintf("%s", string(v))
}

// literal satisfies Expr interface
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

// literal need to have a String() method to satisfy Expr interface (Answer to Q7.13)
func (l literal) String() string {
	return fmt.Sprintf("%f", float64(l))
}

// the method for call evaluates the arguments to pow, sin, or sqrt
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

// Answer to Q7.13
func (u unary) String() string {
	return fmt.Sprintf("%s%s", u.op, u.x)
}

func (b binary) String() string {
	// use string() to convert b.op from rune to string
	return fmt.Sprintf("(%s %s %s)", b.x, string(b.op), b.y)
}

func (c call) String() string {
	switch c.fn {
	case "pow":
		return fmt.Sprintf("%s(%s, %s)", c.fn, c.args[0], c.args[1])
	case "sin", "sqrt":
		return fmt.Sprintf("%s(%s)", c.fn, c.args[0])
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}