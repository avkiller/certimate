package upyunssl_test

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/certmgr/providers/upyun-ssl"
)

var (
	fInputCertPath string
	fInputKeyPath  string
	fUsername      string
	fPassword      string
)

func init() {
	argsPrefix := "UPYUNSSL_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fUsername, argsPrefix+"USERNAME", "", "")
	flag.StringVar(&fPassword, argsPrefix+"PASSWORD", "", "")
}

/*
Shell command to run this test:

	go test -v ./upyun_ssl_test.go -args \
	--UPYUNSSL_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--UPYUNSSL_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--UPYUNSSL_USERNAME="your-username" \
	--UPYUNSSL_PASSWORD="your-password"
*/
func TestDeploy(t *testing.T) {
	flag.Parse()

	t.Run("Deploy", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("USERNAME: %v", fUsername),
			fmt.Sprintf("PASSWORD: %v", fPassword),
		}, "\n"))

		provider, err := provider.NewCertmgr(&provider.CertmgrConfig{
			Username: fUsername,
			Password: fPassword,
		})
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		}

		fInputCertData, _ := os.ReadFile(fInputCertPath)
		fInputKeyData, _ := os.ReadFile(fInputKeyPath)
		res, err := provider.Upload(context.Background(), string(fInputCertData), string(fInputKeyData))
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		}

		sres, _ := json.Marshal(res)
		t.Logf("ok: %s", string(sres))
	})
}
