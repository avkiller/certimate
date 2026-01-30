package mattermost_test

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/notifier/providers/mattermost"
)

const (
	mockSubject = "test_subject"
	mockMessage = "test_message"
)

var (
	fServerUrl string
	fChannelId string
	fUsername  string
	fPassword  string
)

func init() {
	argsPrefix := "MATTERMOST_"

	flag.StringVar(&fServerUrl, argsPrefix+"SERVERURL", "", "")
	flag.StringVar(&fChannelId, argsPrefix+"CHANNELID", "", "")
	flag.StringVar(&fUsername, argsPrefix+"USERNAME", "", "")
	flag.StringVar(&fPassword, argsPrefix+"PASSWORD", "", "")
}

/*
Shell command to run this test:

	go test -v ./mattermost_test.go -args \
	--MATTERMOST_SERVERURL="https://example.com/your-server-url" \
	--MATTERMOST_CHANNELID="your-chanel-id" \
	--MATTERMOST_USERNAME="your-username" \
	--MATTERMOST_PASSWORD="your-password"
*/
func TestNotify(t *testing.T) {
	flag.Parse()

	t.Run("Notify", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("SERVERURL: %v", fServerUrl),
			fmt.Sprintf("CHANNELID: %v", fChannelId),
			fmt.Sprintf("USERNAME: %v", fUsername),
			fmt.Sprintf("PASSWORD: %v", fPassword),
		}, "\n"))

		provider, err := provider.NewNotifier(&provider.NotifierConfig{
			ServerUrl: fServerUrl,
			ChannelId: fChannelId,
			Username:  fUsername,
			Password:  fPassword,
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
