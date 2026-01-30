package tencentcloudcos_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/deployer/providers/tencentcloud-cos"
)

var (
	fInputCertPath string
	fInputKeyPath  string
	fSecretId      string
	fSecretKey     string
	fRegion        string
	fBucket        string
	fDomain        string
)

func init() {
	argsPrefix := "TENCENTCLOUDCOS_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fSecretId, argsPrefix+"SECRETID", "", "")
	flag.StringVar(&fSecretKey, argsPrefix+"SECRETKEY", "", "")
	flag.StringVar(&fRegion, argsPrefix+"REGION", "", "")
	flag.StringVar(&fBucket, argsPrefix+"BUCKET", "", "")
	flag.StringVar(&fDomain, argsPrefix+"DOMAIN", "", "")
}

/*
Shell command to run this test:

	go test -v ./tencentcloud_cos_test.go -args \
	--TENCENTCLOUDCOS_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--TENCENTCLOUDCOS_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--TENCENTCLOUDCOS_SECRETID="your-secret-id" \
	--TENCENTCLOUDCOS_SECRETKEY="your-secret-key" \
	--TENCENTCLOUDCOS_REGION="ap-guangzhou" \
	--TENCENTCLOUDCOS_BUCKET="your-cos-bucket" \
	--TENCENTCLOUDCOS_DOMAIN="example.com"
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
			fmt.Sprintf("BUCKET: %v", fBucket),
			fmt.Sprintf("DOMAIN: %v", fDomain),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			SecretId:  fSecretId,
			SecretKey: fSecretKey,
			Region:    fRegion,
			Bucket:    fBucket,
			Domain:    fDomain,
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
