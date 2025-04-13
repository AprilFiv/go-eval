package ast

import (
	"context"
)

type Node interface{}

type ConstNode struct {
	Node
	ConstValue interface{}
}

type FuncNode struct {
	Node
	FuncName   string
	Executable func(ctx context.Context, args ...interface{}) (interface{}, error)
	Args       []Node
}

type VariableNode struct {
	Node
	VariableName string
}

type BinaryNode struct {
	Node
	Left     Node
	Right    Node
	Operator string
}

type UnaryNode struct {
	Node
	Operand  Node
	Operator string
}
