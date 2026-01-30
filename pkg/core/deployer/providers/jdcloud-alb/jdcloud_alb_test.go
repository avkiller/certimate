package jdcloudalb_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/deployer/providers/jdcloud-alb"
)

var (
	fInputCertPath   string
	fInputKeyPath    string
	fAccessKeyId     string
	fAccessKeySecret string
	fRegionId        string
	fLoadbalancerId  string
	fListenerId      string
)

func init() {
	argsPrefix := "JDCLOUDALB_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fAccessKeyId, argsPrefix+"ACCESSKEYID", "", "")
	flag.StringVar(&fAccessKeySecret, argsPrefix+"ACCESSKEYSECRET", "", "")
	flag.StringVar(&fRegionId, argsPrefix+"REGIONID", "", "")
	flag.StringVar(&fLoadbalancerId, argsPrefix+"LOADBALANCERID", "", "")
	flag.StringVar(&fListenerId, argsPrefix+"LISTENERID", "", "")
}

/*
Shell command to run this test:

	go test -v ./jdcloud_alb_test.go -args \
	--JDCLOUDALB_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--JDCLOUDALB_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--JDCLOUDALB_ACCESSKEYID="your-access-key-id" \
	--JDCLOUDALB_ACCESSKEYSECRET="your-secret-access-key" \
	--JDCLOUDALB_REGION_ID="cn-north-1" \
	--JDCLOUDALB_LOADBALANCERID="your-alb-loadbalancer-id" \
	--JDCLOUDALB_LISTENERID="your-alb-listener-id"
*/
func TestDeploy(t *testing.T) {
	flag.Parse()

	t.Run("Deploy_ToLoadbalancer", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("ACCESSKEYID: %v", fAccessKeyId),
			fmt.Sprintf("ACCESSKEYSECRET: %v", fAccessKeySecret),
			fmt.Sprintf("REGIONID: %v", fRegionId),
			fmt.Sprintf("LOADBALANCERID: %v", fLoadbalancerId),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			AccessKeyId:     fAccessKeyId,
			AccessKeySecret: fAccessKeySecret,
			RegionId:        fRegionId,
			ResourceType:    provider.RESOURCE_TYPE_LOADBALANCER,
			LoadbalancerId:  fLoadbalancerId,
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

	t.Run("Deploy_ToListener", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("ACCESSKEYID: %v", fAccessKeyId),
			fmt.Sprintf("ACCESSKEYSECRET: %v", fAccessKeySecret),
			fmt.Sprintf("REGIONID: %v", fRegionId),
			fmt.Sprintf("LISTENERID: %v", fListenerId),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			AccessKeyId:     fAccessKeyId,
			AccessKeySecret: fAccessKeySecret,
			RegionId:        fRegionId,
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
