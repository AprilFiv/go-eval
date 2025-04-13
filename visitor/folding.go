package visitor

import (
	"context"
	"fmt"
	"github.com/AprilFiv/go-eval/ast"
	"github.com/AprilFiv/go-eval/runtime"
)

type Folding struct {
	Err error
}

func (fold *Folding) appendError(err error) {
	if fold.Err != nil {
		fold.Err = fmt.Errorf("%w \n%w", fold.Err, err)
		return
	}
	fold.Err = err
}

func (fold *Folding) Visit(node *ast.Node) {
	switch t := (*node).(type) {
	case *ast.ConstNode:
	case *ast.VariableNode:
	case *ast.FuncNode:
		ctx := context.Background()
		var params []interface{}
		for _, arg := range t.Args {
			constNode, ok := arg.(*ast.ConstNode)
			if ok {
				params = append(params, constNode.ConstValue)
			}
		}
		//全部都是常量
		if len(params) == len(t.Args) {
			res, err := t.Executable(ctx, params...)
			if err != nil {
				fold.appendError(err)
				return
			}

			*node = &ast.ConstNode{
				ConstValue: res,
			}
		}
	case *ast.UnaryNode:
		constNode, ok := t.Operand.(*ast.ConstNode)
		if ok {
			switch t.Operator {
			case "!":
				res, err := runtime.Not(constNode.ConstValue)
				if err != nil {
					fold.appendError(err)
					return
				}
				*node = &ast.ConstNode{ConstValue: res}
			case "-":
				res, err := runtime.Negative(constNode.ConstValue)
				if err != nil {
					fold.appendError(err)
					return
				}
				*node = &ast.ConstNode{ConstValue: res}
			}
		}
	case *ast.BinaryNode:
		leftNode, ok := t.Left.(*ast.ConstNode)
		if ok {
			rightNode, ok := t.Right.(*ast.ConstNode)
			if ok {
				switch t.Operator {
				case "||":
					res, err := runtime.Or(leftNode.ConstValue, rightNode.ConstValue)
					if err != nil {
						fold.appendError(err)
						return
					}
					*node = &ast.ConstNode{ConstValue: res}
				case "&&":
					res, err := runtime.And(leftNode.ConstValue, rightNode.ConstValue)
					if err != nil {
						fold.appendError(err)
						return
					}
					*node = &ast.ConstNode{ConstValue: res}
				case "!=":
					res, err := runtime.NotEqual(leftNode.ConstValue, rightNode.ConstValue)
					if err != nil {
						fold.appendError(err)
						return
					}
					*node = &ast.ConstNode{ConstValue: res}
				case "==":
					res, err := runtime.Equal(leftNode.ConstValue, rightNode.ConstValue)
					if err != nil {
						fold.appendError(err)
						return
					}
					*node = &ast.ConstNode{ConstValue: res}
				case ">":
					res, err := runtime.GreatThan(leftNode.ConstValue, rightNode.ConstValue)
					if err != nil {
						fold.appendError(err)
						return
					}
					*node = &ast.ConstNode{ConstValue: res}
				case ">=":
					res, err := runtime.GreatThanOrEqual(leftNode.ConstValue, rightNode.ConstValue)
					if err != nil {
						fold.appendError(err)
						return
					}
					*node = &ast.ConstNode{ConstValue: res}
				case "<":
					res, err := runtime.GreatThan(rightNode.ConstValue, leftNode.ConstValue)
					if err != nil {
						fold.appendError(err)
						return
					}
					*node = &ast.ConstNode{ConstValue: res}
				case "<=":
					res, err := runtime.GreatThanOrEqual(rightNode.ConstValue, leftNode.ConstValue)
					if err != nil {
						fold.appendError(err)
						return
					}
					*node = &ast.ConstNode{ConstValue: res}
				case "+":
					res, err := runtime.Add(leftNode.ConstValue, rightNode.ConstValue)
					if err != nil {
						fold.appendError(err)
						return
					}
					*node = &ast.ConstNode{ConstValue: res}
				case "-":
					res, err := runtime.Minus(leftNode.ConstValue, rightNode.ConstValue)
					if err != nil {
						fold.appendError(err)
						return
					}
					*node = &ast.ConstNode{ConstValue: res}
				case "*":
					res, err := runtime.Multiply(leftNode.ConstValue, rightNode.ConstValue)
					if err != nil {
						fold.appendError(err)
						return
					}
					*node = &ast.ConstNode{ConstValue: res}
				case "/":
					res, err := runtime.Divide(leftNode.ConstValue, rightNode.ConstValue)
					if err != nil {
						fold.appendError(err)
						return
					}
					*node = &ast.ConstNode{ConstValue: res}
				case "%":
					res, err := runtime.Mod(leftNode.ConstValue, rightNode.ConstValue)
					if err != nil {
						fold.appendError(err)
						return
					}
					*node = &ast.ConstNode{ConstValue: res}
				}
			}
		}
	}
}
