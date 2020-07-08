package main

import (
	"fmt"
)

type TypeError struct {
	Data Data
	Type string
}

func (e *TypeError) Error() string {
	return fmt.Sprintf("value %v is not of type %s", e.Data, e.Type)
}

type MoreArgsError struct {
	Cons *Cons
	Min  int
}

func (e *MoreArgsError) Error() string {
	return fmt.Sprintf("not enough arguments: %v, at least %d expected", e.Cons, e.Min)
}

type LessArgsError struct {
	Cons *Cons
	Max  int
}

func (e *LessArgsError) Error() string {
	return fmt.Sprintf("too many arguments: %v, at most %d expected", e.Cons, e.Max)
}

func assertMin(cons *Cons, min int) error {
	length := 0
	for c := cons; c != nil; c = c.Next {
		length++
		if length >= min {
			return nil
		}
	}
	return &MoreArgsError{Cons: cons, Min: min}
}

func assertMaxLen(cons *Cons, max int) (int, error) {
	length := 0
	for c := cons; c != nil; c = c.Next {
		length++
		if length > max {
			return length, &LessArgsError{Cons: cons, Max: max}
		}
	}
	return length, nil
}

func assertMax(cons *Cons, max int) error {
	_, err := assertMaxLen(cons, max)
	return err
}

func assertRange(cons *Cons, min, max int) error {
	if length, err := assertMaxLen(cons, max); err != nil {
		return err
	} else if length < min {
		return &MoreArgsError{Cons: cons, Min: min}
	}
	return nil
}

func typeError(data Data, t string) error {
	return &TypeError{Data: data, Type: t}
}
