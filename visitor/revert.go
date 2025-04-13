package visitor

import (
	"encoding/json"
	"fmt"
	"github.com/AprilFiv/go-eval/ast"
	"strings"
)

type Revert struct {
	expression string
	opStack
}

type opStack struct {
	data []string
}

func (s *opStack) push(op string) {
	s.data = append(s.data, op)
}

func (s *opStack) pop() string {
	res := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return res
}

func (s *opStack) len() int {
	return len(s.data)
}

func (revert *Revert) ToExpression() string {
	return revert.pop()
}

func (revert *Revert) Visit(node *ast.Node) {
	switch t := (*node).(type) {
	case *ast.ConstNode:
		plainValue, _ := json.Marshal(t.ConstValue)
		revert.push(string(plainValue))
	case *ast.VariableNode:
		revert.push(t.VariableName)
	case *ast.FuncNode:
		args := make([]string, len(t.Args))
		for i, _ := range t.Args {
			args[len(args)-i-1] = revert.pop()
		}
		expression := fmt.Sprintf("%s(%s)", t.FuncName, strings.Join(args, ","))
		revert.push(expression)
	case *ast.BinaryNode:
		rightOp := revert.pop()
		leftOp := revert.pop()
		expression := fmt.Sprintf(" (%s %s %s) ", leftOp, t.Operator, rightOp)
		revert.push(expression)
	case *ast.UnaryNode:
		op := revert.pop()
		expression := fmt.Sprintf(" (%s %s) ", t.Operator, op)
		revert.push(expression)

	}
}
