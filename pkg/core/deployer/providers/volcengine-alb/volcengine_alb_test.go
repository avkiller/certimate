package volcenginealb_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/deployer/providers/volcengine-alb"
)

var (
	fInputCertPath   string
	fInputKeyPath    string
	fAccessKeyId     string
	fAccessKeySecret string
	fRegion          string
	fListenerId      string
)

func init() {
	argsPrefix := "VOLCENGINEALB_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fAccessKeyId, argsPrefix+"ACCESSKEYID", "", "")
	flag.StringVar(&fAccessKeySecret, argsPrefix+"ACCESSKEYSECRET", "", "")
	flag.StringVar(&fRegion, argsPrefix+"REGION", "", "")
	flag.StringVar(&fListenerId, argsPrefix+"LISTENERID", "", "")
}

/*
Shell command to run this test:

	go test -v ./volcengine_alb_test.go -args \
	--VOLCENGINEALB_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--VOLCENGINEALB_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--VOLCENGINEALB_ACCESSKEYID="your-access-key-id" \
	--VOLCENGINEALB_ACCESSKEYSECRET="your-access-key-secret" \
	--VOLCENGINEALB_REGION="cn-beijing" \
	--VOLCENGINEALB_LISTENERID="your-listener-id"
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
			fmt.Sprintf("LISTENERID: %v", fListenerId),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			AccessKeyId:     fAccessKeyId,
			AccessKeySecret: fAccessKeySecret,
			Region:          fRegion,
			ResourceType:    provider.RESOURCE_TYPE_LISTENER,
			ListenerId:      fListenerId,
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
