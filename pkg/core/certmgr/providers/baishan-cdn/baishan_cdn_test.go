package baishancdn_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/certmgr/providers/baishan-cdn"
)

var (
	fInputCertPath string
	fInputKeyPath  string
	fApiToken      string
)

func init() {
	argsPrefix := "BAISHANCDN_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fApiToken, argsPrefix+"APITOKEN", "", "")
}

/*
Shell command to run this test:

	go test -v ./baishan_cdn_test.go -args \
	--BAISHANCDN_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--BAISHANCDN_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--BAISHANCDN_APITOKEN="your-api-token"
*/
func TestDeploy(t *testing.T) {
	flag.Parse()

	t.Run("Deploy", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("APITOKEN: %v", fApiToken),
		}, "\n"))

		provider, err := provider.NewCertmgr(&provider.CertmgrConfig{
			ApiToken: fApiToken,
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

		t.Logf("ok: %v", res)
	})
}
