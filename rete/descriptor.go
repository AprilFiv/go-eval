package rete

import (
	"context"
)

type Descriptor struct {
	AlphaNodes            map[string]*ReteNode
	ConstOutput           map[string]*NodeContext
	InitializationContext []*NodeContext
}

func (d *Descriptor) Context() *EvalContext {
	nodeContext := make([]*NodeContext, len(d.InitializationContext))
	for idx, _ := range d.InitializationContext {
		if d.InitializationContext[idx].Done {
			nodeContext[idx] = d.InitializationContext[idx]
		}
	}

	output := make(map[string]*NodeContext, len(d.ConstOutput))
	for name, ctx := range d.ConstOutput {
		output[name] = ctx
	}
	return &EvalContext{
		Output:      output,
		NodeContext: nodeContext,
	}
}

func (d *Descriptor) Run(ctx context.Context, c *EvalContext, name string, value interface{}, err error) bool {
	leaf, ok := d.AlphaNodes[name]
	if !ok {
		return false
	}
	//可能被剪枝了，直接return
	lc := leaf.Context(ctx, c.NodeContext)
	if lc.Done {
		return false
	}
	lc.Value = value
	lc.Err = err

	d.processNode(ctx, c, leaf)
	return true
}

func (d *Descriptor) processNode(ctx context.Context, c *EvalContext, node *ReteNode) {
	if node.DoOperator(ctx, c.NodeContext) {
		for _, oName := range node.Activation {
			nc := node.Context(ctx, c.NodeContext)
			c.Output[oName] = nc
		}

		for _, predecessor := range node.GetPredecessor() {
			d.processNode(ctx, c, predecessor)
		}
	}
}
