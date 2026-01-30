package ssh

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
)

func NewClient(conn net.Conn, host string, port int, username string) (*ssh.Client, error) {
	authentications := make([]ssh.AuthMethod, 0)
	return newClientWithAuthMethods(conn, host, port, username, authentications)
}

func NewClientWithPassword(conn net.Conn, host string, port int, username string, password string) (*ssh.Client, error) {
	authentications := make([]ssh.AuthMethod, 0)
	authentications = append(authentications, ssh.Password(password))
	authentications = append(authentications, ssh.KeyboardInteractive(func(user, instruction string, questions []string, echos []bool) ([]string, error) {
		if len(questions) == 1 {
			return []string{password}, nil
		}
		return nil, fmt.Errorf("unexpected keyboard interactive question [%s]", strings.Join(questions, ", "))
	}))
	return newClientWithAuthMethods(conn, host, port, username, authentications)
}

func NewClientWithKey(conn net.Conn, host string, port int, username string, key, keyPassphrase string) (*ssh.Client, error) {
	var signer ssh.Signer
	var err error
	if keyPassphrase != "" {
		signer, err = ssh.ParsePrivateKeyWithPassphrase([]byte(key), []byte(keyPassphrase))
	} else {
		signer, err = ssh.ParsePrivateKey([]byte(key))
	}
	if err != nil {
		return nil, err
	}

	authentications := make([]ssh.AuthMethod, 0)
	authentications = append(authentications, ssh.PublicKeys(signer))
	return newClientWithAuthMethods(conn, host, port, username, authentications)
}

func newClientWithAuthMethods(conn net.Conn, host string, port int, username string, authMethods []ssh.AuthMethod) (*ssh.Client, error) {
	addr := net.JoinHostPort(host, strconv.Itoa(int(port)))

	sshConn, chans, reqs, err := ssh.NewClientConn(conn, addr, &ssh.ClientConfig{
		User:            username,
		Auth:            authMethods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		return nil, err
	}

	return ssh.NewClient(sshConn, chans, reqs), nil
}
