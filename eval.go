package main

import (
	"fmt"
)

func (s *State) call(function Data, args *Cons) (Data, error) {
	var f LispFunc
	if sym, ok := function.(Symbol); ok {
		f = s.Funcs[sym]
	} else if fun, ok := function.(Function); ok {
		f = fun.Func
	} else {
		return nil, fmt.Errorf("illegal function call %v", &Cons{Data: function, Next: args})
	}

	if f == nil {
		return nil, fmt.Errorf("undefined function %v", function)
	}

	return f(args)
}

func (s *State) eval(data Data) (Data, error) {
	if cons, ok := data.(*Cons); ok {
		if cons == nil {
			return cons, nil
		}
		if sym, ok := cons.Data.(Symbol); ok {
			special := s.Specials[sym]
			if special != nil {
				return special(cons.Next)
			}
		}

		args := cons.Next.Copy()
		for a := args; a != nil; a = a.Next {
			if data, err := s.eval(a.Data); err != nil {
				return nil, err
			} else {
				a.Data = data
			}
		}
		return s.call(cons.Data, args)
	}

	if str, ok := data.(Symbol); ok {
		if val, ok := s.Vars[str]; !ok {
			return nil, fmt.Errorf("undefined variable %s", str)
		} else {
			return val, nil
		}
	}

	return data, nil
}

func (s *State) evalFunc(args *Cons) (Data, error) {
	if err := assertRange(args, 1, 1); err != nil {
		return nil, err
	}
	return s.eval(args.Data)
}

func (s *State) callFunc(args *Cons) (Data, error) {
	if err := assertMin(args, 1); err != nil {
		return nil, err
	}
	return s.call(args.Data, args.Next)
}

func (s *State) applyFunc(args *Cons) (Data, error) {
	if err := assertRange(args, 2, 2); err != nil {
		return nil, err
	}

	if data, ok := args.Next.Data.(*Cons); ok {
		return s.call(args.Data, data)
	}
	return nil, typeError(args.Next.Data, "cons")
}
