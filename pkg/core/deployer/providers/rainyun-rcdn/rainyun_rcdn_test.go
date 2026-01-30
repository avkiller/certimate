package rainyunrcdn_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/deployer/providers/rainyun-rcdn"
)

var (
	fInputCertPath string
	fInputKeyPath  string
	fApiKey        string
	fInstanceId    int64
	fDomain        string
)

func init() {
	argsPrefix := "RAINYUNRCDN_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fApiKey, argsPrefix+"APIKEY", "", "")
	flag.Int64Var(&fInstanceId, argsPrefix+"INSTANCEID", 0, "")
	flag.StringVar(&fDomain, argsPrefix+"DOMAIN", "", "")
}

/*
Shell command to run this test:

	go test -v ./rainyun_rcdn_test.go -args \
	--RAINYUNRCDN_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--RAINYUNRCDN_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--RAINYUNRCDN_APIKEY="your-api-key" \
	--RAINYUNRCDN_INSTANCEID="your-rcdn-instance-id" \
	--RAINYUNRCDN_DOMAIN="example.com"
*/
func TestDeploy(t *testing.T) {
	flag.Parse()

	t.Run("Deploy", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("APIKEY: %v", fApiKey),
			fmt.Sprintf("INSTANCEID: %v", fInstanceId),
			fmt.Sprintf("DOMAIN: %v", fDomain),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			ApiKey:             fApiKey,
			InstanceId:         fInstanceId,
			DomainMatchPattern: provider.DOMAIN_MATCH_PATTERN_EXACT,
			Domain:             fDomain,
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
