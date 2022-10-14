package theoryssh

import (
	"bytes"
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
)

func SSHMaker() {
	var stdout, stderr bytes.Buffer
	session, err := SSHConnet("name", "pwd", "ip", "port")
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()
	session.Stderr = &stderr
	session.Stdout = &stdout

	session.Run("pwd")

	fmt.Println(stdout.String())
}

// SSHConnet vi /etc/ssh/sshd_config
// PasswordAuthentication yes
func SSHConnet(user, pwd, host, port string) (*ssh.Session, error) {
	auth := []ssh.AuthMethod{
		ssh.Password(pwd),
	}
	hostKeyCallback := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	clientConfig := &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		HostKeyCallback: hostKeyCallback,
	}
	var client *ssh.Client
	var err error
	if client, err = ssh.Dial("tcp", fmt.Sprintf("%s:%s", host, port), clientConfig); err != nil {
		return nil, err
	}
	var session *ssh.Session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}
	return session, nil
}
