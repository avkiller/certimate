package tencentcloudscf_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/deployer/providers/tencentcloud-scf"
)

var (
	fInputCertPath string
	fInputKeyPath  string
	fSecretId      string
	fSecretKey     string
	fRegion        string
	fDomain        string
)

func init() {
	argsPrefix := "TENCENTCLOUDSCF_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fSecretId, argsPrefix+"SECRETID", "", "")
	flag.StringVar(&fSecretKey, argsPrefix+"SECRETKEY", "", "")
	flag.StringVar(&fRegion, argsPrefix+"REGION", "", "")
	flag.StringVar(&fDomain, argsPrefix+"DOMAIN", "", "")
}

/*
Shell command to run this test:

	go test -v ./tencentcloud_scf_test.go -args \
	--TENCENTCLOUDSCF_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--TENCENTCLOUDSCF_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--TENCENTCLOUDSCF_SECRETID="your-secret-id" \
	--TENCENTCLOUDSCF_SECRETKEY="your-secret-key" \
	--TENCENTCLOUDSCF_REGION="ap-guangzhou" \
	--TENCENTCLOUDSCF_DOMAIN="example.com"
*/
func TestDeploy(t *testing.T) {
	flag.Parse()

	t.Run("Deploy", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("SECRETID: %v", fSecretId),
			fmt.Sprintf("SECRETKEY: %v", fSecretKey),
			fmt.Sprintf("REGION: %v", fRegion),
			fmt.Sprintf("DOMAIN: %v", fDomain),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			SecretId:           fSecretId,
			SecretKey:          fSecretKey,
			Region:             fRegion,
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
