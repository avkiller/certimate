package cachefly_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/deployer/providers/cachefly"
)

var (
	fInputCertPath string
	fInputKeyPath  string
	fApiToken      string
)

func init() {
	argsPrefix := "CACHEFLY_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fApiToken, argsPrefix+"APITOKEN", "", "")
}

/*
Shell command to run this test:

	go test -v ./cachefly_test.go -args \
	--CACHEFLY_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--CACHEFLY_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--CACHEFLY_APITOKEN="your-api-token"
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

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			ApiToken: fApiToken,
		})
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		}

		fInputCertData, _ := os.ReadFile(fInputCertPath)
		fInputKeyData, _ := os.ReadFile(fInputKeyPath)
		res, err := provider.Deploy(context.Background(), string(fInputCertData), string(fInputKeyData))
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		}

		t.Logf("ok: %v", res)
	})
}
