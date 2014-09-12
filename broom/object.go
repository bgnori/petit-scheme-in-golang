package broom

import (
	"fmt"
)

//type interface{} interface{} // Anything.

type Undef interface{} // T.B.D.

type Symbol interface {
	//T.B.D.
	GetValue() string
	Eq(other interface{}) bool
}

type Pair interface {
	Car() interface{}
	Cdr() Pair
	SetCar(v interface{}) Undef
	SetCdr(p Pair) Undef
}

func isNull(v interface{}) bool {
	//null?
	return v == nil
}

func isBoolean(v interface{}) bool {
	//boolean?
	_, ok := v.(bool)
	return ok
}

func isChar(v interface{}) bool {
	//char?
	_, ok := v.(rune)
	return ok
}

func isSymbol(v interface{}) bool {
	//symbol?
	_, ok := v.(Symbol) //FIXME
	return ok
}

//eof-object?

func isNumber(v interface{}) bool {
	//number?
	//see golang builtin
	switch v.(type) {
	case int:
	case int8:
	case int16:
	case int32:
	case int64:
	case uint:
	case uint8:
	case uint16:
	case uint32:
	case uint64:
	case float32:
	case float64:
	case complex64:
	case complex128:
	default:
		return false
	}
	return true
}

func isPair(v interface{}) bool {
	//pair?
	_, ok := v.(Pair)
	return ok
}

//port?

type Recur struct {
	env    Environment
	fomals []Symbol
}

func NewRecur(outer Environment, xs interface{}) *Recur {
	r := new(Recur)
	r.env = NewEnvFrame(outer)
	r.fomals = make([]Symbol, 0)
	r.Bind("recur", r)

	ys, _ := xs.([]interface{})
	var key Symbol
	for i, v := range ys {
		if i%2 == 0 {
			key = v.(Symbol)
			r.fomals = append(r.fomals, key)
		} else {
			r.Bind(key.GetValue(), Eval(r, v))
		}
	}
	return r
}

func (r *Recur) Update(xs []interface{}) {
	next := NewEnvFrame(r.Outer())
	next.Bind("recur", r)
	for i, key := range r.fomals {
		next.Bind(key.GetValue(), xs[i])
	}
	r.env = next
}

func (r *Recur) Bind(name string, v interface{}) {
	r.env.Bind(name, v)
}

func (r *Recur) Resolve(name string) interface{} {
	return r.env.Resolve(name)
}

func (r *Recur) SetOuter(outer Environment) {
	r.env.SetOuter(outer)
}

func (r *Recur) Outer() Environment {
	return r.env.Outer()
}

func (r *Recur) Dump() {
	r.env.Dump()
}

func isRecur(v interface{}) bool {
	_, ok := v.(*Recur)
	return ok
}

type Closure func(env Environment, cdr Pair) interface{}

func isProcedure(v interface{}) bool {
	//procedure?
	_, ok := v.(Closure)
	return ok
}

type Syntax Closure

func isSyntax(v interface{}) bool {
	//syntax?
	_, ok := v.(Syntax)
	return ok
}

func isString(v interface{}) bool {
	//string?
	_, ok := v.(string)
	return ok
}

// vector?
func isArray(v interface{}) bool {
	_, ok := v.([]interface{})
	return ok
}

// bytevector?
// define-record-type

func isMap(v interface{}) bool {
	_, ok := v.(map[interface{}]interface{})
	return ok
}

func DumpMap(x interface{}) {
	mx, _ := x.(map[interface{}]interface{})
	fmt.Println("Dumping", mx)
	for k, vx := range mx {
		fmt.Println(k, vx)
	}
}

func EqMap(x, y interface{}) bool {
	mx, _ := x.(map[interface{}]interface{})
	my, _ := y.(map[interface{}]interface{})
	for k, vx := range mx {
		vy, in := my[k]
		if in && vx == vy {
			continue
		} else {
			return false
		}
	}
	for k, vy := range my {
		vx, in := mx[k]
		if in && vx == vy {
			continue
		} else {
			return false
		}
	}
	return true
}

func EqArray(x, y interface{}) bool {
	println("EqArray")
	xs, _ := x.([]interface{})
	ys, _ := y.([]interface{})
	if len(xs) != len(ys) {
		return false
	}
	for i, v := range xs {
		if !Eq(ys[i], v) {
			return false
		}
	}
	return true
}

func Eq(x, y interface{}) bool {
	switch {
	case isMap(x) && isMap(y):
		return EqMap(x, y)
	case isSymbol(x) && isSymbol(y):
		sx, _ := x.(Symbol)
		sy, _ := y.(Symbol)
		return sx.Eq(sy)
	case isPair(x) && isPair(y):
		return Eq(Car(x), Car(y)) && Eq(Cdr(x), Cdr(y))
	case isArray(x) && isArray(y):
		return EqArray(x, y)
	default:
		return x == y
	}
	return false
}
