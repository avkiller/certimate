package certapply

import (
	"errors"

	"github.com/go-acme/lego/v4/lego"
)

type ACMEClient struct {
	client  *lego.Client
	account *ACMEAccount
}

func NewACMEClient(config *ACMEConfig, email string, configures ...func(*lego.Config) error) (*ACMEClient, error) {
	account, err := NewACMEAccountWithSingleFlight(config, email)
	if err != nil {
		return nil, err
	}

	tconfigures := []func(*lego.Config) error{
		func(legoCfg *lego.Config) error {
			legoCfg.CADirURL = config.CADirUrl
			legoCfg.Certificate.KeyType = config.CertifierKeyType
			return nil
		},
	}
	tconfigures = append(tconfigures, configures...)
	return newACMEClientWithAccount(account, tconfigures...)
}

func NewACMEClientWithAccount(account *ACMEAccount, configures ...func(*lego.Config) error) (*ACMEClient, error) {
	return newACMEClientWithAccount(account, configures...)
}

func newACMEClientWithAccount(account *ACMEAccount, configures ...func(*lego.Config) error) (*ACMEClient, error) {
	if account == nil {
		return nil, errors.New("the acme account is nil")
	}

	legoCfg := lego.NewConfig(account)
	legoCfg.CADirURL = account.ACMEDirUrl

	errs := make([]error, 0)
	for _, configure := range configures {
		if err := configure(legoCfg); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	legoClient, err := lego.NewClient(legoCfg)
	if err != nil {
		return nil, err
	}

	return &ACMEClient{
		client:  legoClient,
		account: account,
	}, nil
}
