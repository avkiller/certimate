package engine

import (
	"context"
)

type WorkflowContext struct {
	WorkflowId string
	RunId      string
	RunGraph   *Graph

	engine    WorkflowEngine
	variables VariableManager
	inputs    InOutManager

	ctx context.Context
}

func (c *WorkflowContext) SetExecutingWorkflow(workflowId string, runId string, runGraph *Graph) *WorkflowContext {
	c.WorkflowId = workflowId
	c.RunId = runId
	c.RunGraph = runGraph
	return c
}

func (c *WorkflowContext) SetEngine(engine WorkflowEngine) *WorkflowContext {
	c.engine = engine
	return c
}

func (c *WorkflowContext) SetVariablesManager(inputs VariableManager) *WorkflowContext {
	c.variables = inputs
	return c
}

func (c *WorkflowContext) SetInputsManager(manager InOutManager) *WorkflowContext {
	c.inputs = manager
	return c
}

func (c *WorkflowContext) SetContext(ctx context.Context) *WorkflowContext {
	c.ctx = ctx
	return c
}

func (c *WorkflowContext) Context() context.Context {
	return c.ctx
}

func (c *WorkflowContext) Clone() *WorkflowContext {
	return &WorkflowContext{
		WorkflowId: c.WorkflowId,
		RunId:      c.RunId,
		RunGraph:   c.RunGraph,

		engine:    c.engine,
		variables: c.variables,
		inputs:    c.inputs,

		ctx: c.ctx,
	}
}
