package tencentcloudclb_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/deployer/providers/tencentcloud-clb"
)

var (
	fInputCertPath  string
	fInputKeyPath   string
	fSecretId       string
	fSecretKey      string
	fRegion         string
	fLoadbalancerId string
	fListenerId     string
	fDomain         string
)

func init() {
	argsPrefix := "TENCENTCLOUDCDN_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fSecretId, argsPrefix+"SECRETID", "", "")
	flag.StringVar(&fSecretKey, argsPrefix+"SECRETKEY", "", "")
	flag.StringVar(&fRegion, argsPrefix+"REGION", "", "")
	flag.StringVar(&fLoadbalancerId, argsPrefix+"LOADBALANCERID", "", "")
	flag.StringVar(&fListenerId, argsPrefix+"LISTENERID", "", "")
	flag.StringVar(&fDomain, argsPrefix+"DOMAIN", "", "")
}

/*
Shell command to run this test:

	go test -v ./tencentcloud_clb_test.go -args \
	--TENCENTCLOUDCLB_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--TENCENTCLOUDCLB_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--TENCENTCLOUDCLB_SECRETID="your-secret-id" \
	--TENCENTCLOUDCLB_SECRETKEY="your-secret-key" \
	--TENCENTCLOUDCLB_REGION="ap-guangzhou" \
	--TENCENTCLOUDCLB_LOADBALANCERID="your-clb-lb-id" \
	--TENCENTCLOUDCLB_LISTENERID="your-clb-lbl-id" \
	--TENCENTCLOUDCLB_DOMAIN="example.com"
*/
func TestDeploy(t *testing.T) {
	flag.Parse()

	t.Run("Deploy_ToLoadbalancer", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("SECRETID: %v", fSecretId),
			fmt.Sprintf("SECRETKEY: %v", fSecretKey),
			fmt.Sprintf("REGION: %v", fRegion),
			fmt.Sprintf("LOADBALANCERID: %v", fLoadbalancerId),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			SecretId:       fSecretId,
			SecretKey:      fSecretKey,
			Region:         fRegion,
			ResourceType:   provider.RESOURCE_TYPE_LOADBALANCER,
			LoadbalancerId: fLoadbalancerId,
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
			fmt.Sprintf("SECRETID: %v", fSecretId),
			fmt.Sprintf("SECRETKEY: %v", fSecretKey),
			fmt.Sprintf("REGION: %v", fRegion),
			fmt.Sprintf("LOADBALANCERID: %v", fLoadbalancerId),
			fmt.Sprintf("LISTENERID: %v", fListenerId),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			SecretId:       fSecretId,
			SecretKey:      fSecretKey,
			Region:         fRegion,
			ResourceType:   provider.RESOURCE_TYPE_LISTENER,
			LoadbalancerId: fLoadbalancerId,
			ListenerId:     fListenerId,
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

	t.Run("Deploy_ToRuleDomain", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("SECRETID: %v", fSecretId),
			fmt.Sprintf("SECRETKEY: %v", fSecretKey),
			fmt.Sprintf("REGION: %v", fRegion),
			fmt.Sprintf("LOADBALANCERID: %v", fLoadbalancerId),
			fmt.Sprintf("LISTENERID: %v", fListenerId),
			fmt.Sprintf("DOMAIN: %v", fDomain),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			SecretId:       fSecretId,
			SecretKey:      fSecretKey,
			Region:         fRegion,
			ResourceType:   provider.RESOURCE_TYPE_RULEDOMAIN,
			LoadbalancerId: fLoadbalancerId,
			ListenerId:     fListenerId,
			Domain:         fDomain,
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
