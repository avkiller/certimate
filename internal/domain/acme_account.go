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
	CA          string        `db:"ca"          json:"ca"`
	Email       string        `db:"email"       json:"email"`
	PrivateKey  string        `db:"privateKey"  json:"privateKey"`
	ACMEAccount *acme.Account `db:"acmeAccount" json:"acmeAccount"`
	ACMEAcctUrl string        `db:"acmeAcctUrl" json:"acmeAcctUrl"`
	ACMEDirUrl  string        `db:"acmeDirUrl"  json:"acmeDirUrl"`
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

	rs, _ := xcert.ParsePrivateKeyFromPEM(a.PrivateKey)
	return rs
}
