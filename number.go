package main

import "math"

type Number interface {
	Data
	Float() Float
	Int() Int
	IsInt() bool
	Add(n Number) Number
	Sub(n Number) Number
	Mul(n Number) Number
	Negate() Number
	NumEq(n Number) bool
}

type Int int64

func (a Int) Float() Float {
	return Float(int64(a))
}

func (a Int) Int() Int {
	return a
}

func (_ Int) IsInt() bool {
	return true
}

func (_ Int) Type() Symbol {
	return "int"
}

func (a Int) Add(b Number) Number {
	if n, ok := b.(Int); ok {
		return a + n
	}
	return a.Float().Add(b)
}

func (a Int) Sub(b Number) Number {
	if n, ok := b.(Int); ok {
		return a - n
	}
	return a.Float().Sub(b)
}

func (a Int) Mul(b Number) Number {
	if n, ok := b.(Int); ok {
		return a * n
	}
	return a.Float().Mul(b)
}

func (a Int) Negate() Number {
	return -a
}

func (a Int) NumEq(n Number) bool {
	return a == n.Int()
}

type Float float64

func (a Float) Float() Float {
	return a
}

func (a Float) Int() Int {
	return Int(float64(a))
}

func (a Float) IsInt() bool {
	return math.Mod(float64(a), 1) == 0
}

func (_ Float) Type() Symbol {
	return "float"
}

func (a Float) Add(b Number) Number {
	return a + b.Float()
}

func (a Float) Sub(b Number) Number {
	return a - b.Float()
}

func (a Float) Mul(b Number) Number {
	return a * b.Float()
}

func (a Float) Negate() Number {
	return -a
}

func (a Float) NumEq(b Number) bool {
	if n, ok := b.(Float); ok {
		return a == n
	}
	return a.IsInt() && a.Int().NumEq(b)
}
