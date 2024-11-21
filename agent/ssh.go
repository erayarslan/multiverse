package agent

import (
	"errors"
	"fmt"
	"io"
	"os"

	goSsh "golang.org/x/crypto/ssh"
)

type ssh struct {
	stdout   io.Writer
	stderr   io.Writer
	stdin    io.Reader
	session  *goSsh.Session
	client   *goSsh.Client
	host     string
	username string
	pemBytes []byte
	port     int
	height   int
	width    int
}

type SSH interface {
	Start() error
	Close() error
}

func (s *ssh) Start() error {
	signer, err := goSsh.ParsePrivateKey(s.pemBytes)
	if err != nil {
		return err
	}

	config := &goSsh.ClientConfig{
		User: s.username,
		Auth: []goSsh.AuthMethod{
			goSsh.PublicKeys(signer),
		},
		HostKeyCallback: goSsh.InsecureIgnoreHostKey(), // nolint:gosec
	}

	s.client, err = goSsh.Dial("tcp", fmt.Sprintf("%s:%d", s.host, s.port), config)
	if err != nil {
		return err
	}

	s.session, err = s.client.NewSession()
	if err != nil {
		return err
	}

	modes := goSsh.TerminalModes{
		goSsh.ECHO:          1,
		goSsh.TTY_OP_ISPEED: 14400,
		goSsh.TTY_OP_OSPEED: 14400,
	}

	term := os.Getenv("TERM")
	if term == "" {
		term = "xterm"
	}

	if err := s.session.RequestPty(term, s.height, s.width, modes); err != nil {
		return err
	}

	s.session.Stdout = s.stdout
	s.session.Stderr = s.stderr
	s.session.Stdin = s.stdin

	if err := s.session.Shell(); err != nil {
		return err
	}

	if err := s.session.Wait(); err != nil {
		var e *goSsh.ExitError
		if errors.As(err, &e) && e.ExitStatus() == 130 {
			return nil
		}

		return err
	}

	return nil
}

func (s *ssh) Close() error {
	err := s.session.Close()
	if err == io.EOF {
		err = nil
	}

	return errors.Join(err, s.client.Close())
}

func NewSSH(host string,
	port int,
	username string,
	pemBytes []byte,
	stdout io.Writer,
	stderr io.Writer,
	stdin io.Reader,
	height int,
	width int,
) SSH {
	return &ssh{
		host:     host,
		port:     port,
		username: username,
		pemBytes: pemBytes,
		stdout:   stdout,
		stderr:   stderr,
		stdin:    stdin,
		height:   height,
		width:    width,
	}
}
