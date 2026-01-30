package baotapanelconsole_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/deployer/providers/baotapanel-console"
)

var (
	fInputCertPath string
	fInputKeyPath  string
	fServerUrl     string
	fApiKey        string
)

func init() {
	argsPrefix := "BAOTAPANELCONSOLE_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fServerUrl, argsPrefix+"SERVERURL", "", "")
	flag.StringVar(&fApiKey, argsPrefix+"APIKEY", "", "")
}

/*
Shell command to run this test:

	go test -v ./baotapanel_console_test.go -args \
	--BAOTAPANELCONSOLE_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--BAOTAPANELCONSOLE_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--BAOTAPANELCONSOLE_SERVERURL="http://127.0.0.1:8888" \
	--BAOTAPANELCONSOLE_APIKEY="your-api-key"
*/
func TestDeploy(t *testing.T) {
	flag.Parse()

	t.Run("Deploy", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("SERVERURL: %v", fServerUrl),
			fmt.Sprintf("APIKEY: %v", fApiKey),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			ServerUrl:                fServerUrl,
			ApiKey:                   fApiKey,
			AllowInsecureConnections: true,
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
