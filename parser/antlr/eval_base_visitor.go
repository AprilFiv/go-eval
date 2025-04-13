// Code generated from eval.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // eval

import "github.com/antlr4-go/antlr/v4"

type BaseevalVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseevalVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseevalVisitor) VisitBool(ctx *BoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseevalVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseevalVisitor) VisitFloat(ctx *FloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseevalVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseevalVisitor) VisitVar(ctx *VarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseevalVisitor) VisitOr(ctx *OrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseevalVisitor) VisitComparison1(ctx *Comparison1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseevalVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseevalVisitor) VisitAddSUb(ctx *AddSUbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseevalVisitor) VisitComparison2(ctx *Comparison2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseevalVisitor) VisitV(ctx *VContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseevalVisitor) VisitF(ctx *FContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseevalVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseevalVisitor) VisitAnd(ctx *AndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseevalVisitor) VisitUnary(ctx *UnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseevalVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}
