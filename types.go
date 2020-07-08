package main

import (
	"fmt"
	"strconv"
)

type Data interface {
	Type() Symbol
}

type PrintData interface {
	Data
	Print(repr bool) string
}

// TODO: Next Data
type Cons struct {
	Data Data
	Next *Cons
}

func (_ *Cons) Type() Symbol {
	// TODO: nil?
	return "cons"
}

type Symbol string

func (_ Symbol) Type() Symbol {
	return "symbol"
}

type String string

func (_ String) Type() Symbol {
	return "string"
}

type trueType struct{}

func (_ trueType) Type() Symbol {
	return "true"
}

func (_ trueType) String() string {
	return "t"
}

var (
	True  = trueType{}
	Nil   = (*Cons)(nil)
	False = Nil
)

func toBool(b bool) Data {
	if b {
		return True
	}
	return False
}

type LispFunc func(*Cons) (Data, error)

type Function struct {
	Name string
	Func LispFunc
}

func (_ Function) Type() Symbol {
	return "function"
}

func (f Function) String() string {
	return fmt.Sprintf("<function %s>", f.Name)
}

func (s String) Print(repr bool) string {
	if repr {
		return strconv.Quote(string(s))
	}
	return string(s)
}

func (s String) String() string {
	return s.Print(true)
}

func (c *Cons) Print(repr bool) string {
	if c == nil {
		return "nil"
	}
	if c.Data == Symbol("quote") && c.Next != nil && c.Next.Next == nil {
		return "'" + fmt.Sprint(c.Next.Data)
	}
	if c.Data == Symbol("function") && c.Next != nil && c.Next.Next == nil {
		return "#'" + fmt.Sprint(c.Next.Data)
	}
	result := "("
	for {
		if printData, ok := c.Data.(PrintData); ok {
			result += printData.Print(repr)
		} else {
			result += fmt.Sprint(c.Data)
		}
		if c.Next == nil {
			result += ")"
			return result
		}
		result += " "
		c = c.Next
	}
}

func (c *Cons) String() string {
	return c.Print(true)
}

func (c *Cons) Copy() *Cons {
	if c == nil {
		return nil
	}
	var result = *c
	cons := &result
	for c = c.Next; c != nil; c = c.Next {
		copy := *c
		cons.Next = &copy
		cons = cons.Next
	}
	return &result
}
