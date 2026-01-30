package ssh

import (
	"errors"
	"fmt"
	"net"
	"path/filepath"
	"strconv"

	"github.com/go-acme/lego/v4/challenge/http01"
	"golang.org/x/crypto/ssh"

	"github.com/certimate-go/certimate/pkg/core"
	xssh "github.com/certimate-go/certimate/pkg/utils/ssh"
)

type ServerConfig struct {
	// SSH 主机。
	// 零值时默认值 "localhost"。
	SshHost string `json:"sshHost,omitempty"`
	// SSH 端口。
	// 零值时默认值 22。
	SshPort int32 `json:"sshPort,omitempty"`
	// SSH 认证方式。
	// 可取值 "none"、"password"、"key"。
	// 零值时根据有无密码或私钥字段决定。
	SshAuthMethod string `json:"sshAuthMethod,omitempty"`
	// SSH 登录用户名。
	// 零值时默认值 "root"。
	SshUsername string `json:"sshUsername,omitempty"`
	// SSH 登录密码。
	SshPassword string `json:"sshPassword,omitempty"`
	// SSH 登录私钥。
	SshKey string `json:"sshKey,omitempty"`
	// SSH 登录私钥口令。
	SshKeyPassphrase string `json:"sshKeyPassphrase,omitempty"`
}

type ChallengeProviderConfig struct {
	ServerConfig

	// 跳板机配置数组。
	JumpServers []ServerConfig `json:"jumpServers,omitempty"`
	// 是否回退使用 SCP。
	UseSCP bool `json:"useSCP,omitempty"`
	// 网站根目录路径。
	WebRootPath string `json:"webRootPath"`
}

func NewChallengeProvider(config *ChallengeProviderConfig) (core.ACMEChallenger, error) {
	if config == nil {
		return nil, errors.New("the configuration of the acme challenge provider is nil")
	}

	provider := &provider{config: config}
	return provider, nil
}

type provider struct {
	config *ChallengeProviderConfig
}

func (p *provider) Present(domain, token, keyAuth string) error {
	var err error

	// 创建 TCP 链接
	var targetConn net.Conn
	if len(p.config.JumpServers) > 0 {
		var jumpClient *ssh.Client
		for i, jumpServerConf := range p.config.JumpServers {
			var jumpConn net.Conn
			// 第一个连接是主机发起，后续通过跳板机发起
			if jumpClient == nil {
				jumpConn, err = net.Dial("tcp", net.JoinHostPort(jumpServerConf.SshHost, strconv.Itoa(int(jumpServerConf.SshPort))))
			} else {
				jumpConn, err = jumpClient.Dial("tcp", net.JoinHostPort(jumpServerConf.SshHost, strconv.Itoa(int(jumpServerConf.SshPort))))
			}
			if err != nil {
				return fmt.Errorf("failed to connect to jump server [%d]: %w", i+1, err)
			}
			defer jumpConn.Close()

			newClient, err := p.createSshClient(
				jumpConn,
				jumpServerConf.SshHost,
				jumpServerConf.SshPort,
				jumpServerConf.SshAuthMethod,
				jumpServerConf.SshUsername,
				jumpServerConf.SshPassword,
				jumpServerConf.SshKey,
				jumpServerConf.SshKeyPassphrase,
			)
			if err != nil {
				return fmt.Errorf("failed to create jump server ssh client[%d]: %w", i+1, err)
			}
			defer newClient.Close()

			jumpClient = newClient
		}

		// 通过跳板机发起 TCP 连接到目标服务器
		targetConn, err = jumpClient.Dial("tcp", net.JoinHostPort(p.config.SshHost, strconv.Itoa(int(p.config.SshPort))))
		if err != nil {
			return fmt.Errorf("failed to connect to target server: %w", err)
		}
	} else {
		// 直接发起 TCP 连接到目标服务器
		targetConn, err = net.Dial("tcp", net.JoinHostPort(p.config.SshHost, strconv.Itoa(int(p.config.SshPort))))
		if err != nil {
			return fmt.Errorf("failed to connect to target server: %w", err)
		}
	}
	defer targetConn.Close()

	// 创建 SSH 客户端
	client, err := p.createSshClient(
		targetConn,
		p.config.SshHost,
		p.config.SshPort,
		p.config.SshAuthMethod,
		p.config.SshUsername,
		p.config.SshPassword,
		p.config.SshKey,
		p.config.SshKeyPassphrase,
	)
	if err != nil {
		return fmt.Errorf("failed to create ssh client: %w", err)
	}
	defer client.Close()

	// 写入质询文件
	challengeFilePath := filepath.Join(p.config.WebRootPath, http01.ChallengePath(token))
	if err := xssh.WriteRemoteString(client, challengeFilePath, keyAuth, p.config.UseSCP); err != nil {
		return fmt.Errorf("failed to write file in webroot for HTTP challenge: %w", err)
	}

	return nil
}

