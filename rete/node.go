package rete

import (
	"context"
	"fmt"
	"github.com/AprilFiv/go-eval/ast"
	"github.com/AprilFiv/go-eval/runtime"
)

type ReteNode struct {
	Idx int64
	ast.Node
	Activation  []string
	Predecessor []*ReteNode
	Successor   []*ReteNode
}

func Identifier(node *ast.Node) string {
	switch typedNode := (*node).(type) {
	//常数节点，不存在则设置进图内即可
	case *ast.ConstNode:
		return fmt.Sprintf("%s|%T|%v", "const", typedNode.ConstValue, typedNode.ConstValue)
	case *ast.VariableNode:
		return fmt.Sprintf("%s|%s", "variable", typedNode.VariableName)
	case *ast.FuncNode:
		s := typedNode.FuncName + "("
		for i, arg := range typedNode.Args {
			s = s + Identifier(&arg)
			if i != len(typedNode.Args)-1 {
				s = s + ","
			}
		}
		s = s + ")"
		return fmt.Sprintf("%s|%s", "func", s)
	case *ast.BinaryNode:
		return fmt.Sprintf("%s|%s[%s,%s]", "binary", typedNode.Operator, Identifier(&typedNode.Left), Identifier(&typedNode.Right))
	case *ast.UnaryNode:
		return fmt.Sprintf("%s|%s[%s]", "unary", typedNode.Operator, Identifier(&typedNode.Operand))
	}
	return ""
}

func (rn *ReteNode) GetPredecessor() []*ReteNode {
	return rn.Predecessor
}

func (rn *ReteNode) SetPredecessor(p []*ReteNode) {
	rn.Predecessor = p

}
func (rn *ReteNode) GetSuccessor() []*ReteNode {
	return rn.Successor
}

func (rn *ReteNode) SetSuccessor(s []*ReteNode) {
	rn.Successor = s
}

func (rn *ReteNode) Context(ctx context.Context, contextMap []*NodeContext) *NodeContext {
	nc := contextMap[rn.Idx]
	if nc == nil {
		nc = &NodeContext{}
		contextMap[rn.Idx] = nc
	}
	return nc
}

func (rn *ReteNode) CutOff(ctx context.Context, contextMap []*NodeContext) {
	//存在激活，则不进行剪枝
	if len(rn.Activation) != 0 {
		return
	}
	nc := rn.Context(ctx, contextMap)
	if nc.Done {
		return
	}

	//存在没有完成的父节点
	for _, p := range rn.Predecessor {
		pc := p.Context(ctx, contextMap)
		if !pc.Done {
			return
		}
	}

	nc.Done = true

	//所有的子节点也全部剪枝
	for _, s := range rn.Successor {
		s.CutOff(ctx, contextMap)
	}
}

