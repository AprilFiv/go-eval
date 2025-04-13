package visitor

import (
	"fmt"
	"github.com/AprilFiv/go-eval/ast"
	"github.com/AprilFiv/go-eval/runtime"
)

type TypeCheck struct {
	Err error
}

func (typeCheck *TypeCheck) appendError(err error) {
	if typeCheck.Err != nil {
		typeCheck.Err = fmt.Errorf("%w \n%w", typeCheck.Err, err)
		return
	}
	typeCheck.Err = err
}

func (typeCheck *TypeCheck) Visit(node *ast.Node) {
	switch t := (*node).(type) {
	//常数节点，不存在则设置进图内即可
	case *ast.ConstNode:
	case *ast.VariableNode:
	case *ast.FuncNode:
	case *ast.UnaryNode:
		constNode, ok := t.Operand.(*ast.ConstNode)
		if ok {
			typeCheck.checkOperator(constNode, t.Operator)
		}
	case *ast.BinaryNode:
		leftNode, ok := t.Left.(*ast.ConstNode)
		if ok {
			typeCheck.checkOperator(leftNode, t.Operator)
		}
		rightNode, ok := t.Right.(*ast.ConstNode)
		if ok {
			typeCheck.checkOperator(rightNode, t.Operator)
		}

	}
}

func (typeCheck *TypeCheck) checkOperator(constNode *ast.ConstNode, operator string) bool {
	switch operator {
	case "||", "&&", "!":
		if !runtime.ShouldBool(constNode.ConstValue) {
			typeCheck.appendError(fmt.Errorf("operator %s not support %T", operator, constNode.ConstValue))
		}
	case "!=":
	case "==":
	case ">", ">=", "<", "<=":
		if !(runtime.ShouldInt64(constNode.ConstValue) ||
			runtime.ShouldFloat64(constNode.ConstValue) ||
			runtime.ShouldString(constNode.ConstValue)) {
			typeCheck.appendError(fmt.Errorf("operator %s not support %T", operator, constNode.ConstValue))
		}
	case "+", "-", "*", "/", "%":
		if !(runtime.ShouldInt64(constNode.ConstValue) ||
			runtime.ShouldFloat64(constNode.ConstValue)) {
			typeCheck.appendError(fmt.Errorf("operator %s not support %T", operator, constNode.ConstValue))
		}
	}
	return false
}
