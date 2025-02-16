package nodeprocessor

import (
	"context"
	"errors"
	"time"

	"github.com/usual2970/certimate/internal/domain"
)

type NodeProcessor interface {
	Process(ctx context.Context) error

	GetLog(ctx context.Context) *domain.WorkflowRunLog
	AppendLogRecord(ctx context.Context, level domain.WorkflowRunLogLevel, content string, err ...string)
}

type nodeLogger struct {
	log *domain.WorkflowRunLog
}

type certificateRepository interface {
	GetByWorkflowNodeId(ctx context.Context, workflowNodeId string) (*domain.Certificate, error)
}

type workflowOutputRepository interface {
	GetByNodeId(ctx context.Context, workflowNodeId string) (*domain.WorkflowOutput, error)
	Save(ctx context.Context, workflowOutput *domain.WorkflowOutput) (*domain.WorkflowOutput, error)
	SaveWithCertificate(ctx context.Context, workflowOutput *domain.WorkflowOutput, certificate *domain.Certificate) (*domain.WorkflowOutput, error)
}

type settingsRepository interface {
	GetByName(ctx context.Context, name string) (*domain.Settings, error)
}

func newNodeLogger(node *domain.WorkflowNode) *nodeLogger {
	return &nodeLogger{
		log: &domain.WorkflowRunLog{
			NodeId:   node.Id,
			NodeName: node.Name,
			Records:  make([]domain.WorkflowRunLogRecord, 0),
		},
	}
}

func (l *nodeLogger) GetLog(ctx context.Context) *domain.WorkflowRunLog {
	return l.log
}

func (l *nodeLogger) AppendLogRecord(ctx context.Context, level domain.WorkflowRunLogLevel, content string, err ...string) {
	record := domain.WorkflowRunLogRecord{
		Time:    time.Now().UTC().Format(time.RFC3339),
		Level:   level,
		Content: content,
	}
	if len(err) > 0 {
		record.Error = err[0]
		l.log.Error = err[0]
	}

	l.log.Records = append(l.log.Records, record)
}

func GetProcessor(node *domain.WorkflowNode) (NodeProcessor, error) {
	switch node.Type {
	case domain.WorkflowNodeTypeStart:
		return NewStartNode(node), nil
	case domain.WorkflowNodeTypeCondition:
		return NewConditionNode(node), nil
	case domain.WorkflowNodeTypeApply:
		return NewApplyNode(node), nil
	case domain.WorkflowNodeTypeUpload:
		return NewUploadNode(node), nil
	case domain.WorkflowNodeTypeDeploy:
		return NewDeployNode(node), nil
	case domain.WorkflowNodeTypeNotify:
		return NewNotifyNode(node), nil
	case domain.WorkflowNodeTypeExecuteSuccess:
		return NewExecuteSuccessNode(node), nil
	case domain.WorkflowNodeTypeExecuteFailure:
		return NewExecuteFailureNode(node), nil
	}
	return nil, errors.New("not implemented")
}

func getContextWorkflowId(ctx context.Context) string {
	return ctx.Value("workflow_id").(string)
}

func getContextWorkflowRunId(ctx context.Context) string {
	return ctx.Value("workflow_run_id").(string)
}
