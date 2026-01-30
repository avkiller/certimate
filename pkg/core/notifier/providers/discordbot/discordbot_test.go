package discordbot_test

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/notifier/providers/discordbot"
)

const (
	mockSubject = "test_subject"
	mockMessage = "test_message"
)

var (
	fApiToken  string
	fChannelId string
)

func init() {
	argsPrefix := "DISCORDBOT_"

	flag.StringVar(&fApiToken, argsPrefix+"APITOKEN", "", "")
	flag.StringVar(&fChannelId, argsPrefix+"CHANNELID", "", "")
}

/*
Shell command to run this test:

	go test -v ./discordbot_test.go -args \
	--DISCORDBOT_APITOKEN="your-bot-token" \
	--DISCORDBOT_CHANNELID="your-channel-id"
*/
func TestNotify(t *testing.T) {
	flag.Parse()

	t.Run("Notify", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("APITOKEN: %v", fApiToken),
			fmt.Sprintf("CHANNELID: %v", fChannelId),
		}, "\n"))

		provider, err := provider.NewNotifier(&provider.NotifierConfig{
			BotToken:  fApiToken,
			ChannelId: fChannelId,
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
