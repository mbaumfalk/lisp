package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

type State struct {
	Funcs map[Symbol]LispFunc
	Vars  map[Symbol]Data

	Specials map[Symbol]LispFunc
}

func NewState() *State {
	state := &State{}
	state.Funcs = map[Symbol]LispFunc{
		"+":     add,
		"*":     mul,
		"-":     sub,
		"=":     numEq,
		"print": printLisp,
		"car":   car,
		"cdr":   cdr,
		"list":  list,
		"eval":  state.evalFunc,
		"call":  state.callFunc,
		"apply": state.applyFunc,
		"type":  typeof,
		"equal": equalFunc,
	}
	state.Specials = map[Symbol]LispFunc{
		"quote":    quote,
		"function": state.function,
	}
	return state
}

func (s *State) ReadFrom(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	scanner.Split(splitFunc)

	for scanner.Scan() {
		if cons, err := parseData(scanner, scanner.Text()); err != nil {
			return err
		} else if _, err := s.eval(cons); err != nil {
			return err
		}
	}
	return scanner.Err()
}

func (s *State) Eval(reader io.Reader) (Data, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(splitFunc)

	if scanner.Scan() {
		if cons, err := parseData(scanner, scanner.Text()); err != nil {
			return nil, err
		} else if data, err := s.eval(cons); err != nil {
			return nil, err
		} else {
			return data, nil
		}
	}
	return nil, scanner.Err()
}

type InteractiveState State

func NewInteractiveState() *InteractiveState {
	return (*InteractiveState)(NewState())
}

func (i *InteractiveState) ReadFrom(reader io.Reader) {
	s := (*State)(i)
	scanner := bufio.NewScanner(reader)
	scanner.Split(splitFunc)

	fmt.Print("> ")
	for scanner.Scan() {
		if cons, err := parseData(scanner, scanner.Text()); err != nil {
			fmt.Println(err)
		} else if data, err := s.eval(cons); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(data)
		}
		fmt.Print("> ")
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Println()
}