func (p *provider) CleanUp(domain, token, keyAuth string) error {
	var err error

	// 创建 TCP 链接
	var targetConn net.Conn
	if len(p.config.JumpServers) > 0 {
		var jumpClient *ssh.Client
		for i, jumpServerConf := range p.config.JumpServers {
			var jumpConn net.Conn
			// 第一个连接是主机发起，后续通过跳板机发起
			if jumpClient == nil {
				jumpConn, err = net.Dial("tcp", net.JoinHostPort(jumpServerConf.SshHost, strconv.Itoa(int(jumpServerConf.SshPort))))
			} else {
				jumpConn, err = jumpClient.Dial("tcp", net.JoinHostPort(jumpServerConf.SshHost, strconv.Itoa(int(jumpServerConf.SshPort))))
			}
			if err != nil {
				return fmt.Errorf("failed to connect to jump server [%d]: %w", i+1, err)
			}
			defer jumpConn.Close()

			newClient, err := p.createSshClient(
				jumpConn,
				jumpServerConf.SshHost,
				jumpServerConf.SshPort,
				jumpServerConf.SshAuthMethod,
				jumpServerConf.SshUsername,
				jumpServerConf.SshPassword,
				jumpServerConf.SshKey,
				jumpServerConf.SshKeyPassphrase,
			)
			if err != nil {
				return fmt.Errorf("failed to create jump server ssh client[%d]: %w", i+1, err)
			}
			defer newClient.Close()

			jumpClient = newClient
		}

		// 通过跳板机发起 TCP 连接到目标服务器
		targetConn, err = jumpClient.Dial("tcp", net.JoinHostPort(p.config.SshHost, strconv.Itoa(int(p.config.SshPort))))
		if err != nil {
			return fmt.Errorf("failed to connect to target server: %w", err)
		}
	} else {
		// 直接发起 TCP 连接到目标服务器
		targetConn, err = net.Dial("tcp", net.JoinHostPort(p.config.SshHost, strconv.Itoa(int(p.config.SshPort))))
		if err != nil {
			return fmt.Errorf("failed to connect to target server: %w", err)
		}
	}
	defer targetConn.Close()

	// 创建 SSH 客户端
	client, err := p.createSshClient(
		targetConn,
		p.config.SshHost,
		p.config.SshPort,
		p.config.SshAuthMethod,
		p.config.SshUsername,
		p.config.SshPassword,
		p.config.SshKey,
		p.config.SshKeyPassphrase,
	)
	if err != nil {
		return fmt.Errorf("failed to create ssh client: %w", err)
	}
	defer client.Close()

	// 删除质询文件
	challengeFilePath := filepath.Join(p.config.WebRootPath, http01.ChallengePath(token))
	xssh.RemoveRemote(client, challengeFilePath, p.config.UseSCP)

	return nil
}

func (p *provider) createSshClient(conn net.Conn, host string, port int32, authMethod string, username, password, key, keyPassphrase string) (*ssh.Client, error) {
	if host == "" {
		host = "localhost"
	}

	if port == 0 {
		port = 22
	}

	if username == "" {
		username = "root"
	}

	const AUTH_METHOD_NONE = "none"
	const AUTH_METHOD_PASSWORD = "password"
	const AUTH_METHOD_KEY = "key"
	if authMethod == "" {
		if key != "" {
			authMethod = AUTH_METHOD_KEY
		} else if password != "" {
			authMethod = AUTH_METHOD_PASSWORD
		} else {
			authMethod = AUTH_METHOD_NONE
		}
	}

	switch authMethod {
	case AUTH_METHOD_NONE:
		return xssh.NewClient(conn, host, int(port), username)

	case AUTH_METHOD_PASSWORD:
		return xssh.NewClientWithPassword(conn, host, int(port), username, password)

	case AUTH_METHOD_KEY:
		return xssh.NewClientWithKey(conn, host, int(port), username, key, keyPassphrase)

	default:
		return nil, fmt.Errorf("unsupported auth method '%s'", authMethod)
	}
}
