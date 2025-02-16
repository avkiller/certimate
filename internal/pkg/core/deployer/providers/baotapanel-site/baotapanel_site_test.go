﻿package baotapanelsite_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/usual2970/certimate/internal/pkg/core/deployer/providers/baotapanel-site"
)

var (
	fInputCertPath string
	fInputKeyPath  string
	fApiUrl        string
	fApiKey        string
	fSiteName      string
)

func init() {
	argsPrefix := "CERTIMATE_DEPLOYER_BAOTAPANELSITE_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fApiUrl, argsPrefix+"APIURL", "", "")
	flag.StringVar(&fApiKey, argsPrefix+"APIKEY", "", "")
	flag.StringVar(&fSiteName, argsPrefix+"SITENAME", "", "")
}

/*
Shell command to run this test:

	go test -v ./baotapanel_site_test.go -args \
	--CERTIMATE_DEPLOYER_BAOTAPANELSITE_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--CERTIMATE_DEPLOYER_BAOTAPANELSITE_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--CERTIMATE_DEPLOYER_BAOTAPANELSITE_APIURL="your-baota-panel-url" \
	--CERTIMATE_DEPLOYER_BAOTAPANELSITE_APIKEY="your-baota-panel-key" \
	--CERTIMATE_DEPLOYER_BAOTAPANELSITE_SITENAME="your-baota-site-name"
*/
func TestDeploy(t *testing.T) {
	flag.Parse()

	t.Run("Deploy", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("APIURL: %v", fApiUrl),
			fmt.Sprintf("APIKEY: %v", fApiKey),
			fmt.Sprintf("SITENAME: %v", fSiteName),
		}, "\n"))

		deployer, err := provider.New(&provider.BaotaPanelSiteDeployerConfig{
			ApiUrl:   fApiUrl,
			ApiKey:   fApiKey,
			SiteName: fSiteName,
		})
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		}

		fInputCertData, _ := os.ReadFile(fInputCertPath)
		fInputKeyData, _ := os.ReadFile(fInputKeyPath)
		res, err := deployer.Deploy(context.Background(), string(fInputCertData), string(fInputKeyData))
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		}

		t.Logf("ok: %v", res)
	})
}
