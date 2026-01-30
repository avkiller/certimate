package domain

import (
	"log/slog"
	"strings"
)

const CollectionNameWorkflowLog = "workflow_logs"

type WorkflowLog struct {
	Meta
	WorkflowId     string         `db:"workflowRef" json:"workflowId"`
	RunId          string         `db:"runRef"      json:"runId"`
	NodeId         string         `db:"nodeId"      json:"nodeId"`
	NodeName       string         `db:"nodeName"    json:"nodeName"`
	TimestampMilli int64          `db:"timestamp"   json:"timestamp"`
	Level          int32          `db:"level"       json:"level"`
	Message        string         `db:"message"     json:"message"`
	Data           map[string]any `db:"data"        json:"data"`
}

type WorkflowLogs []WorkflowLog

func (r WorkflowLogs) ErrorString() string {
	var builder strings.Builder
	for _, log := range r {
		if log.Level >= int32(slog.LevelError) {
			builder.WriteString(log.Message)
			builder.WriteString("\n")
		}
	}
	return strings.TrimSpace(builder.String())
}
