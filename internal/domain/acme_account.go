package domain

import (
	"crypto"

	"github.com/go-acme/lego/v4/acme"
	"github.com/go-acme/lego/v4/registration"

	xcert "github.com/certimate-go/certimate/pkg/utils/cert"
)

const CollectionNameACMEAccount = "acme_accounts"

type ACMEAccount struct {
	Meta
	CA          string        `json:"ca" db:"ca"`
	Email       string        `json:"email" db:"email"`
	PrivateKey  string        `json:"privateKey" db:"privateKey"`
	ACMEAccount *acme.Account `json:"acmeAccount" db:"acmeAccount"`
	ACMEAcctUrl string        `json:"acmeAcctUrl" db:"acmeAcctUrl"`
	ACMEDirUrl  string        `json:"acmeDirUrl" db:"acmeDirUrl"`
}

func (a *ACMEAccount) GetEmail() string {
	return a.Email
}

func (a *ACMEAccount) GetRegistration() *registration.Resource {
	if a.ACMEAccount == nil {
		return nil
	}

	return &registration.Resource{
		Body: *a.ACMEAccount,
		URI:  a.ACMEAcctUrl,
	}
}

func (a *ACMEAccount) GetPrivateKey() crypto.PrivateKey {
	if a.PrivateKey == "" {
		return nil
	}

	rs, _ := xcert.ParseECPrivateKeyFromPEM(a.PrivateKey)
	return rs
}
