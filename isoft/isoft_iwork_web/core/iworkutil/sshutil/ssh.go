package sshutil

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
)

func SSHConnectTest(user, password, host string, port int) error {
	client, session, err := SSHConnect(user, password, host, port)
	if err != nil {
		return err
	}
	defer client.Close()
	defer session.Close()
	return nil
}

func SSHConnect(user, password, host string, port int) (*ssh.Client, *ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: auth,
		// Timeout:             30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}
	// connet to sshutil
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, nil, err
	}
	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, nil, err
	}
	return client, session, nil
}
