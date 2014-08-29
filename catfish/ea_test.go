package catfish

import (
	"fmt"
	"testing"
)

func TestEvalNumber(t *testing.T) {
	e := NewGlobalRootFrame()
	if Eval(1, e) != 1 {
		t.Error("expected 1")
	}
}

func TestEvalString(t *testing.T) {
	e := NewGlobalRootFrame()
	if Eval("あいう", e) != "あいう" {
		t.Error("expected あいう")
	}
}

func TestEvalVariable(t *testing.T) {
	e := NewGlobalRootFrame()
	e.Bind("a", 42)
	v := Eval(sym("a"), e)
	if v != 42 {
		t.Error("expected 42")
		fmt.Println(v)
	}
}

func TestEvalQuotedSymbol(t *testing.T) {
	e := NewGlobalRootFrame()
	v := Eval(Cons(sym("quote"), Cons(sym("A"), nil)), e)
	if !sym("A").Eq(v) {
		t.Error("expected sym A")
		fmt.Println(v)
	}
}

func TestEvalDefine(t *testing.T) {
	e := NewGlobalRootFrame()
	v := Eval(Cons(sym("define"), Cons(sym("A"), Cons(42, nil))), e)
	if e.Resolve("A") != 42 {
		t.Error("expected 42")
	}
	if !sym("A").Eq(v) {
		t.Error("expected sym A")
	}
}

func TestEvalIF(t *testing.T) {
	e := NewGlobalRootFrame()
	v := Eval(Cons(sym("if"), Cons(true, Cons(42, Cons(123, nil)))), e)
	if v != 42 {
		t.Error("expected 42")
	}
	v = Eval(Cons(sym("if"), Cons(false, Cons(42, Cons(123, nil)))), e)
	if v != 123 {
		t.Error("expected 123")
	}
}

func TestEvalLambda(t *testing.T) {
	e := NewGlobalRootFrame()
	f := Eval(List(nil, sym("lambda"), List(nil, sym("x")), sym("x")), e)
	if _, ok := f.(Procedure); !ok {
		t.Error("expected Procedure")
	}
	println("having", f)
	v := Eval(List(nil, f, 123), e)
	if v != 123 {
		t.Error("expected 123")
	}
}