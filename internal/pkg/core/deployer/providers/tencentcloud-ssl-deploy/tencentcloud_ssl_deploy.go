﻿package tencentcloudssldeploy

import (
	"context"
	"errors"
	"fmt"
	"time"

	xerrors "github.com/pkg/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tcSsl "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssl/v20191205"

	"github.com/usual2970/certimate/internal/pkg/core/deployer"
	"github.com/usual2970/certimate/internal/pkg/core/logger"
	"github.com/usual2970/certimate/internal/pkg/core/uploader"
	uploaderp "github.com/usual2970/certimate/internal/pkg/core/uploader/providers/tencentcloud-ssl"
)

type TencentCloudSSLDeployDeployerConfig struct {
	// 腾讯云 SecretId。
	SecretId string `json:"secretId"`
	// 腾讯云 SecretKey。
	SecretKey string `json:"secretKey"`
	// 腾讯云地域。
	Region string `json:"region"`
	// 腾讯云云资源类型。
	ResourceType string `json:"resourceType"`
	// 腾讯云云资源 ID 数组。
	ResourceIds []string `json:"resourceIds"`
}

type TencentCloudSSLDeployDeployer struct {
	config      *TencentCloudSSLDeployDeployerConfig
	logger      logger.Logger
	sdkClient   *tcSsl.Client
	sslUploader uploader.Uploader
}

var _ deployer.Deployer = (*TencentCloudSSLDeployDeployer)(nil)

func New(config *TencentCloudSSLDeployDeployerConfig) (*TencentCloudSSLDeployDeployer, error) {
	return NewWithLogger(config, logger.NewNilLogger())
}

func NewWithLogger(config *TencentCloudSSLDeployDeployerConfig, logger logger.Logger) (*TencentCloudSSLDeployDeployer, error) {
	if config == nil {
		return nil, errors.New("config is nil")
	}

	if logger == nil {
		return nil, errors.New("logger is nil")
	}

	client, err := createSdkClient(config.SecretId, config.SecretKey, config.Region)
	if err != nil {
		return nil, xerrors.Wrap(err, "failed to create sdk client")
	}

	uploader, err := uploaderp.New(&uploaderp.TencentCloudSSLUploaderConfig{
		SecretId:  config.SecretId,
		SecretKey: config.SecretKey,
	})
	if err != nil {
		return nil, xerrors.Wrap(err, "failed to create ssl uploader")
	}

	return &TencentCloudSSLDeployDeployer{
		logger:      logger,
		config:      config,
		sdkClient:   client,
		sslUploader: uploader,
	}, nil
}

func (d *TencentCloudSSLDeployDeployer) Deploy(ctx context.Context, certPem string, privkeyPem string) (*deployer.DeployResult, error) {
	if d.config.ResourceType == "" {
		return nil, errors.New("config `resourceType` is required")
	}
	if len(d.config.ResourceIds) == 0 {
		return nil, errors.New("config `resourceIds` is required")
	}

	// 上传证书到 SSL
	upres, err := d.sslUploader.Upload(ctx, certPem, privkeyPem)
	if err != nil {
		return nil, xerrors.Wrap(err, "failed to upload certificate file")
	}

	d.logger.Logt("certificate file uploaded", upres)

	// 证书部署到云资源实例列表
	// REF: https://cloud.tencent.com/document/product/400/91667
	deployCertificateInstanceReq := tcSsl.NewDeployCertificateInstanceRequest()
	deployCertificateInstanceReq.CertificateId = common.StringPtr(upres.CertId)
	deployCertificateInstanceReq.ResourceType = common.StringPtr(d.config.ResourceType)
	deployCertificateInstanceReq.InstanceIdList = common.StringPtrs(d.config.ResourceIds)
	deployCertificateInstanceReq.Status = common.Int64Ptr(1)
	deployCertificateInstanceResp, err := d.sdkClient.DeployCertificateInstance(deployCertificateInstanceReq)
	if err != nil {
		return nil, xerrors.Wrap(err, "failed to execute sdk request 'ssl.DeployCertificateInstance'")
	} else if deployCertificateInstanceResp.Response == nil || deployCertificateInstanceResp.Response.DeployRecordId == nil {
		return nil, errors.New("failed to create deploy record")
	}

	d.logger.Logt("已部署证书到云资源实例", deployCertificateInstanceResp.Response)

	// 循环获取部署任务详情，等待任务状态变更
	// REF: https://cloud.tencent.com.cn/document/api/400/91658
	for {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

		describeHostDeployRecordDetailReq := tcSsl.NewDescribeHostDeployRecordDetailRequest()
		describeHostDeployRecordDetailReq.DeployRecordId = common.StringPtr(fmt.Sprintf("%d", *deployCertificateInstanceResp.Response.DeployRecordId))
		describeHostDeployRecordDetailReq.Limit = common.Uint64Ptr(100)
		describeHostDeployRecordDetailResp, err := d.sdkClient.DescribeHostDeployRecordDetail(describeHostDeployRecordDetailReq)
		if err != nil {
			return nil, xerrors.Wrap(err, "failed to execute sdk request 'ssl.DescribeHostDeployRecordDetail'")
		}

		if describeHostDeployRecordDetailResp.Response.TotalCount == nil {
			return nil, errors.New("部署任务状态异常")
		} else {
			acc := int64(0)
			if describeHostDeployRecordDetailResp.Response.SuccessTotalCount != nil {
				acc += *describeHostDeployRecordDetailResp.Response.SuccessTotalCount
			}
			if describeHostDeployRecordDetailResp.Response.FailedTotalCount != nil {
				acc += *describeHostDeployRecordDetailResp.Response.FailedTotalCount
			}

			if acc == *describeHostDeployRecordDetailResp.Response.TotalCount {
				d.logger.Logt("已获取部署任务详情", describeHostDeployRecordDetailResp)
				break
			}
		}

		d.logger.Logt("部署任务未完成 ...")
		time.Sleep(time.Second * 5)
	}

	return &deployer.DeployResult{}, nil
}

func createSdkClient(secretId, secretKey, region string) (*tcSsl.Client, error) {
	credential := common.NewCredential(secretId, secretKey)

	client, err := tcSsl.NewClient(credential, region, profile.NewClientProfile())
	if err != nil {
		return nil, err
	}

	return client, nil
}
