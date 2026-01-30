package notifiers

import (
	"fmt"

	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/pkg/core"
	"github.com/certimate-go/certimate/pkg/core/notifier/providers/email"
	xmaps "github.com/certimate-go/certimate/pkg/utils/maps"
)

func init() {
	if err := Registries.Register(domain.NotificationProviderTypeEmail, func(options *ProviderFactoryOptions) (core.Notifier, error) {
		credentials := domain.AccessConfigForEmail{}
		if err := xmaps.Populate(options.ProviderAccessConfig, &credentials); err != nil {
			return nil, fmt.Errorf("failed to populate provider access config: %w", err)
		}

		provider, err := email.NewNotifierProvider(&email.NotifierProviderConfig{
			SmtpHost:        credentials.SmtpHost,
			SmtpPort:        credentials.SmtpPort,
			SmtpTls:         credentials.SmtpTls,
			Username:        credentials.Username,
			Password:        credentials.Password,
			SenderAddress:   credentials.SenderAddress,
			SenderName:      credentials.SenderName,
			ReceiverAddress: xmaps.GetOrDefaultString(options.ProviderExtendedConfig, "receiverAddress", credentials.ReceiverAddress),
		})
		return provider, err
	}); err != nil {
		panic(err)
	}
}
