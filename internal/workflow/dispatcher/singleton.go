package dispatcher

import (
	"sync"
)

var (
	instance    WorkflowDispatcher
	intanceOnce sync.Once
)

func GetSingletonDispatcher() WorkflowDispatcher {
	intanceOnce.Do(func() {
		instance = newWorkflowDispatcher()
	})
	return instance
}
