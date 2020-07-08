package main

import (
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	state := NewState()
	if data, err := state.Eval(strings.NewReader("(+ 629 424 369)")); err != nil {
		t.Fatal(err)
	} else if data != Int(1422) {
		t.Fatalf("got %v (type %[1]T), expected 1422 (type Int)", data)
	}

	if data, err := state.Eval(strings.NewReader("(+ 24 10.12)")); err != nil {
		t.Fatal(err)
	} else if data != Float(34.12) {
		t.Fatalf("got %v (type %[1]T), expected 34.12 (type Float)", data)
	}

	if data, err := state.Eval(strings.NewReader("(+)")); err != nil {
		t.Fatal(err)
	} else if data != Int(0) {
		t.Fatalf("got %v (type %[1]T), expected 0 (type Int)", data)
	}

	if _, err := state.Eval(strings.NewReader("(+ 10 20 'a)")); err == nil {
		t.Fatalf("got error %v, expected non-nil", err)
	} else if _, ok := err.(*TypeError); !ok {
		t.Fatalf("got error %v (type %[1]T), expected TypeError", err)
	}
}

func TestSub(t *testing.T) {
	state := NewState()
	if data, err := state.Eval(strings.NewReader("(- 629 424 369)")); err != nil {
		t.Fatal(err)
	} else if data != Int(-164) {
		t.Fatalf("got %v (type %[1]T), expected -164 (type Int)", data)
	}

	if data, err := state.Eval(strings.NewReader("(- 629)")); err != nil {
		t.Fatal(err)
	} else if data != Int(-629) {
		t.Fatalf("got %v (type %[1]T), expected -629 (type Int)", data)
	}

	if data, err := state.Eval(strings.NewReader("(- 44 9.88)")); err != nil {
		t.Fatal(err)
	} else if data != Float(34.12) {
		t.Fatalf("got %v (type %[1]T), expected 34.12 (type Float)", data)
	}

	if _, err := state.Eval(strings.NewReader("(-)")); err == nil {
		t.Fatalf("got error %v, expected non-nil", err)
	} else if _, ok := err.(*MoreArgsError); !ok {
		t.Fatalf("got error %v (type %[1]T), expected MoreArgsError", err)
	}

	if _, err := state.Eval(strings.NewReader("(- 10 20 'a)")); err == nil {
		t.Fatalf("got error %v, expected non-nil", err)
	} else if _, ok := err.(*TypeError); !ok {
		t.Fatalf("got error %v (type %[1]T), expected TypeError", err)
	}
}

func TestMul(t *testing.T) {
	state := NewState()
	if data, err := state.Eval(strings.NewReader("(* 629 424 369)")); err != nil {
		t.Fatal(err)
	} else if data != Int(98410824) {
		t.Fatalf("got %v (type %[1]T), expected 94810824 (type Int)", data)
	}

	if data, err := state.Eval(strings.NewReader("(* 2 3.141592653589793)")); err != nil {
		t.Fatal(err)
	} else if data != Float(6.283185307179586) {
		t.Fatalf("got %v (type %[1]T), expected 6.283185307179586 (type Float)", data)
	}

	if data, err := state.Eval(strings.NewReader("(*)")); err != nil {
		t.Fatal(err)
	} else if data != Int(1) {
		t.Fatalf("got %v (type %[1]T), expected 1 (type Int)", data)
	}

	if _, err := state.Eval(strings.NewReader("(* 10 20 'a)")); err == nil {
		t.Fatalf("got error %v, expected non-nil", err)
	} else if _, ok := err.(*TypeError); !ok {
		t.Fatalf("got error %v (type %[1]T), expected TypeError", err)
	}
}

