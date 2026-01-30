package engine

import (
	"log/slog"
)

type endNodeExecutor struct {
	nodeExecutor
}

func (ne *endNodeExecutor) Execute(execCtx *NodeExecutionContext) (*NodeExecutionResult, error) {
	execRes := newNodeExecutionResult(execCtx.Node)
	execRes.Terminated = true

	ne.logger.Info("the workflow is ending")

	return execRes, nil
}

func newEndNodeExecutor() NodeExecutor {
	return &endNodeExecutor{
		nodeExecutor: nodeExecutor{logger: slog.Default()},
	}
}
