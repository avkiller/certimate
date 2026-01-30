package tencentcloudeo_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/deployer/providers/tencentcloud-eo"
)

var (
	fInputCertPath string
	fInputKeyPath  string
	fSecretId      string
	fSecretKey     string
	fZoneId        string
	fDomains       string
)

func init() {
	argsPrefix := "TENCENTCLOUDEO_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fSecretId, argsPrefix+"SECRETID", "", "")
	flag.StringVar(&fSecretKey, argsPrefix+"SECRETKEY", "", "")
	flag.StringVar(&fZoneId, argsPrefix+"ZONEID", "", "")
	flag.StringVar(&fDomains, argsPrefix+"DOMAINS", "", "")
}

/*
Shell command to run this test:

	go test -v ./tencentcloud_eo_test.go -args \
	--TENCENTCLOUDEO_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--TENCENTCLOUDEO_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--TENCENTCLOUDEO_SECRETID="your-secret-id" \
	--TENCENTCLOUDEO_SECRETKEY="your-secret-key" \
	--TENCENTCLOUDEO_ZONEID="your-zone-id" \
	--TENCENTCLOUDEO_DOMAINS="example.com"
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
			fmt.Sprintf("ZONEID: %v", fZoneId),
			fmt.Sprintf("DOMAINS: %v", fDomains),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			SecretId:           fSecretId,
			SecretKey:          fSecretKey,
			ZoneId:             fZoneId,
			DomainMatchPattern: provider.DOMAIN_MATCH_PATTERN_EXACT,
			Domains:            strings.Split(fDomains, ";"),
			EnableMultipleSSL:  true,
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
