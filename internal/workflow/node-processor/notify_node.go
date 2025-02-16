package nodeprocessor

import (
	"context"

	"github.com/usual2970/certimate/internal/domain"
	"github.com/usual2970/certimate/internal/notify"
	"github.com/usual2970/certimate/internal/repository"
)

type notifyNode struct {
	node *domain.WorkflowNode
	*nodeLogger

	settingsRepo settingsRepository
}

func NewNotifyNode(node *domain.WorkflowNode) *notifyNode {
	return &notifyNode{
		node:       node,
		nodeLogger: newNodeLogger(node),

		settingsRepo: repository.NewSettingsRepository(),
	}
}

func (n *notifyNode) Process(ctx context.Context) error {
	n.AppendLogRecord(ctx, domain.WorkflowRunLogLevelInfo, "进入推送通知节点")

	nodeConfig := n.node.GetConfigForNotify()

	// 获取通知配置
	settings, err := n.settingsRepo.GetByName(ctx, "notifyChannels")
	if err != nil {
		n.AppendLogRecord(ctx, domain.WorkflowRunLogLevelError, "获取通知配置失败", err.Error())
		return err
	}

	// 获取通知渠道
	channelConfig, err := settings.GetNotifyChannelConfig(nodeConfig.Channel)
	if err != nil {
		n.AppendLogRecord(ctx, domain.WorkflowRunLogLevelError, "获取通知渠道配置失败", err.Error())
		return err
	}

	// 发送通知
	if err := notify.SendToChannel(nodeConfig.Subject, nodeConfig.Message, nodeConfig.Channel, channelConfig); err != nil {
		n.AppendLogRecord(ctx, domain.WorkflowRunLogLevelError, "发送通知失败", err.Error())
		return err
	}
	n.AppendLogRecord(ctx, domain.WorkflowRunLogLevelInfo, "发送通知成功")

	return nil
}
