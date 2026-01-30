package dispatcher

import (
	"context"
)

type taskInfo struct {
	WorkflowId string
	RunId      string

	ctx    context.Context
	cancel context.CancelFunc
}
