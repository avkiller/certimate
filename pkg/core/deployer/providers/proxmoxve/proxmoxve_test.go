package proxmoxve_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/deployer/providers/proxmoxve"
)

var (
	fInputCertPath  string
	fInputKeyPath   string
	fServerUrl      string
	fApiToken       string
	fApiTokenSecret string
	fNodeName       string
)

func init() {
	argsPrefix := "PROXMOXVE_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fServerUrl, argsPrefix+"SERVERURL", "", "")
	flag.StringVar(&fApiToken, argsPrefix+"APITOKEN", "", "")
	flag.StringVar(&fApiTokenSecret, argsPrefix+"APITOKENSECRET", "", "")
	flag.StringVar(&fNodeName, argsPrefix+"NODENAME", "", "")
}

/*
Shell command to run this test:

	go test -v ./proxmoxve_test.go -args \
	--PROXMOXVE_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--PROXMOXVE_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--PROXMOXVE_SERVERURL="http://127.0.0.1:8006" \
	--PROXMOXVE_APITOKEN="your-api-token" \
	--PROXMOXVE_APITOKENSECRET="your-api-token-secret" \
	--PROXMOXVE_NODENAME="your-cluster-node-name"
*/
func TestDeploy(t *testing.T) {
	flag.Parse()

	t.Run("Deploy", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("SERVERURL: %v", fServerUrl),
			fmt.Sprintf("APITOKEN: %v", fApiToken),
			fmt.Sprintf("APITOKENSECRET: %v", fApiTokenSecret),
			fmt.Sprintf("NODENAME: %v", fNodeName),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			ServerUrl:                fServerUrl,
			ApiToken:                 fApiToken,
			ApiTokenSecret:           fApiTokenSecret,
			AllowInsecureConnections: true,
			NodeName:                 fNodeName,
			AutoRestart:              true,
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
