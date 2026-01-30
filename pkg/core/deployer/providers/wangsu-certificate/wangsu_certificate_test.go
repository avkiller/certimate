package wangsucertificate_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/deployer/providers/wangsu-certificate"
)

var (
	fInputCertPath   string
	fInputKeyPath    string
	fAccessKeyId     string
	fAccessKeySecret string
	fCertificateId   string
)

func init() {
	argsPrefix := "WANGSUCERTIFICATE_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fAccessKeyId, argsPrefix+"ACCESSKEYID", "", "")
	flag.StringVar(&fAccessKeySecret, argsPrefix+"ACCESSKEYSECRET", "", "")
	flag.StringVar(&fCertificateId, argsPrefix+"CERTIFICATEID", "", "")
}

/*
Shell command to run this test:

	go test -v ./wangsu_certificate_test.go -args \
	--WANGSUCERTIFICATE_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--WANGSUCERTIFICATE_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--WANGSUCERTIFICATE_ACCESSKEYID="your-access-key-id" \
	--WANGSUCERTIFICATE_ACCESSKEYSECRET="your-access-key-secret" \
	--WANGSUCERTIFICATE_CERTIFICATEID="your-certificate-id"
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
			fmt.Sprintf("CERTIFICATEID: %v", fCertificateId),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			AccessKeyId:     fAccessKeyId,
			AccessKeySecret: fAccessKeySecret,
			CertificateId:   fCertificateId,
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
