package main

import (
	"fmt"
)

func add(args *Cons) (Data, error) {
	result := Number(Int(0))
	for ; args != nil; args = args.Next {
		n, ok := args.Data.(Number)
		if !ok {
			return nil, typeError(args.Data, "number")
		}
		result = result.Add(n)
	}
	return result, nil
}

func mul(args *Cons) (Data, error) {
	result := Number(Int(1))
	for ; args != nil; args = args.Next {
		n, ok := args.Data.(Number)
		if !ok {
			return nil, typeError(args.Data, "number")
		}
		result = result.Mul(n)
	}
	return result, nil
}

func sub(args *Cons) (Data, error) {
	if err := assertMin(args, 1); err != nil {
		return nil, err
	}

	result, ok := args.Data.(Number)
	if !ok {
		return nil, typeError(args.Data, "number")
	}

	if args.Next == nil {
		return result.Negate(), nil
	}

	for args = args.Next; args != nil; args = args.Next {
		n, ok := args.Data.(Number)
		if !ok {
			return nil, typeError(args.Data, "number")
		}
		result = result.Sub(n)
	}
	return result, nil
}

func numEq(args *Cons) (Data, error) {
	if err := assertMin(args, 1); err != nil {
		return nil, err
	}

	first, ok := args.Data.(Number)
	if !ok {
		return nil, typeError(args.Data, "number")
	}

	result := true
	for args = args.Next; args != nil; args = args.Next {
		n, ok := args.Data.(Number)
		if !ok {
			return nil, typeError(args.Data, "number")
		}
		result = result && first.NumEq(n)
	}

	return toBool(result), nil
}

func printLisp(args *Cons) (Data, error) {
	for args != nil {
		if printData, ok := args.Data.(PrintData); ok {
			fmt.Println(printData.Print(false))
		} else {
			fmt.Println(args.Data)
		}
		args = args.Next
	}
	return Nil, nil
}

func car(args *Cons) (Data, error) {
	if err := assertRange(args, 1, 1); err != nil {
		return nil, err
	}

	if cons, ok := args.Data.(*Cons); !ok {
		return nil, typeError(args.Data, "cons")
	} else {
		if cons == nil {
			return cons, nil
		}
		return cons.Data, nil
	}
}

func cdr(args *Cons) (Data, error) {
	if err := assertRange(args, 1, 1); err != nil {
		return nil, err
	}

	if cons, ok := args.Data.(*Cons); !ok {
		return nil, typeError(args.Data, "cons")
	} else {
		if cons == nil {
			return cons, nil
		}
		return cons.Next, nil
	}
}

func list(args *Cons) (Data, error) {
	return args.Copy(), nil
}

func typeof(args *Cons) (Data, error) {
	if err := assertRange(args, 1, 1); err != nil {
		return nil, err
	}
	return args.Data.Type(), nil
}

func equal(a, b Data) bool {
	if a.Type() != b.Type() {
		return false
	}

	if aFunc, ok := a.(Function); ok {
		bFunc := b.(Function)
		// Go can not compare funcs to each other, so we can only compare their names
		return aFunc.Name == bFunc.Name
	}

	if a == b {
		return true
	}

	if aCons, ok := a.(*Cons); ok {
		bCons := b.(*Cons)
		for {
			if aCons == nil {
				if bCons == nil {
					return true
				}
				return false
			}
			if !equal(aCons.Data, bCons.Data) {
				return false
			}
			aCons = aCons.Next
			bCons = bCons.Next
		}
		return false
	}

	return false
}

func equalFunc(args *Cons) (Data, error) {
	if err := assertRange(args, 2, 2); err != nil {
		return nil, err
	}
	return toBool(equal(args.Data, args.Next.Data)), nil
}

func quote(args *Cons) (Data, error) {
	if err := assertRange(args, 1, 1); err != nil {
		return nil, err
	}

	if sym, ok := args.Data.(Symbol); ok {
		return sym, nil
	}
	if cons, ok := args.Data.(*Cons); ok {
		if cons == nil {
			return cons, nil
		}
	}
	return args.Data, nil
}

func (s *State) function(args *Cons) (Data, error) {
	if err := assertRange(args, 1, 1); err != nil {
		return nil, err
	}

	if sym, ok := args.Data.(Symbol); !ok {
		return nil, fmt.Errorf("illegal function name %v", args.Data)
	} else if f := s.Funcs[sym]; f != nil {
		return Function{Name: string(sym), Func: f}, nil
	} else {
		return nil, fmt.Errorf("undefined function %v", sym)
	}
}
