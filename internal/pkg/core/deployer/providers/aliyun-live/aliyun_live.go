﻿package aliyunlive

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	aliyunOpen "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	aliyunLive "github.com/alibabacloud-go/live-20161101/client"
	"github.com/alibabacloud-go/tea/tea"
	xerrors "github.com/pkg/errors"

	"github.com/usual2970/certimate/internal/pkg/core/deployer"
	"github.com/usual2970/certimate/internal/pkg/core/logger"
)

type AliyunLiveDeployerConfig struct {
	// 阿里云 AccessKeyId。
	AccessKeyId string `json:"accessKeyId"`
	// 阿里云 AccessKeySecret。
	AccessKeySecret string `json:"accessKeySecret"`
	// 阿里云地域。
	Region string `json:"region"`
	// 直播流域名（支持泛域名）。
	Domain string `json:"domain"`
}

type AliyunLiveDeployer struct {
	config    *AliyunLiveDeployerConfig
	logger    logger.Logger
	sdkClient *aliyunLive.Client
}

var _ deployer.Deployer = (*AliyunLiveDeployer)(nil)

func New(config *AliyunLiveDeployerConfig) (*AliyunLiveDeployer, error) {
	return NewWithLogger(config, logger.NewNilLogger())
}

func NewWithLogger(config *AliyunLiveDeployerConfig, logger logger.Logger) (*AliyunLiveDeployer, error) {
	if config == nil {
		return nil, errors.New("config is nil")
	}

	if logger == nil {
		return nil, errors.New("logger is nil")
	}

	client, err := createSdkClient(config.AccessKeyId, config.AccessKeySecret, config.Region)
	if err != nil {
		return nil, xerrors.Wrap(err, "failed to create sdk client")
	}

	return &AliyunLiveDeployer{
		logger:    logger,
		config:    config,
		sdkClient: client,
	}, nil
}

func (d *AliyunLiveDeployer) Deploy(ctx context.Context, certPem string, privkeyPem string) (*deployer.DeployResult, error) {
	// "*.example.com" → ".example.com"，适配阿里云 Live 要求的泛域名格式
	domain := strings.TrimPrefix(d.config.Domain, "*")

	// 设置域名证书
	// REF: https://help.aliyun.com/zh/live/developer-reference/api-live-2016-11-01-setlivedomaincertificate
	setLiveDomainSSLCertificateReq := &aliyunLive.SetLiveDomainCertificateRequest{
		DomainName:  tea.String(domain),
		CertName:    tea.String(fmt.Sprintf("certimate-%d", time.Now().UnixMilli())),
		CertType:    tea.String("upload"),
		SSLProtocol: tea.String("on"),
		SSLPub:      tea.String(certPem),
		SSLPri:      tea.String(privkeyPem),
	}
	setLiveDomainSSLCertificateResp, err := d.sdkClient.SetLiveDomainCertificate(setLiveDomainSSLCertificateReq)
	if err != nil {
		return nil, xerrors.Wrap(err, "failed to execute sdk request 'live.SetLiveDomainCertificate'")
	}

	d.logger.Logt("已设置域名证书", setLiveDomainSSLCertificateResp)

	return &deployer.DeployResult{}, nil
}

func createSdkClient(accessKeyId, accessKeySecret, region string) (*aliyunLive.Client, error) {
	// 接入点一览 https://help.aliyun.com/zh/live/developer-reference/api-live-2016-11-01-endpoint
	var endpoint string
	switch region {
	case
		"cn-qingdao",
		"cn-beijing",
		"cn-shanghai",
		"cn-shenzhen",
		"ap-northeast-1",
		"ap-southeast-5",
		"me-central-1":
		endpoint = "live.aliyuncs.com"
	default:
		endpoint = fmt.Sprintf("live.%s.aliyuncs.com", region)
	}

	config := &aliyunOpen.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
		Endpoint:        tea.String(endpoint),
	}

	client, err := aliyunLive.NewClient(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}
