package notifiers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	slackbot "github.com/certimate-go/certimate/pkg/core/notifier/providers/slackbot"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.NotificationProviderTypeSlackBot, func(options *ProviderFactoryOptions) (core.Notifier, error) {
		credentials := domain.AccessConfigForSlackBot{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := slackbot.NewNotifierProvider(&slackbot.NotifierProviderConfig{
			BotToken:  credentials.BotToken,
			ChannelId: xmaps.GetOrDefaultString(options.ProviderExtendedConfig, "channelId", credentials.ChannelId),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
