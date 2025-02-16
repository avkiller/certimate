package nodeprocessor

import (
	"context"

	"github.com/usual2970/certimate/internal/domain"
)

type executeFailureNode struct {
	node *domain.WorkflowNode
	*nodeLogger
}

func NewExecuteFailureNode(node *domain.WorkflowNode) *executeFailureNode {
	return &executeFailureNode{
		node:       node,
		nodeLogger: newNodeLogger(node),
	}
}

func (n *executeFailureNode) Process(ctx context.Context) error {
	// 此类型节点不需要执行任何操作，直接返回
	n.AppendLogRecord(ctx, domain.WorkflowRunLogLevelInfo, "进入执行失败分支")

	return nil
}
