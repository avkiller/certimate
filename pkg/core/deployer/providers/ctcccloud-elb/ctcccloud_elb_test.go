package ctcccloudelb_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/deployer/providers/ctcccloud-elb"
)

var (
	fInputCertPath   string
	fInputKeyPath    string
	fAccessKeyId     string
	fSecretAccessKey string
	fRegionId        string
	fLoadbalancerId  string
	fListenerId      string
)

func init() {
	argsPrefix := "CTCCCLOUDELB_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fAccessKeyId, argsPrefix+"ACCESSKEYID", "", "")
	flag.StringVar(&fSecretAccessKey, argsPrefix+"SECRETACCESSKEY", "", "")
	flag.StringVar(&fRegionId, argsPrefix+"REGIONID", "", "")
	flag.StringVar(&fLoadbalancerId, argsPrefix+"LOADBALANCERID", "", "")
	flag.StringVar(&fListenerId, argsPrefix+"LISTENERID", "", "")
}

/*
Shell command to run this test:

	go test -v ./ctcccloud_elb_test.go -args \
	--CTCCCLOUDELB_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--CTCCCLOUDELB_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--CTCCCLOUDELB_ACCESSKEYID="your-access-key-id" \
	--CTCCCLOUDELB_SECRETACCESSKEY="your-secret-access-key" \
	--CTCCCLOUDELB_REGIONID="your-region-id" \
	--CTCCCLOUDELB_LOADBALANCERID="your-elb-instance-id" \
	--CTCCCLOUDELB_LISTENERID="your-elb-listener-id"
*/
func TestDeploy(t *testing.T) {
	flag.Parse()

	t.Run("Deploy_ToLoadbalancer", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("ACCESSKEYID: %v", fAccessKeyId),
			fmt.Sprintf("SECRETACCESSKEY: %v", fSecretAccessKey),
			fmt.Sprintf("REGIONID: %v", fRegionId),
			fmt.Sprintf("LOADBALANCERID: %v", fLoadbalancerId),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			AccessKeyId:     fAccessKeyId,
			SecretAccessKey: fSecretAccessKey,
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
			fmt.Sprintf("SECRETACCESSKEY: %v", fSecretAccessKey),
			fmt.Sprintf("REGIONID: %v", fRegionId),
			fmt.Sprintf("LISTENERID: %v", fListenerId),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			AccessKeyId:     fAccessKeyId,
			SecretAccessKey: fSecretAccessKey,
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
