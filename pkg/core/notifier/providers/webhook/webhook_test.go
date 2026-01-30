package webhook_test

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/notifier/providers/webhook"
)

const (
	mockSubject = "test_subject"
	mockMessage = "test_message"
)

var (
	fWebhookUrl         string
	fWebhookContentType string
)

func init() {
	argsPrefix := "WEBHOOK_"

	flag.StringVar(&fWebhookUrl, argsPrefix+"URL", "", "")
	flag.StringVar(&fWebhookContentType, argsPrefix+"CONTENTTYPE", "application/json", "")
}

/*
Shell command to run this test:

	go test -v ./webhook_test.go -args \
	--WEBHOOK_URL="https://example.com/your-webhook-url" \
	--WEBHOOK_CONTENTTYPE="application/json"
*/
func TestNotify(t *testing.T) {
	flag.Parse()

	t.Run("Notify", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("URL: %v", fWebhookUrl),
		}, "\n"))

		provider, err := provider.NewNotifier(&provider.NotifierConfig{
			WebhookUrl: fWebhookUrl,
			Method:     "POST",
			Headers: map[string]string{
				"Content-Type": fWebhookContentType,
			},
			AllowInsecureConnections: true,
		})
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		}

		res, err := provider.Notify(context.Background(), mockSubject, mockMessage)
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		}

		t.Logf("ok: %v", res)
	})
}
