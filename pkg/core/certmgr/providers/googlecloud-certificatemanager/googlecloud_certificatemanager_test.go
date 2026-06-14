package googlecloudcertificatemanager_test

import (
	"testing"

	"github.com/certimate-go/certimate/pkg/core/certmgr/internal/tester"
	impl "github.com/certimate-go/certimate/pkg/core/certmgr/providers/googlecloud-certificatemanager"
)

var (
	fp                 = tester.Args("GOOGLECLOUDCERTIFICATEMANAGER_")
	fTestCertPath      string
	fTestKeyPath       string
	fServiceAccountKey string
)

func init() {
	fp.DefineString(&fTestCertPath, "TESTCERTPATH")
	fp.DefineString(&fTestKeyPath, "TESTKEYPATH")
	fp.DefineString(&fServiceAccountKey, "SERVICEACCOUNTKEY")
}

/*
Shell command to run this test:

	go test -v ./googlecloud_certificatemanager_test.go -args \
	--GOOGLECLOUDCERTIFICATEMANAGER_TESTCERTPATH="/path/to/your-test-cert.pem" \
	--GOOGLECLOUDCERTIFICATEMANAGER_TESTKEYPATH="/path/to/your-test-key.pem" \
	--GOOGLECLOUDCERTIFICATEMANAGER_SERVICEACCOUNTKEY="{...}"
*/
func TestProvider(t *testing.T) {
	fp.Parse()

	t.Run("Upload", func(t *testing.T) {
		provider, err := impl.NewCertmgr(&impl.CertmgrConfig{
			ServiceAccountKey: fServiceAccountKey,
		})
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		}

		tester.TestUpload(t, provider, tester.TestUploadArgs{CertPath: fTestCertPath, KeyPath: fTestKeyPath})
	})
}
