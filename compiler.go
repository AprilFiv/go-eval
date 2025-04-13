package eval

import (
	"context"
	"errors"
	"fmt"
	"github.com/AprilFiv/go-eval/rete"

	"github.com/AprilFiv/go-eval/ast"
	"github.com/AprilFiv/go-eval/parser"
	gen "github.com/AprilFiv/go-eval/parser/antlr"
	"github.com/AprilFiv/go-eval/visitor"
	"github.com/antlr4-go/antlr/v4"
)

func Compile(expr string) (ast.Node, error) {

	is := antlr.NewInputStream(expr)

	// Create the Lexer
	lexer := gen.NewevalLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := gen.NewevalParser(stream)

	errorListener := &parser.DefaultErrorListener{}
	p.AddErrorListener(errorListener)
	parseResult := p.Program().Accept(&parser.DefaultVisitor{})

	if errorListener.Error != "" {
		return nil, errors.New(errorListener.Error)
	}

	err, ok := parseResult.(error)
	if ok {
		return nil, err
	}

	return parseResult.(ast.Node), nil

}

type Options struct {
	Resolver func(funcName string) (func(ctx context.Context, args ...interface{}) (interface{}, error), error)
}

type ParseOption func(opt *Options)

func WithResolver(resolver func(funcName string) (func(ctx context.Context, args ...interface{}) (interface{}, error), error)) func(o *Options) {
	return func(opt *Options) {
		opt.Resolver = resolver
	}
}

func Parse(ctx context.Context, expressions map[string]string, parseOptions ...ParseOption) (*rete.Descriptor, error) {

	opt := &Options{}

	for _, parseOption := range parseOptions {
		parseOption(opt)
	}

	//rete visitor是面向全部expression的
	reteVisitor := &visitor.ReteVisitor{
		AlphaNodes:            map[string]*rete.ReteNode{},
		GraphNodes:            map[string]*rete.ReteNode{},
		InitializationContext: []*rete.NodeContext{},
	}

	constOutput := map[string]*rete.NodeContext{}

	/**
	1.type check
	2.func resolve
	3.const fold
	4.rete
	*/

	for name, expr := range expressions {
		root, err := Compile(expr)
		if err != nil {
			return nil, fmt.Errorf("%w \nname:%s expr:%s", err, name, expr)
		}

		typeCheck := &visitor.TypeCheck{}
		ast.Walk(&root, typeCheck)
		if typeCheck.Err != nil {
			return nil, fmt.Errorf("%w \nname:%s expr:%s", typeCheck.Err, name, expr)
		}

		funcResolve := &visitor.FuncResolve{
			Resolver: opt.Resolver,
		}
		ast.Walk(&root, funcResolve)
		if funcResolve.Err != nil {
			return nil, fmt.Errorf("%w \nname:%s expr:%s", funcResolve.Err, name, expr)
		}

		fold := &visitor.Folding{}
		ast.Walk(&root, fold)
		if fold.Err != nil {
			return nil, fmt.Errorf("%w \nname:%s expr:%s", fold.Err, name, expr)
		}

		constNode, ok := root.(*ast.ConstNode)
		if ok {
			constOutput[name] = &rete.NodeContext{
				Done:  true,
				Err:   nil,
				Value: constNode.ConstValue,
			}
			continue
		}

		ast.Walk(&root, reteVisitor)

		reteOutputNode := reteVisitor.GraphNodes[rete.Identifier(&root)]
		reteOutputNode.Activation = append(reteOutputNode.Activation, name)
	}

	return &rete.Descriptor{
		AlphaNodes:            reteVisitor.AlphaNodes,
		ConstOutput:           constOutput,
		InitializationContext: reteVisitor.InitializationContext,
	}, nil

}
