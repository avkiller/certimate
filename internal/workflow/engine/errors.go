package engine

import (
	"errors"
)

var (
	// 表示工作流引擎执行被中断，可能已结束
	ErrTerminated = errors.New("workflow engine: execution was terminated")
	// 表示工作流引擎在执行子节点时发生异常
	ErrBlocksException = errors.New("workflow engine: error occurred when executing blocks")
)
