package discordbot_test

import (
	"testing"

	impl "github.com/certimate-go/certimate/pkg/core/notifier/providers/discordbot"
	tester "github.com/certimate-go/certimate/pkg/core/notifier/testing"
)

var (
	fp         = tester.Args("DISCORDBOT_")
	fApiToken  string
	fChannelId string
)

func init() {
	fp.DefineString(&fApiToken, "APITOKEN")
	fp.DefineString(&fChannelId, "CHANNELID")
}

/*
Shell command to run this test:

	go test -v ./discordbot_test.go -args \
	--DISCORDBOT_APITOKEN="your-bot-token" \
	--DISCORDBOT_CHANNELID="your-channel-id"
*/
func TestProvider(t *testing.T) {
	fp.Parse()

	t.Run("Notify", func(t *testing.T) {
		provider, err := impl.NewNotifier(&impl.NotifierConfig{
			BotToken:  fApiToken,
			ChannelId: fChannelId,
		})
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		}

		tester.TestNotify(t, provider, tester.TestNotifyArgs{})
	})
}
