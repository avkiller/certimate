package tencentcloudclb

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/samber/lo"
	tcclb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/clb/v20180317"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"

	"github.com/certimate-go/certimate/pkg/core"
	sslmgrsp "github.com/certimate-go/certimate/pkg/core/ssl-manager/providers/tencentcloud-ssl"
)

type SSLDeployerProviderConfig struct {
	// 腾讯云 SecretId。
	SecretId string `json:"secretId"`
	// 腾讯云 SecretKey。
	SecretKey string `json:"secretKey"`
	// 腾讯云接口端点。
	Endpoint string `json:"endpoint,omitempty"`
	// 腾讯云地域。
	Region string `json:"region"`
	// 部署资源类型。
	ResourceType ResourceType `json:"resourceType"`
	// 负载均衡器 ID。
	// 部署资源类型为 [RESOURCE_TYPE_SSLDEPLOY]、[RESOURCE_TYPE_LOADBALANCER]、[RESOURCE_TYPE_RULEDOMAIN] 时必填。
	LoadbalancerId string `json:"loadbalancerId,omitempty"`
	// 负载均衡监听 ID。
	// 部署资源类型为 [RESOURCE_TYPE_SSLDEPLOY]、[RESOURCE_TYPE_LOADBALANCER]、[RESOURCE_TYPE_LISTENER]、[RESOURCE_TYPE_RULEDOMAIN] 时必填。
	ListenerId string `json:"listenerId,omitempty"`
	// SNI 域名或七层转发规则域名（支持泛域名）。
	// 部署资源类型为 [RESOURCE_TYPE_SSLDEPLOY] 时选填；部署资源类型为 [RESOURCE_TYPE_RULEDOMAIN] 时必填。
	Domain string `json:"domain,omitempty"`
}

type SSLDeployerProvider struct {
	config     *SSLDeployerProviderConfig
	logger     *slog.Logger
	sdkClient  *tcclb.Client
	sslManager core.SSLManager
}

var _ core.SSLDeployer = (*SSLDeployerProvider)(nil)

func NewSSLDeployerProvider(config *SSLDeployerProviderConfig) (*SSLDeployerProvider, error) {
	if config == nil {
		return nil, errors.New("the configuration of the ssl deployer provider is nil")
	}

	client, err := createSDKClient(config.SecretId, config.SecretKey, config.Endpoint, config.Region)
	if err != nil {
		return nil, fmt.Errorf("could not create sdk client: %w", err)
	}

	sslmgr, err := sslmgrsp.NewSSLManagerProvider(&sslmgrsp.SSLManagerProviderConfig{
		SecretId:  config.SecretId,
		SecretKey: config.SecretKey,
		Endpoint: lo.
			If(strings.HasSuffix(config.Endpoint, "intl.tencentcloudapi.com"), "ssl.intl.tencentcloudapi.com"). // 国际站使用独立的接口端点
			Else(""),
	})
	if err != nil {
		return nil, fmt.Errorf("could not create ssl manager: %w", err)
	}

	return &SSLDeployerProvider{
		config:     config,
		logger:     slog.Default(),
		sdkClient:  client,
		sslManager: sslmgr,
	}, nil
}

func (d *SSLDeployerProvider) SetLogger(logger *slog.Logger) {
	if logger == nil {
		d.logger = slog.New(slog.DiscardHandler)
	} else {
		d.logger = logger
	}

	d.sslManager.SetLogger(logger)
}

func (d *SSLDeployerProvider) Deploy(ctx context.Context, certPEM string, privkeyPEM string) (*core.SSLDeployResult, error) {
	// 上传证书
	upres, err := d.sslManager.Upload(ctx, certPEM, privkeyPEM)
	if err != nil {
		return nil, fmt.Errorf("failed to upload certificate file: %w", err)
	} else {
		d.logger.Info("ssl certificate uploaded", slog.Any("result", upres))
	}

	// 根据部署资源类型决定部署方式
	switch d.config.ResourceType {
	case RESOURCE_TYPE_LOADBALANCER:
		if err := d.deployToLoadbalancer(ctx, upres.CertId); err != nil {
			return nil, err
		}

	case RESOURCE_TYPE_LISTENER:
		if err := d.deployToListener(ctx, upres.CertId); err != nil {
			return nil, err
		}

	case RESOURCE_TYPE_RULEDOMAIN:
		if err := d.deployToRuleDomain(ctx, upres.CertId); err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("unsupported resource type '%s'", d.config.ResourceType)
	}

	return &core.SSLDeployResult{}, nil
}

