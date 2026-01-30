package certapply

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-acme/lego/v4/acme"
	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/challenge/http01"
	"github.com/go-acme/lego/v4/log"
	"github.com/samber/lo"

	"github.com/certimate-go/certimate/internal/certapply/applicators"
	"github.com/certimate-go/certimate/internal/domain"
)

type ObtainCertificateRequest struct {
	Domains    []string
	KeyType    certcrypto.KeyType
	ValidityTo time.Time

	// 提供商相关
	ChallengeType          string
	Provider               string
	ProviderAccessConfig   map[string]any
	ProviderExtendedConfig map[string]any

	// 解析相关
	DisableFollowCNAME bool
	Nameservers        []string

	// DNS-01 质询相关
	DnsPropagationWait    int32
	DnsPropagationTimeout int32
	DnsTTL                int32

	// HTTP-01 质询相关
	HttpDelayWait int32

	// ACME 相关
	ACMEProfile string

	// ARI 相关
	ARIReplacesAcctUrl string
	ARIReplacesCertId  string
}

type ObtainCertificateResponse struct {
	CSR                  string
	FullChainCertificate string
	IssuerCertificate    string
	PrivateKey           string
	ACMEAcctUrl          string
	ACMECertUrl          string
	ACMECertStableUrl    string
	ARIReplaced          bool
}

func (c *ACMEClient) ObtainCertificate(ctx context.Context, request *ObtainCertificateRequest) (*ObtainCertificateResponse, error) {
	type result struct {
		res *ObtainCertificateResponse
		err error
	}

	done := make(chan result, 1)

	go func() {
		res, err := c.sendObtainCertificateRequest(request)
		done <- result{res, err}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case r := <-done:
		return r.res, r.err
	}
}

func (c *ACMEClient) sendObtainCertificateRequest(request *ObtainCertificateRequest) (*ObtainCertificateResponse, error) {
	if request == nil {
		return nil, errors.New("the request is nil")
	}

	os.Setenv("LEGO_DISABLE_CNAME_SUPPORT", strconv.FormatBool(request.DisableFollowCNAME))

	switch request.ChallengeType {
	case "dns-01":
		{
			providerFactory, err := applicators.ACMEDns01Registries.Get(domain.ACMEDns01ProviderType(request.Provider))
			if err != nil {
				return nil, err
			}

			provider, err := providerFactory(&applicators.ProviderFactoryOptions{
				ProviderAccessConfig:   request.ProviderAccessConfig,
				ProviderExtendedConfig: request.ProviderExtendedConfig,
				DnsPropagationWait:     request.DnsPropagationWait,
				DnsPropagationTimeout:  request.DnsPropagationTimeout,
				DnsTTL:                 request.DnsTTL,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to initialize dns-01 provider '%s': %w", request.Provider, err)
			}

			c.client.Challenge.SetDNS01Provider(provider,
				dns01.CondOption(
					len(request.Nameservers) > 0,
					dns01.AddRecursiveNameservers(dns01.ParseNameservers(request.Nameservers)),
				),
				dns01.CondOption(
					request.DnsPropagationWait > 0,
					dns01.PropagationWait(time.Duration(request.DnsPropagationWait)*time.Second, true),
				),
				dns01.CondOption(
					len(request.Nameservers) > 0 || request.DnsPropagationWait > 0,
					dns01.DisableAuthoritativeNssPropagationRequirement(),
				),
			)
		}

	case "http-01":
		{
			providerFactory, err := applicators.ACMEHttp01Registries.Get(domain.ACMEHttp01ProviderType(request.Provider))
			if err != nil {
				return nil, err
			}

			provider, err := providerFactory(&applicators.ProviderFactoryOptions{
				ProviderAccessConfig:   request.ProviderAccessConfig,
				ProviderExtendedConfig: request.ProviderExtendedConfig,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to initialize http-01 provider '%s': %w", request.Provider, err)
			}

			c.client.Challenge.SetHTTP01Provider(provider,
				http01.SetDelay(time.Duration(request.HttpDelayWait)*time.Second),
			)
		}

	default:
		return nil, fmt.Errorf("unsupported challenge type: '%s'", request.ChallengeType)
	}

	req := certificate.ObtainRequest{
		Domains:        request.Domains,
		Bundle:         true,
		Profile:        request.ACMEProfile,
		NotAfter:       request.ValidityTo,
		ReplacesCertID: lo.If(request.ARIReplacesAcctUrl == c.account.ACMEAcctUrl, request.ARIReplacesCertId).Else(""),
	}
	resp, err := c.client.Certificate.Obtain(req)
	if err != nil {
		ariErr := &acme.AlreadyReplacedError{}
		if !errors.As(err, &ariErr) {
			return nil, err
		}

		log.Warnf("the certificate has already been replaced, try to obtain again without ARI ...")

		// reset ARI and retry if failure
		req.ReplacesCertID = ""
		resp, err = c.client.Certificate.Obtain(req)
		if err != nil {
			return nil, err
		}
	}

	return &ObtainCertificateResponse{
		CSR:                  strings.TrimSpace(string(resp.CSR)),
		FullChainCertificate: strings.TrimSpace(string(resp.Certificate)),
		IssuerCertificate:    strings.TrimSpace(string(resp.IssuerCertificate)),
		PrivateKey:           strings.TrimSpace(string(resp.PrivateKey)),
		ACMEAcctUrl:          c.account.ACMEAcctUrl,
		ACMECertUrl:          resp.CertURL,
		ACMECertStableUrl:    resp.CertStableURL,
		ARIReplaced:          req.ReplacesCertID != "",
	}, nil
}
