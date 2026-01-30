package upyuncdn_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/deployer/providers/upyun-cdn"
)

var (
	fInputCertPath string
	fInputKeyPath  string
	fUsername      string
	fPassword      string
	fDomain        string
)

func init() {
	argsPrefix := "UPYUNCDN_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fUsername, argsPrefix+"USERNAME", "", "")
	flag.StringVar(&fPassword, argsPrefix+"PASSWORD", "", "")
	flag.StringVar(&fDomain, argsPrefix+"DOMAIN", "", "")
}

/*
Shell command to run this test:

	go test -v ./upyun_cdn_test.go -args \
	--UPYUNCDN_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--UPYUNCDN_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--UPYUNCDN_USERNAME="your-username" \
	--UPYUNCDN_PASSWORD="your-password" \
	--UPYUNCDN_DOMAIN="example.com"
*/
func TestDeploy(t *testing.T) {
	flag.Parse()

	t.Run("Deploy", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("USERNAME: %v", fUsername),
			fmt.Sprintf("PASSWORD: %v", fPassword),
			fmt.Sprintf("DOMAIN: %v", fDomain),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			Username: fUsername,
			Password: fPassword,
			Domain:   fDomain,
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
