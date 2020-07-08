package main

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

func parseData(scanner *bufio.Scanner, text string) (Data, error) {
	switch text {
	case "(":
		return parseCons(scanner)
	case ")":
		return nil, errors.New("unmatched parenthesis")
	case "'":
		scanner.Scan()
		if data, err := parseData(scanner, scanner.Text()); err != nil {
			return nil, err
		} else {
			return &Cons{Data: Symbol("quote"), Next: &Cons{Data: data}}, nil
		}
	case "#":
		scanner.Scan()
		if scanner.Text() != "'" {
			return nil, errors.New("invalid syntax")
		}
		scanner.Scan()
		if data, err := parseData(scanner, scanner.Text()); err != nil {
			return nil, err
		} else {
			return &Cons{Data: Symbol("function"), Next: &Cons{Data: data}}, nil
		}
	case "nil":
		return Nil, nil
	case "t":
		return True, nil
	}

	if len(text) > 0 && text[0] == '"' {
		return String(text[1 : len(text)-1]), nil
	}

	if n, err := strconv.ParseInt(text, 10, 64); err == nil {
		return Int(n), nil
	}
	if n, err := strconv.ParseFloat(text, 64); err == nil {
		return Float(n), nil
	}
	return Symbol(text), nil

}

func parseCons(scanner *bufio.Scanner) (*Cons, error) {
	// Assume ( has been scanned
	var result *Cons
	cons := result
	for scanner.Scan() {
		text := scanner.Text()
		if text == ")" {
			return result, nil
		}

		data, err := parseData(scanner, text)
		if err != nil {
			return nil, err
		}
		new := &Cons{Data: data}
		if result == nil {
			result = new
		} else {
			cons.Next = new
		}
		cons = new
	}
	return nil, nil
}

func splitString(data []byte, atEOF bool) (advance int, token []byte, err error) {
	advance, token, err = bufio.ScanLines(data, atEOF)
	if len(token) == 0 || err != nil {
		return
	}
	str := "\""
	tail := string(token[1:])
	length := 1
	for index := strings.IndexAny(tail, "\\\""); index != -1; index = strings.IndexAny(tail, "\\\"") {
		if tail[index] == '"' {
			str += tail[:index+1]
			return length + index + 1, []byte(str), nil
		}
		val, _, newTail, errQ := strconv.UnquoteChar(tail[index:], '"')
		if errQ != nil {
			err = errQ
			return
		}
		str += tail[:index] + string(val)
		length += len(tail) - len(newTail)
		tail = newTail
	}
	err = errors.New("invalid string syntax")
	return
}

func splitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	advance, token, err = bufio.ScanWords(data, atEOF)
	if err != nil {
		return
	}
	if len(token) > 0 {
		if token[0] == '"' {
			return splitString(data, atEOF)
		}
		if token[0] == '#' {
			j := 1
			if advance-len(token) > 1 {
				j = advance - len(token)
			}
			return j, token[:1], nil
		}
	}
	for i, c := range token {
		if c == '(' || c == ')' || c == '\'' || c == '"' {
			if i == 0 {
				i = 1
			}
			j := i
			if advance-len(token) > 1 {
				j += advance - len(token) - 1
			}
			return j, token[:i], nil
		}
	}
	return
}
