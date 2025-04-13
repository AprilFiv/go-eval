package visitor

import (
	"fmt"
	"github.com/AprilFiv/go-eval/ast"
	"github.com/AprilFiv/go-eval/runtime"
)

type FuncResolve struct {
	Resolver func(funcName string) (runtime.Executable, error)
	Err      error
}

func (funcResolve *FuncResolve) appendError(err error) {
	if funcResolve.Err != nil {
		funcResolve.Err = fmt.Errorf("%w \n%w", funcResolve.Err, err)
		return
	}
	funcResolve.Err = err
}

func (funcResolve *FuncResolve) Visit(node *ast.Node) {
	switch t := (*node).(type) {
	case *ast.FuncNode:
		executable, err := funcResolve.Resolver(t.FuncName)
		if err != nil {
			funcResolve.appendError(err)
			return
		}
		t.Executable = executable
	}
}
