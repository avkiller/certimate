package aliyunapigw_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/deployer/providers/aliyun-apigw"
)

var (
	fInputCertPath   string
	fInputKeyPath    string
	fAccessKeyId     string
	fAccessKeySecret string
	fRegion          string
	fServiceType     string
	fGatewayId       string
	fGroupId         string
	fDomain          string
)

func init() {
	argsPrefix := "ALIYUNAPIGW_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fAccessKeyId, argsPrefix+"ACCESSKEYID", "", "")
	flag.StringVar(&fAccessKeySecret, argsPrefix+"ACCESSKEYSECRET", "", "")
	flag.StringVar(&fRegion, argsPrefix+"REGION", "", "")
	flag.StringVar(&fGatewayId, argsPrefix+"GATEWARYID", "", "")
	flag.StringVar(&fGroupId, argsPrefix+"GROUPID", "", "")
	flag.StringVar(&fServiceType, argsPrefix+"SERVICETYPE", "", "")
	flag.StringVar(&fDomain, argsPrefix+"DOMAIN", "", "")
}

/*
Shell command to run this test:

	go test -v ./aliyun_apigw_test.go -args \
	--ALIYUNAPIGW_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--ALIYUNAPIGW_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--ALIYUNAPIGW_ACCESSKEYID="your-access-key-id" \
	--ALIYUNAPIGW_ACCESSKEYSECRET="your-access-key-secret" \
	--ALIYUNAPIGW_REGION="cn-hangzhou" \
	--ALIYUNAPIGW_SERVICETYPE="cloudnative" \
	--ALIYUNAPIGW_GATEWAYID="your-api-gateway-id" \
	--ALIYUNAPIGW_GROUPID="your-api-group-id" \
	--ALIYUNAPIGW_DOMAIN="example.com"
*/
func TestDeploy(t *testing.T) {
	flag.Parse()

	t.Run("Deploy", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("ACCESSKEYID: %v", fAccessKeyId),
			fmt.Sprintf("ACCESSKEYSECRET: %v", fAccessKeySecret),
			fmt.Sprintf("REGION: %v", fRegion),
			fmt.Sprintf("SERVICETYPE: %v", fServiceType),
			fmt.Sprintf("GATEWAYID: %v", fGatewayId),
			fmt.Sprintf("GROUPID: %v", fGroupId),
			fmt.Sprintf("DOMAIN: %v", fDomain),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			AccessKeyId:        fAccessKeyId,
			AccessKeySecret:    fAccessKeySecret,
			Region:             fRegion,
			ServiceType:        fServiceType,
			GatewayId:          fGatewayId,
			GroupId:            fGroupId,
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