func (d *SSLDeployerProvider) deployToLoadbalancer(ctx context.Context, cloudCertId string) error {
	if d.config.LoadbalancerId == "" {
		return errors.New("config `loadbalancerId` is required")
	}

	// 查询监听器列表
	// REF: https://cloud.tencent.com/document/api/214/30686
	listenerIds := make([]string, 0)
	describeListenersReq := tcclb.NewDescribeListenersRequest()
	describeListenersReq.LoadBalancerId = common.StringPtr(d.config.LoadbalancerId)
	describeListenersResp, err := d.sdkClient.DescribeListeners(describeListenersReq)
	d.logger.Debug("sdk request 'clb.DescribeListeners'", slog.Any("request", describeListenersReq), slog.Any("response", describeListenersResp))
	if err != nil {
		return fmt.Errorf("failed to execute sdk request 'clb.DescribeListeners': %w", err)
	} else {
		if describeListenersResp.Response.Listeners != nil {
			for _, listener := range describeListenersResp.Response.Listeners {
				if listener.Protocol == nil || (*listener.Protocol != "HTTPS" && *listener.Protocol != "TCP_SSL" && *listener.Protocol != "QUIC") {
					continue
				}

				listenerIds = append(listenerIds, *listener.ListenerId)
			}
		}
	}

	// 遍历更新监听器证书
	if len(listenerIds) == 0 {
		d.logger.Info("no clb listeners to deploy")
	} else {
		d.logger.Info("found https/tcpssl/quic listeners to deploy", slog.Any("listenerIds", listenerIds))
		var errs []error

		for _, listenerId := range listenerIds {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				if err := d.modifyListenerCertificate(ctx, d.config.LoadbalancerId, listenerId, cloudCertId); err != nil {
					errs = append(errs, err)
				}
			}
		}

		if len(errs) > 0 {
			return errors.Join(errs...)
		}
	}

	return nil
}

func (d *SSLDeployerProvider) deployToListener(ctx context.Context, cloudCertId string) error {
	if d.config.LoadbalancerId == "" {
		return errors.New("config `loadbalancerId` is required")
	}
	if d.config.ListenerId == "" {
		return errors.New("config `listenerId` is required")
	}

	// 更新监听器证书
	if err := d.modifyListenerCertificate(ctx, d.config.LoadbalancerId, d.config.ListenerId, cloudCertId); err != nil {
		return err
	}

	return nil
}

func (d *SSLDeployerProvider) deployToRuleDomain(ctx context.Context, cloudCertId string) error {
	if d.config.LoadbalancerId == "" {
		return errors.New("config `loadbalancerId` is required")
	}
	if d.config.ListenerId == "" {
		return errors.New("config `listenerId` is required")
	}
	if d.config.Domain == "" {
		return errors.New("config `domain` is required")
	}

	// 修改负载均衡七层监听器转发规则的域名级别属性
	// REF: https://cloud.tencent.com/document/api/214/38092
	modifyDomainAttributesReq := tcclb.NewModifyDomainAttributesRequest()
	modifyDomainAttributesReq.LoadBalancerId = common.StringPtr(d.config.LoadbalancerId)
	modifyDomainAttributesReq.ListenerId = common.StringPtr(d.config.ListenerId)
	modifyDomainAttributesReq.Domain = common.StringPtr(d.config.Domain)
	modifyDomainAttributesReq.Certificate = &tcclb.CertificateInput{
		SSLMode: common.StringPtr("UNIDIRECTIONAL"),
		CertId:  common.StringPtr(cloudCertId),
	}
	modifyDomainAttributesResp, err := d.sdkClient.ModifyDomainAttributes(modifyDomainAttributesReq)
	d.logger.Debug("sdk request 'clb.ModifyDomainAttributes'", slog.Any("request", modifyDomainAttributesReq), slog.Any("response", modifyDomainAttributesResp))
	if err != nil {
		return fmt.Errorf("failed to execute sdk request 'clb.ModifyDomainAttributes': %w", err)
	}

	// 循环查询异步任务状态，等待任务状态变更
	// REF: https://cloud.tencent.com/document/product/214/30683
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		describeTaskStatusReq := tcclb.NewDescribeTaskStatusRequest()
		describeTaskStatusReq.TaskId = modifyDomainAttributesResp.Response.RequestId
		describeTaskStatusResp, err := d.sdkClient.DescribeTaskStatus(describeTaskStatusReq)
		d.logger.Debug("sdk request 'clb.DescribeTaskStatus'", slog.Any("request", describeTaskStatusReq), slog.Any("response", describeTaskStatusResp))
		if err != nil {
			return fmt.Errorf("failed to execute sdk request 'clb.DescribeTaskStatus': %w", err)
		}

		if describeTaskStatusResp.Response.Status == nil || *describeTaskStatusResp.Response.Status == 1 {
			return errors.New("unexpected tencentcloud task status")
		} else if *describeTaskStatusResp.Response.Status == 0 {
			break
		}

		d.logger.Info("waiting for tencentcloud task completion ...")
		time.Sleep(time.Second * 5)
	}

	return nil
}

