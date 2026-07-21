package wecombot_test

import (
	"testing"

	impl "github.com/certimate-go/certimate/pkg/core/notifier/providers/wecombot"
	tester "github.com/certimate-go/certimate/pkg/core/notifier/testing"
)

var (
	fp          = tester.Args("WECOMBOT_")
	fWebhookUrl string
)

func init() {
	fp.DefineString(&fWebhookUrl, "WEBHOOKURL")
}

/*
Shell command to run this test:

	go test -v ./wecombot_test.go -args \
	--WECOMBOT_WEBHOOKURL="https://example.com/your-webhook-url" \
*/
func TestProvider(t *testing.T) {
	fp.Parse()

	t.Run("Notify", func(t *testing.T) {
		provider, err := impl.NewNotifier(&impl.NotifierConfig{
			WebhookUrl: fWebhookUrl,
		})
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		}

		tester.TestNotify(t, provider, tester.TestNotifyArgs{})
	})
}
