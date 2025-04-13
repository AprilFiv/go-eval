package ast

type Visitor interface {
	Visit(node *Node)
}

func Walk(node *Node, v Visitor) {
	switch n := (*node).(type) {
	case *VariableNode:
	case *ConstNode:
	case *UnaryNode:
		Walk(&n.Operand, v)
	case *BinaryNode:
		Walk(&n.Left, v)
		Walk(&n.Right, v)
	case *FuncNode:
		for i := range n.Args {
			Walk(&n.Args[i], v)
		}
	}
	v.Visit(node)
}
