package huaweicloudwaf_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/deployer/providers/huaweicloud-waf"
)

var (
	fInputCertPath   string
	fInputKeyPath    string
	fAccessKeyId     string
	fSecretAccessKey string
	fRegion          string
	fResourceType    string
	fDomain          string
)

func init() {
	argsPrefix := "HUAWEICLOUDWAF_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fAccessKeyId, argsPrefix+"ACCESSKEYID", "", "")
	flag.StringVar(&fSecretAccessKey, argsPrefix+"SECRETACCESSKEY", "", "")
	flag.StringVar(&fRegion, argsPrefix+"REGION", "", "")
	flag.StringVar(&fResourceType, argsPrefix+"RESOURCETYPE", "", "")
	flag.StringVar(&fDomain, argsPrefix+"DOMAIN", "", "")
}

/*
Shell command to run this test:

	go test -v ./huaweicloud_waf_test.go -args \
	--HUAWEICLOUDWAF_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--HUAWEICLOUDWAF_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--HUAWEICLOUDWAF_ACCESSKEYID="your-access-key-id" \
	--HUAWEICLOUDWAF_SECRETACCESSKEY="your-secret-access-key" \
	--HUAWEICLOUDWAF_REGION="cn-north-1" \
	--HUAWEICLOUDWAF_RESOURCETYPE="premium" \
	--HUAWEICLOUDWAF_DOMAIN="example.com"
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
			fmt.Sprintf("REGION: %v", fRegion),
			fmt.Sprintf("RESOURCETYPE: %v", fResourceType),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			AccessKeyId:     fAccessKeyId,
			SecretAccessKey: fSecretAccessKey,
			Region:          fRegion,
			ResourceType:    fResourceType,
			Domain:          fDomain,
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
