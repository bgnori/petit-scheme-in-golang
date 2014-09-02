package broom

import (
  "fmt"
  "strings"
)

type pairImpl struct {
	car Value
	cdr Value
}

func Cons(car, cdr Value) Pair {
	return &pairImpl{car: car, cdr: cdr}
}

func Car(v Value) Value {
	u, ok := v.(Pair)
	if !ok {
		panic("non pair object for Car()")
	}
	return u.Car()
}

func Cdr(v Value) Value {
	u, ok := v.(Pair)
	if !ok {
		panic("non pair object for Cdr()")
	}
	return u.Cdr()
}

func (p *pairImpl) Car() Value {
	return p.car
}

func (p *pairImpl) Cdr() Value {
	return p.cdr
}

func (p *pairImpl) SetCar(v Value) Undef {
	p.car = v
	return nil
}

func (p *pairImpl) SetCdr(v Value) Undef {
	p.cdr = v
	return nil
}

func (p *pairImpl) String() string {
  //assume that proper list
  xs := List2Arr(p)
  ss := make([]string, 0)
  for _, x := range xs {
    ss = append(ss, fmt.Sprint(x))
  }
  return "(" + strings.Join(ss, " ") + ")"
}

func sub(v Value, xs []Value) []Value {
	if v == nil {
		return xs
	} else {
                xs = append(xs, Car(v))
		return sub(Cdr(v), xs)
	}
}

func List2Arr(v Value) []Value {
	return sub(v, make([]Value, 0))
}

func List(cdr Value, xs ...Value) Value {
	//(list obj... )
	// this function supports . cdr, for none proper list
	if len(xs) == 0 {
		return cdr
	}
	return Cons(xs[0], List(cdr, xs[1:]...))
}

func isList(xs Value) bool {
	if isNull(xs) {
		return true
	}
	if isPair(xs) {
		return isList(Cdr(xs))
	}
	return false
}

func Length(xs Value) int {
	if isNull(xs) {
		return 0
	}
	if isPair(xs) {
		return Length(Cdr(xs)) + 1
	}
	panic("proper list required")
}