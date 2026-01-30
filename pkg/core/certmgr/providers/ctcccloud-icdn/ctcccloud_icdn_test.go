package ctcccloudicdn_test

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/certmgr/providers/ctcccloud-icdn"
)

var (
	fInputCertPath   string
	fInputKeyPath    string
	fAccessKeyId     string
	fSecretAccessKey string
)

func init() {
	argsPrefix := "CTCCCLOUDICDN_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fAccessKeyId, argsPrefix+"ACCESSKEYID", "", "")
	flag.StringVar(&fSecretAccessKey, argsPrefix+"SECRETACCESSKEY", "", "")
}

/*
Shell command to run this test:

	go test -v ./ctcccloud_icdn_test.go -args \
	--CTCCCLOUDICDN_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--CTCCCLOUDICDN_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--CTCCCLOUDICDN_ACCESSKEYID="your-access-key-id" \
	--CTCCCLOUDICDN_SECRETACCESSKEY="your-secret-access-key"
*/
func TestDeploy(t *testing.T) {
	flag.Parse()

	t.Run("Deploy", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("ACCESSKEYID: %v", fAccessKeyId),
			fmt.Sprintf("SECRETACCESSKEY: %v", fSecretAccessKey),
		}, "\n"))

		provider, err := provider.NewCertmgr(&provider.CertmgrConfig{
			AccessKeyId:     fAccessKeyId,
			SecretAccessKey: fSecretAccessKey,
		})
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		}

		fInputCertData, _ := os.ReadFile(fInputCertPath)
		fInputKeyData, _ := os.ReadFile(fInputKeyPath)
		res, err := provider.Upload(context.Background(), string(fInputCertData), string(fInputKeyData))
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		}

		sres, _ := json.Marshal(res)
		t.Logf("ok: %s", string(sres))
	})
}
