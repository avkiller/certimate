package goedge_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/deployer/providers/goedge"
)

var (
	fInputCertPath string
	fInputKeyPath  string
	fServerUrl     string
	fAccessKeyId   string
	fAccessKey     string
	fCertificateId int64
)

func init() {
	argsPrefix := "GOEDGE_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fServerUrl, argsPrefix+"SERVERURL", "", "")
	flag.StringVar(&fAccessKeyId, argsPrefix+"ACCESSKEYID", "", "")
	flag.StringVar(&fAccessKey, argsPrefix+"ACCESSKEY", "", "")
	flag.Int64Var(&fCertificateId, argsPrefix+"CERTIFICATEID", 0, "")
}

/*
Shell command to run this test:

	go test -v ./goedge_test.go -args \
	--GOEDGE_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--GOEDGE_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--GOEDGE_SERVERURL="http://127.0.0.1:7788" \
	--GOEDGE_ACCESSKEYID="your-access-key-id" \
	--GOEDGE_ACCESSKEY="your-access-key" \
	--GOEDGE_CERTIFICATEID="your-certificate-id"
*/
func TestDeploy(t *testing.T) {
	flag.Parse()

	t.Run("Deploy_ToCertificate", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("SERVERURL: %v", fServerUrl),
			fmt.Sprintf("ACCESSKEYID: %v", fAccessKeyId),
			fmt.Sprintf("ACCESSKEY: %v", fAccessKey),
			fmt.Sprintf("CERTIFICATEID: %v", fCertificateId),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			ServerUrl:                fServerUrl,
			ApiRole:                  "user",
			AccessKeyId:              fAccessKeyId,
			AccessKey:                fAccessKey,
			AllowInsecureConnections: true,
			ResourceType:             provider.RESOURCE_TYPE_CERTIFICATE,
			CertificateId:            fCertificateId,
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