func TestNumEq(t *testing.T) {
	state := NewState()
	if data, err := state.Eval(strings.NewReader("(= 629 424)")); err != nil {
		t.Fatal(err)
	} else if data != False {
		t.Fatalf("got %v, expected nil", data)
	}

	if data, err := state.Eval(strings.NewReader("(= 629 629)")); err != nil {
		t.Fatal(err)
	} else if data != True {
		t.Fatalf("got %v, expected t", data)
	}

	if data, err := state.Eval(strings.NewReader("(= 629 629 629.0)")); err != nil {
		t.Fatal(err)
	} else if data != True {
		t.Fatalf("got %v, expected t", data)
	}

	if data, err := state.Eval(strings.NewReader("(= 629 630 629.0)")); err != nil {
		t.Fatal(err)
	} else if data != False {
		t.Fatalf("got %v, expected nil", data)
	}

	if data, err := state.Eval(strings.NewReader("(= 9007199254740993 9007199254740993.0)")); err != nil {
		t.Fatal(err)
	} else if data != False {
		t.Fatalf("got %v, expected nil", data)
	}

	if data, err := state.Eval(strings.NewReader("(= 9007199254740992 9007199254740993.0)")); err != nil {
		t.Fatal(err)
	} else if data != True {
		t.Fatalf("got %v, expected t", data)
	}

	if _, err := state.Eval(strings.NewReader("(=)")); err == nil {
		t.Fatalf("got error %v, expected non-nil", err)
	} else if _, ok := err.(*MoreArgsError); !ok {
		t.Fatalf("got error %v (type %[1]T), expected MoreArgsError", err)
	}
}

func TestCar(t *testing.T) {
	state := NewState()
	if data, err := state.Eval(strings.NewReader("(car '(a 1 2 b))")); err != nil {
		t.Fatal(err)
	} else if data != Symbol("a") {
		t.Fatalf("got %v (type %[1]T), expected a (type Symbol)", data)
	}

	if data, err := state.Eval(strings.NewReader("(car nil)")); err != nil {
		t.Fatal(err)
	} else if data != Nil {
		t.Fatalf("got %v (type %[1]T), expected nil", data)
	}

	if _, err := state.Eval(strings.NewReader("(car'a)")); err == nil {
		t.Fatalf("got error %v, expected non-nil", err)
	} else if _, ok := err.(*TypeError); !ok {
		t.Fatalf("got error %v (type %[1]T), expected TypeError", err)
	}

	if _, err := state.Eval(strings.NewReader("(car)")); err == nil {
		t.Fatalf("got error %v, expected non-nil", err)
	} else if _, ok := err.(*MoreArgsError); !ok {
		t.Fatalf("got error %v (type %[1]T), expected MoreArgsError", err)
	}

	if _, err := state.Eval(strings.NewReader("(car '(1 2) '(3 4))")); err == nil {
		t.Fatalf("got error %v, expected non-nil", err)
	} else if _, ok := err.(*LessArgsError); !ok {
		t.Fatalf("got error %v (type %[1]T), expected LessArgsError", err)
	}
}

func TestCdr(t *testing.T) {
	state := NewState()
	if data, err := state.Eval(strings.NewReader("(cdr '(a 1 2))")); err != nil {
		t.Fatal(err)
	} else if !equal(data, &Cons{Data: Int(1), Next: &Cons{Data: Int(2)}}) {
		t.Fatalf("got %v (type %[1]T), expected (1 2) (type *Cons)", data)
	}

	if data, err := state.Eval(strings.NewReader("(cdr nil)")); err != nil {
		t.Fatal(err)
	} else if data != Nil {
		t.Fatalf("got %v (type %[1]T), expected nil", data)
	}

	if _, err := state.Eval(strings.NewReader("(cdr'a)")); err == nil {
		t.Fatalf("got error %v, expected non-nil", err)
	} else if _, ok := err.(*TypeError); !ok {
		t.Fatalf("got error %v (type %[1]T), expected TypeError", err)
	}

	if _, err := state.Eval(strings.NewReader("(cdr)")); err == nil {
		t.Fatalf("got error %v, expected non-nil", err)
	} else if _, ok := err.(*MoreArgsError); !ok {
		t.Fatalf("got error %v (type %[1]T), expected MoreArgsError", err)
	}

	if _, err := state.Eval(strings.NewReader("(cdr '(1 2) '(3 4))")); err == nil {
		t.Fatalf("got error %v, expected non-nil", err)
	} else if _, ok := err.(*LessArgsError); !ok {
		t.Fatalf("got error %v (type %[1]T), expected LessArgsError", err)
	}
}
