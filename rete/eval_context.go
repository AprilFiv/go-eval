package rete

type EvalContext struct {
	NodeContext []*NodeContext
	Output      map[string]*NodeContext
}
