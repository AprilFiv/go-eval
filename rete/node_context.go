package rete

type NodeContext struct {
	Done  bool
	Err   error
	Value interface{}
}
