package notify

import (
	"context"
	"fmt"

	"github.com/certimate-go/certimate/internal/domain/dtos"
)

const (
	notifyTestSubject = "[Certimate] Notification Testing"
	notifyTestMessage = "Welcome to use Certimate!"
)

type NotifyService struct{}

func NewNotifyService() *NotifyService {
	return &NotifyService{}
}

func (n *NotifyService) TestPush(ctx context.Context, req *dtos.NotifyTestPushReq) error {
	// TODO: 测试通知
	return fmt.Errorf("not implemented")
}
