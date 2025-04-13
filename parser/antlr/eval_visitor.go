// Code generated from eval.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // eval

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by evalParser.
type evalVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by evalParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by evalParser#Bool.
	VisitBool(ctx *BoolContext) interface{}

	// Visit a parse tree produced by evalParser#String.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by evalParser#Float.
	VisitFloat(ctx *FloatContext) interface{}

	// Visit a parse tree produced by evalParser#Integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by evalParser#Var.
	VisitVar(ctx *VarContext) interface{}

	// Visit a parse tree produced by evalParser#Or.
	VisitOr(ctx *OrContext) interface{}

	// Visit a parse tree produced by evalParser#Comparison1.
	VisitComparison1(ctx *Comparison1Context) interface{}

	// Visit a parse tree produced by evalParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by evalParser#AddSUb.
	VisitAddSUb(ctx *AddSUbContext) interface{}

	// Visit a parse tree produced by evalParser#Comparison2.
	VisitComparison2(ctx *Comparison2Context) interface{}

	// Visit a parse tree produced by evalParser#V.
	VisitV(ctx *VContext) interface{}

	// Visit a parse tree produced by evalParser#F.
	VisitF(ctx *FContext) interface{}

	// Visit a parse tree produced by evalParser#Parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by evalParser#And.
	VisitAnd(ctx *AndContext) interface{}

	// Visit a parse tree produced by evalParser#Unary.
	VisitUnary(ctx *UnaryContext) interface{}

	// Visit a parse tree produced by evalParser#function.
	VisitFunction(ctx *FunctionContext) interface{}
}
