package visitor

import (
	"github.com/AprilFiv/go-eval/ast"
	"github.com/AprilFiv/go-eval/rete"
)

/*
*
可以对ast.Node做表达式压缩做BetaNode,这里不做实现。直接存储ast.Node结果。
*/
type ReteVisitor struct {
	AlphaNodes            map[string]*rete.ReteNode
	IndexAllocator        int64
	GraphNodes            map[string]*rete.ReteNode
	InitializationContext []*rete.NodeContext
}

func (reteVisitor *ReteVisitor) Visit(node *ast.Node) {
	if reteVisitor.GraphNodes[rete.Identifier(node)] != nil {
		return
	}

	idx := reteVisitor.IndexAllocator
	reteNode := &rete.ReteNode{
		Idx:  idx,
		Node: *node,
	}
	reteVisitor.GraphNodes[rete.Identifier(node)] = reteNode
	reteVisitor.IndexAllocator++
	reteVisitor.InitializationContext = append(reteVisitor.InitializationContext, &rete.NodeContext{})

	switch typedNode := (*node).(type) {
	//常数节点，不存在则设置进图内即可
	case *ast.ConstNode:
		reteVisitor.InitializationContext[idx].Value = typedNode.ConstValue
		reteVisitor.InitializationContext[idx].Done = true
		//变量节点，不存在时，要加入到AlphaNode列表中(VariableNode即AlphaNode
	case *ast.VariableNode:
		reteVisitor.AlphaNodes[typedNode.VariableName] = reteNode
		//函数节点是上层节点，要设置和子节点的连通关系
	case *ast.FuncNode:
		for _, arg := range typedNode.Args {
			child := reteVisitor.GraphNodes[rete.Identifier(&arg)]
			reteVisitor.link(reteNode, child)
		}
		//一元节点是上层节点，要设置和子节点的连通关系
	case *ast.UnaryNode:
		child := reteVisitor.GraphNodes[rete.Identifier(&typedNode.Operand)]
		reteVisitor.link(reteNode, child)
		//二元节点是上层节点，要设置和子节点的连通关系
	case *ast.BinaryNode:
		left := reteVisitor.GraphNodes[rete.Identifier(&typedNode.Left)]
		right := reteVisitor.GraphNodes[rete.Identifier(&typedNode.Right)]
		reteVisitor.link(reteNode, left)
		reteVisitor.link(reteNode, right)
	}

}

func (reteVisitor *ReteVisitor) link(parent *rete.ReteNode, child *rete.ReteNode) {
	parent.SetSuccessor(append(parent.GetSuccessor(), child))
	child.SetPredecessor(append(child.GetPredecessor(), parent))
}