func (d *SSLDeployerProvider) modifyListenerCertificate(ctx context.Context, cloudLoadbalancerId, cloudListenerId, cloudCertId string) error {
	// 查询负载均衡的监听器列表
	// REF: https://cloud.tencent.com/document/api/214/30686
	describeListenersReq := tcclb.NewDescribeListenersRequest()
	describeListenersReq.LoadBalancerId = common.StringPtr(cloudLoadbalancerId)
	describeListenersReq.ListenerIds = common.StringPtrs([]string{cloudListenerId})
	describeListenersResp, err := d.sdkClient.DescribeListeners(describeListenersReq)
	d.logger.Debug("sdk request 'clb.DescribeListeners'", slog.Any("request", describeListenersReq), slog.Any("response", describeListenersResp))
	if err != nil {
		return fmt.Errorf("failed to execute sdk request 'clb.DescribeListeners': %w", err)
	} else if len(describeListenersResp.Response.Listeners) == 0 {
		return fmt.Errorf("listener %s not found", cloudListenerId)
	}

	// 修改监听器属性
	// REF: https://cloud.tencent.com/document/api/214/30681
	modifyListenerReq := tcclb.NewModifyListenerRequest()
	modifyListenerReq.LoadBalancerId = common.StringPtr(cloudLoadbalancerId)
	modifyListenerReq.ListenerId = common.StringPtr(cloudListenerId)
	modifyListenerReq.Certificate = &tcclb.CertificateInput{CertId: common.StringPtr(cloudCertId)}
	if describeListenersResp.Response.Listeners[0].Certificate != nil && describeListenersResp.Response.Listeners[0].Certificate.SSLMode != nil {
		modifyListenerReq.Certificate.SSLMode = describeListenersResp.Response.Listeners[0].Certificate.SSLMode
		modifyListenerReq.Certificate.CertCaId = describeListenersResp.Response.Listeners[0].Certificate.CertCaId
	} else {
		modifyListenerReq.Certificate.SSLMode = common.StringPtr("UNIDIRECTIONAL")
	}
	modifyListenerResp, err := d.sdkClient.ModifyListener(modifyListenerReq)
	d.logger.Debug("sdk request 'clb.ModifyListener'", slog.Any("request", modifyListenerReq), slog.Any("response", modifyListenerResp))
	if err != nil {
		return fmt.Errorf("failed to execute sdk request 'clb.ModifyListener': %w", err)
	}

	// 循环查询异步任务状态，等待任务状态变更
	// REF: https://cloud.tencent.com/document/product/214/30683
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		describeTaskStatusReq := tcclb.NewDescribeTaskStatusRequest()
		describeTaskStatusReq.TaskId = modifyListenerResp.Response.RequestId
		describeTaskStatusResp, err := d.sdkClient.DescribeTaskStatus(describeTaskStatusReq)
		d.logger.Debug("sdk request 'clb.DescribeTaskStatus'", slog.Any("request", describeTaskStatusReq), slog.Any("response", describeTaskStatusResp))
		if err != nil {
			return fmt.Errorf("failed to execute sdk request 'clb.DescribeTaskStatus': %w", err)
		}

		if describeTaskStatusResp.Response.Status == nil || *describeTaskStatusResp.Response.Status == 1 {
			return errors.New("unexpected tencentcloud task status")
		} else if *describeTaskStatusResp.Response.Status == 0 {
			break
		}

		d.logger.Info("waiting for tencentcloud task completion ...")
		time.Sleep(time.Second * 5)
	}

	return nil
}

func createSDKClient(secretId, secretKey, endpoint, region string) (*tcclb.Client, error) {
	credential := common.NewCredential(secretId, secretKey)

	cpf := profile.NewClientProfile()
	if endpoint != "" {
		cpf.HttpProfile.Endpoint = endpoint
	}

	client, err := tcclb.NewClient(credential, region, cpf)
	if err != nil {
		return nil, err
	}

	return client, nil
}
