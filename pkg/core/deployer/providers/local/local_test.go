package local_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/certimate-go/certimate/pkg/core/deployer/providers/local"
)

var (
	fInputCertPath  string
	fInputKeyPath   string
	fOutputCertPath string
	fOutputKeyPath  string
	fPfxPassword    string
	fJksAlias       string
	fJksKeypass     string
	fJksStorepass   string
	fShellEnv       string
	fPreCommand     string
	fPostCommand    string
)

func init() {
	argsPrefix := "LOCAL_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fOutputCertPath, argsPrefix+"OUTPUTCERTPATH", "", "")
	flag.StringVar(&fOutputKeyPath, argsPrefix+"OUTPUTKEYPATH", "", "")
	flag.StringVar(&fPfxPassword, argsPrefix+"PFXPASSWORD", "", "")
	flag.StringVar(&fJksAlias, argsPrefix+"JKSALIAS", "", "")
	flag.StringVar(&fJksKeypass, argsPrefix+"JKSKEYPASS", "", "")
	flag.StringVar(&fJksStorepass, argsPrefix+"JKSSTOREPASS", "", "")
	flag.StringVar(&fShellEnv, argsPrefix+"SHELLENV", "", "")
	flag.StringVar(&fPreCommand, argsPrefix+"PRECOMMAND", "", "")
	flag.StringVar(&fPostCommand, argsPrefix+"POSTCOMMAND", "", "")
}

/*
Shell command to run this test:

	go test -v ./local_test.go -args \
	--LOCAL_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--LOCAL_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--LOCAL_OUTPUTCERTPATH="/path/to/your-output-cert" \
	--LOCAL_OUTPUTKEYPATH="/path/to/your-output-key" \
	--LOCAL_PFXPASSWORD="your-pfx-password" \
	--LOCAL_JKSALIAS="your-jks-alias" \
	--LOCAL_JKSKEYPASS="your-jks-keypass" \
	--LOCAL_JKSSTOREPASS="your-jks-storepass" \
	--LOCAL_SHELLENV="sh" \
	--LOCAL_PRECOMMAND="echo 'hello world'" \
	--LOCAL_POSTCOMMAND="echo 'bye-bye world'"
*/
func TestDeploy(t *testing.T) {
	flag.Parse()

	t.Run("Deploy_PEM", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("OUTPUTCERTPATH: %v", fOutputCertPath),
			fmt.Sprintf("OUTPUTKEYPATH: %v", fOutputKeyPath),
			fmt.Sprintf("SHELLENV: %v", fShellEnv),
			fmt.Sprintf("PRECOMMAND: %v", fPreCommand),
			fmt.Sprintf("POSTCOMMAND: %v", fPostCommand),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			OutputFormat:   provider.OUTPUT_FORMAT_PEM,
			OutputCertPath: fOutputCertPath + ".pem",
			OutputKeyPath:  fOutputKeyPath + ".pem",
			ShellEnv:       fShellEnv,
			PreCommand:     fPreCommand,
			PostCommand:    fPostCommand,
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

		fstat1, err := os.Stat(fOutputCertPath + ".pem")
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		} else if fstat1.Size() == 0 {
			t.Errorf("err: empty output certificate file")
			return
		}

		fstat2, err := os.Stat(fOutputKeyPath + ".pem")
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		} else if fstat2.Size() == 0 {
			t.Errorf("err: empty output private key file")
			return
		}

		t.Logf("ok: %v", res)
	})

	t.Run("Deploy_PFX", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("OUTPUTCERTPATH: %v", fOutputCertPath),
			fmt.Sprintf("PFXPASSWORD: %v", fPfxPassword),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			OutputFormat:   provider.OUTPUT_FORMAT_PFX,
			OutputCertPath: fOutputCertPath + ".pfx",
			PfxPassword:    fPfxPassword,
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

		fstat, err := os.Stat(fOutputCertPath + ".pfx")
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		} else if fstat.Size() == 0 {
			t.Errorf("err: empty output certificate file")
			return
		}

		t.Logf("ok: %v", res)
	})

	t.Run("Deploy_JKS", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("OUTPUTCERTPATH: %v", fOutputCertPath),
			fmt.Sprintf("JKSALIAS: %v", fJksAlias),
			fmt.Sprintf("JKSKEYPASS: %v", fJksKeypass),
			fmt.Sprintf("JKSSTOREPASS: %v", fJksStorepass),
		}, "\n"))

		provider, err := provider.NewDeployer(&provider.DeployerConfig{
			OutputFormat:   provider.OUTPUT_FORMAT_JKS,
			OutputCertPath: fOutputCertPath + ".jks",
			JksAlias:       fJksAlias,
			JksKeypass:     fJksKeypass,
			JksStorepass:   fJksStorepass,
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

		fstat, err := os.Stat(fOutputCertPath + ".jks")
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		} else if fstat.Size() == 0 {
			t.Errorf("err: empty output certificate file")
			return
		}

		t.Logf("ok: %v", res)
	})
}