func (rn *ReteNode) DoOperator(ctx context.Context, contextMap []*NodeContext) bool {
	nc := rn.Context(ctx, contextMap)
	if nc.Done {
		return false
	}

	switch typedNode := (rn.Node).(type) {
	case *ast.ConstNode:
		nc.Value = typedNode.ConstValue
		nc.Done = true
		return true
		//ast中的variableNode即alphaNode，外部已经完成值的填充
	case *ast.VariableNode:
		nc.Done = true
		return true
	case *ast.FuncNode:
		var params []interface{}
		for index, successor := range rn.Successor {
			sc := successor.Context(ctx, contextMap)
			if sc.Err != nil {
				nc.Err = sc.Err
				nc.Done = true
				for i := 0; i < index; i++ {
					rn.Successor[i].CutOff(ctx, contextMap)
				}
				for i := index + 1; i < len(rn.Successor); i++ {
					rn.Successor[i].CutOff(ctx, contextMap)
				}
				return true
			}

			if !sc.Done {
				return false
			}
			params = append(params, sc.Value)
		}

		res, err := typedNode.Executable(ctx, params...)
		nc.Done = true
		nc.Err = err
		nc.Value = res
		return true
	case *ast.BinaryNode:

		leftOperand := rn.Successor[0]
		rightOperand := rn.Successor[1]
		lc := leftOperand.Context(ctx, contextMap)
		rc := rightOperand.Context(ctx, contextMap)
		switch typedNode.Operator {

		//对于不存在短路特性的二元操作，若一个算子发生错误，则可直接返回

		case "+":
			if lc.Err != nil {
				nc.Err = lc.Err
				nc.Done = true
				rightOperand.CutOff(ctx, contextMap)
				return true
			}
			if rc.Err != nil {
				nc.Err = rc.Err
				nc.Done = true
				leftOperand.CutOff(ctx, contextMap)
				return true
			}

			leftValue := lc.Value
			rightValue := rc.Value

			if leftValue != nil && rightValue != nil {
				res, err := runtime.Add(leftValue, rightValue)
				nc.Done = true
				nc.Err = err
				nc.Value = res
				return true
			}
		case "-":
			if lc.Err != nil {
				nc.Err = lc.Err
				nc.Done = true
				rightOperand.CutOff(ctx, contextMap)
				return true
			}
			if rc.Err != nil {
				nc.Err = rc.Err
				nc.Done = true
				leftOperand.CutOff(ctx, contextMap)
				return true
			}

			leftValue := lc.Value
			rightValue := rc.Value

			if leftValue != nil && rightValue != nil {
				res, err := runtime.Minus(leftValue, rightValue)
				nc.Done = true
				nc.Err = err
				nc.Value = res
				return true
			}
		case "*":
			if lc.Err != nil {
				nc.Err = lc.Err
				nc.Done = true
				rightOperand.CutOff(ctx, contextMap)
				return true
			}
			if rc.Err != nil {
				nc.Err = rc.Err
				nc.Done = true
				leftOperand.CutOff(ctx, contextMap)
				return true
			}

			leftValue := lc.Value
			rightValue := rc.Value

			if leftValue != nil && rightValue != nil {
				res, err := runtime.Multiply(leftValue, rightValue)
				nc.Done = true
				nc.Err = err
				nc.Value = res
				return true
			}
		case "/":
			if lc.Err != nil {
				nc.Err = lc.Err
				nc.Done = true
				rightOperand.CutOff(ctx, contextMap)
				return true
			}
			if rc.Err != nil {
				nc.Err = rc.Err
				nc.Done = true
				leftOperand.CutOff(ctx, contextMap)
				return true
			}

			leftValue := lc.Value
			rightValue := rc.Value

			if leftValue != nil && rightValue != nil {
				res, err := runtime.Divide(leftValue, rightValue)
				nc.Done = true
				nc.Err = err
				nc.Value = res
				return true
			}
		case "%":
			if lc.Err != nil {
				nc.Err = lc.Err
				nc.Done = true
				rightOperand.CutOff(ctx, contextMap)
				return true
			}
			if rc.Err != nil {
				nc.Err = rc.Err
				nc.Done = true
				leftOperand.CutOff(ctx, contextMap)
				return true
			}

			leftValue := lc.Value
			rightValue := rc.Value

			if leftValue != nil && rightValue != nil {
				res, err := runtime.Mod(leftValue, rightValue)
				nc.Done = true
				nc.Err = err
				nc.Value = res
				return true
			}
		case "==":
			if lc.Err != nil {
				nc.Err = lc.Err
				nc.Done = true
				rightOperand.CutOff(ctx, contextMap)
				return true
			}
			if rc.Err != nil {
				nc.Err = rc.Err
				nc.Done = true
				leftOperand.CutOff(ctx, contextMap)
				return true
			}

			leftValue := lc.Value
			rightValue := rc.Value

			if leftValue != nil && rightValue != nil {
				res, err := runtime.Equal(leftValue, rightValue)
				nc.Done = true
				nc.Err = err
				nc.Value = res
				return true
			}
		case "!=":
			if lc.Err != nil {
				nc.Err = lc.Err
				nc.Done = true
				rightOperand.CutOff(ctx, contextMap)
				return true
			}
			if rc.Err != nil {
				nc.Err = rc.Err
				nc.Done = true
				leftOperand.CutOff(ctx, contextMap)
				return true
			}

			leftValue := lc.Value
			rightValue := rc.Value

			if leftValue != nil && rightValue != nil {
				res, err := runtime.NotEqual(leftValue, rightValue)
				nc.Done = true
				nc.Err = err
				nc.Value = res
				return true
			}
		case ">":
			if lc.Err != nil {
				nc.Err = lc.Err
				nc.Done = true
				rightOperand.CutOff(ctx, contextMap)
				return true
			}
			if rc.Err != nil {
				nc.Err = rc.Err
				nc.Done = true
				leftOperand.CutOff(ctx, contextMap)
				return true
			}

			leftValue := lc.Value
			rightValue := rc.Value

			if leftValue != nil && rightValue != nil {
				res, err := runtime.GreatThan(leftValue, rightValue)
				nc.Done = true
				nc.Err = err
				nc.Value = res
				return true
			}
		case ">=":
			if lc.Err != nil {
				nc.Err = lc.Err
				nc.Done = true
				rightOperand.CutOff(ctx, contextMap)
				return true
			}
			if rc.Err != nil {
				nc.Err = rc.Err
				nc.Done = true
				leftOperand.CutOff(ctx, contextMap)
				return true
			}

			leftValue := lc.Value
			rightValue := rc.Value

			if leftValue != nil && rightValue != nil {
				res, err := runtime.GreatThanOrEqual(leftValue, rightValue)
				nc.Done = true
				nc.Err = err
				nc.Value = res
				return true
			}
		case "<":
			if lc.Err != nil {
				nc.Err = lc.Err
				nc.Done = true
				rightOperand.CutOff(ctx, contextMap)
				return true
			}
			if rc.Err != nil {
				nc.Err = rc.Err
				nc.Done = true
				leftOperand.CutOff(ctx, contextMap)
				return true
			}

			leftValue := lc.Value
			rightValue := rc.Value

			if leftValue != nil && rightValue != nil {
				res, err := runtime.LessThan(leftValue, rightValue)
				nc.Done = true
				nc.Err = err
				nc.Value = res
				return true
			}
		case "<=":
			if lc.Err != nil {
				nc.Err = lc.Err
				nc.Done = true
				rightOperand.CutOff(ctx, contextMap)
				return true
			}
			if rc.Err != nil {
				nc.Err = rc.Err
				nc.Done = true
				leftOperand.CutOff(ctx, contextMap)
				return true
			}

			leftValue := lc.Value
			rightValue := rc.Value

			if leftValue != nil && rightValue != nil {
				res, err := runtime.LessThanOrEqual(leftValue, rightValue)
				nc.Done = true
				nc.Err = err
				nc.Value = res
				return true
			}

		case "&&":
			leftValue := lc.Value
			//任意一个为false，表达式为false，则返回
			if leftValue != nil && !leftValue.(bool) && lc.Err == nil {
				nc.Done = true
				nc.Value = false
				rightOperand.CutOff(ctx, contextMap)
				return true
			}

			rightValue := rc.Value
			if rightValue != nil && !rightValue.(bool) && rc.Err == nil {
				nc.Done = true
				nc.Value = false
				leftOperand.CutOff(ctx, contextMap)
				return true
			}

			//一个有错误，另一个完成了，则返回错误
			if lc.Err != nil && rc.Done {
				nc.Done = true
				nc.Err = lc.Err
				return true
			}
			if rc.Err != nil && lc.Done {
				nc.Done = true
				nc.Err = rc.Err
				return true
			}

			if leftValue != nil && rightValue != nil {
				res, err := runtime.And(leftValue, rightValue)
				nc.Done = true
				nc.Err = err
				nc.Value = res
				return true
			}
		case "||":
			leftValue := lc.Value
			//任意一个为true，表达式为true，则返回
			if leftValue != nil && leftValue.(bool) && lc.Err == nil {
				nc.Done = true
				nc.Value = true
				rightOperand.CutOff(ctx, contextMap)
				return true
			}

			rightValue := rc.Value
			if rightValue != nil && rightValue.(bool) && rc.Err == nil {
				nc.Done = true
				nc.Value = true
				leftOperand.CutOff(ctx, contextMap)
				return true
			}

			//一个有错误，另一个完成了，则返回错误
			if lc.Err != nil && rc.Done {
				nc.Done = true
				nc.Err = lc.Err
				return true
			}
			if rc.Err != nil && lc.Done {
				nc.Done = true
				nc.Err = rc.Err
				return true
			}

			if leftValue != nil && rightValue != nil {
				res, err := runtime.Or(leftValue, rightValue)
				nc.Done = true
				nc.Err = err
				nc.Value = res
				return true
			}
		}
	case *ast.UnaryNode:
		operand := rn.Successor[0]
		oc := operand.Context(ctx, contextMap)

		switch typedNode.Operator {
		case "-":
			if oc.Err != nil {
				nc.Done = true
				nc.Err = oc.Err
				return true
			}
			value := oc.Value
			if value != nil {
				res, err := runtime.Negative(value)
				nc.Done = true
				nc.Err = err
				nc.Value = res
				return true
			}
		case "!":
			if oc.Err != nil {
				nc.Done = true
				nc.Err = oc.Err
				return true
			}
			value := oc.Value
			if value != nil {
				res, err := runtime.Not(value)
				nc.Done = true
				nc.Err = err
				nc.Value = res
				return true
			}
		}

	}
	return false

}
