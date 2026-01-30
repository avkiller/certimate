package engine

import (
	"context"
	"log/slog"
	"sync"
)

type NodeExecutor interface {
	withLogger

	Execute(execCtx *NodeExecutionContext) (*NodeExecutionResult, error)
}

type nodeExecutor struct {
	logger *slog.Logger
}

func (e *nodeExecutor) SetLogger(logger *slog.Logger) {
	e.logger = logger
}

type NodeExecutionContext struct {
	WorkflowContext

	Node *Node
}

func (c *NodeExecutionContext) SetExecutingWorkflow(workflowId string, runId string, runGraph *Graph) *NodeExecutionContext {
	c.WorkflowContext.SetExecutingWorkflow(workflowId, runId, runGraph)
	return c
}

func (c *NodeExecutionContext) SetExecutingNode(node *Node) *NodeExecutionContext {
	c.Node = node
	return c
}

func (c *NodeExecutionContext) SetEngine(engine WorkflowEngine) *NodeExecutionContext {
	c.WorkflowContext.SetEngine(engine)
	return c
}

func (c *NodeExecutionContext) SetVariablesManager(variables VariableManager) *NodeExecutionContext {
	c.WorkflowContext.SetVariablesManager(variables)
	return c
}

func (c *NodeExecutionContext) SetInputsManager(inputs InOutManager) *NodeExecutionContext {
	c.WorkflowContext.SetInputsManager(inputs)
	return c
}

func (c *NodeExecutionContext) SetContext(ctx context.Context) *NodeExecutionContext {
	c.WorkflowContext.SetContext(ctx)
	return c
}

func newNodeExecutionContext(wfCtx *WorkflowContext, node *Node) *NodeExecutionContext {
	return (&NodeExecutionContext{}).
		SetExecutingWorkflow(wfCtx.WorkflowId, wfCtx.RunId, wfCtx.RunGraph).
		SetExecutingNode(node).
		SetEngine(wfCtx.engine).
		SetVariablesManager(wfCtx.variables).
		SetInputsManager(wfCtx.inputs).
		SetContext(wfCtx.ctx)
}

type NodeExecutionResult struct {
	node *Node

	Terminated bool // 是否终止执行（通常由 End 节点主动触发）

	variablesMtx sync.Mutex
	Variables    []VariableState

	outputForced bool // 即使 Outputs 为空，也强制持久化输出
	outputsMtx   sync.Mutex
	Outputs      []InOutState
}

func (r *NodeExecutionResult) AddVariable(key string, value any, valueType string) {
	r.AddVariableWithScope("", key, value, valueType)
}

func (r *NodeExecutionResult) AddVariableWithScope(scope string, key string, value any, valueType string) {
	r.addVariableState(VariableState{
		Scope:     scope,
		Key:       key,
		Value:     value,
		ValueType: valueType,
	})
}

func (r *NodeExecutionResult) addVariableState(state VariableState) {
	r.variablesMtx.Lock()
	defer r.variablesMtx.Unlock()

	if r.Variables == nil {
		r.Variables = make([]VariableState, 0)
	}

	for i, item := range r.Variables {
		if item.Scope == state.Scope && item.Key == state.Key {
			r.Variables[i] = state
			return
		}
	}
	r.Variables = append(r.Variables, state)
}

func (r *NodeExecutionResult) AddOutput(stype string, key string, value any, valueType string) {
	r.addOutputState(InOutState{
		NodeId:     r.node.Id,
		Type:       stype,
		Name:       key,
		Value:      value,
		ValueType:  valueType,
		Persistent: false,
	})
}

func (r *NodeExecutionResult) AddOutputWithPersistent(stype string, key string, value any, valueType string) {
	r.addOutputState(InOutState{
		NodeId:     r.node.Id,
		Type:       stype,
		Name:       key,
		Value:      value,
		ValueType:  valueType,
		Persistent: true,
	})
}

func (r *NodeExecutionResult) addOutputState(state InOutState) {
	r.outputsMtx.Lock()
	defer r.outputsMtx.Unlock()

	if r.Outputs == nil {
		r.Outputs = make([]InOutState, 0)
	}

	for i, t := range r.Outputs {
		if t.NodeId == state.NodeId && t.Name == state.Name {
			r.Outputs[i] = state
			return
		}
	}
	r.Outputs = append(r.Outputs, state)
}

func newNodeExecutionResult(node *Node) *NodeExecutionResult {
	return &NodeExecutionResult{
		node:      node,
		Variables: make([]VariableState, 0),
		Outputs:   make([]InOutState, 0),
	}
}
