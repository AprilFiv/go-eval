package parser

import (
	"strconv"

	"github.com/AprilFiv/go-eval/ast"
	gen "github.com/AprilFiv/go-eval/parser/antlr"
	"github.com/antlr4-go/antlr/v4"
)

type DefaultVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *DefaultVisitor) VisitProgram(ctx *gen.ProgramContext) interface{} {
	return ctx.Statement().Accept(v)
}

func (v *DefaultVisitor) VisitBool(ctx *gen.BoolContext) interface{} {
	value, err := strconv.ParseBool(ctx.GetText())
	if err != nil {
		return err
	}
	return &ast.ConstNode{ConstValue: value}
}

func (v *DefaultVisitor) VisitString(ctx *gen.StringContext) interface{} {
	value := ctx.GetText()[1 : len(ctx.GetText())-1]
	return &ast.ConstNode{ConstValue: value}
}

func (v *DefaultVisitor) VisitInteger(ctx *gen.IntegerContext) interface{} {
	value, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		return err
	}
	return &ast.ConstNode{ConstValue: value}
}

func (v *DefaultVisitor) VisitFloat(ctx *gen.FloatContext) interface{} {
	value, err := strconv.ParseFloat(ctx.GetText(), 10)
	if err != nil {
		return err
	}
	return &ast.ConstNode{ConstValue: value}
}

func (v *DefaultVisitor) VisitV(ctx *gen.VContext) interface{} {
	return ctx.Variable().Accept(v)
}

func (v *DefaultVisitor) VisitMulDiv(ctx *gen.MulDivContext) interface{} {
	return v.VisitBinary(ctx)
}
func (v *DefaultVisitor) VisitAddSUb(ctx *gen.AddSUbContext) interface{} {
	return v.VisitBinary(ctx)
}
func (v *DefaultVisitor) VisitComparison1(ctx *gen.Comparison1Context) interface{} {
	return v.VisitBinary(ctx)
}
func (v *DefaultVisitor) VisitComparison2(ctx *gen.Comparison2Context) interface{} {
	return v.VisitBinary(ctx)
}
func (v *DefaultVisitor) VisitAnd(ctx *gen.AndContext) interface{} {
	return v.VisitBinary(ctx)
}
func (v *DefaultVisitor) VisitOr(ctx *gen.OrContext) interface{} {
	return v.VisitBinary(ctx)
}

func (v *DefaultVisitor) VisitUnary(ctx *gen.UnaryContext) interface{} {
	child := ctx.Statement().Accept(v)
	err, ok := child.(error)
	if ok {
		return err
	}
	return &ast.UnaryNode{
		Operator: ctx.GetOp().GetText(),
		Operand:  child.(ast.Node),
	}
}

func (v *DefaultVisitor) VisitVar(ctx *gen.VarContext) interface{} {
	return &ast.VariableNode{
		VariableName: ctx.GetText(),
	}
}

func (v *DefaultVisitor) VisitF(ctx *gen.FContext) interface{} {
	//F是function别名
	return ctx.Function().Accept(v)
}

func (v *DefaultVisitor) VisitParens(ctx *gen.ParensContext) interface{} {
	//直接忽视括号
	return ctx.Statement().Accept(v)
}

type IBinaryContext interface {
	Statement(i int) gen.IStatementContext
	GetOp() antlr.Token
}

func (v *DefaultVisitor) VisitBinary(ctx IBinaryContext) interface{} {
	leftChild := ctx.Statement(0).Accept(v)
	err, ok := leftChild.(error)
	if ok {
		return err
	}
	rightChild := ctx.Statement(1).Accept(v)
	err, ok = rightChild.(error)
	if ok {
		return err
	}

	return &ast.BinaryNode{
		Operator: ctx.GetOp().GetText(),
		Left:     leftChild.(ast.Node),
		Right:    rightChild.(ast.Node),
	}
}

func (v *DefaultVisitor) VisitFunction(ctx *gen.FunctionContext) interface{} {
	var childNodes []ast.Node
	for _, variable := range ctx.AllStatement() {
		child := variable.Accept(v)
		err, ok := child.(error)
		if ok {
			return err
		}
		childNodes = append(childNodes, child.(ast.Node))
	}
	return &ast.FuncNode{
		FuncName: ctx.GetF().GetText(),
		Args:     childNodes,
	}
}
