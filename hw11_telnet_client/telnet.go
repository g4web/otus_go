package main

import (
	"io"
	"net"
	"time"
)

const network = "tcp"

type TelnetHandler interface {
	Connect() error
	Close() error
	Send() error
	Receive() error
}

type TelnetClient struct {
	address string
	timeout time.Duration
	in      io.Reader
	out     io.Writer
	conn    net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetHandler {
	return &TelnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (t *TelnetClient) Connect() error {
	conn, err := net.DialTimeout(network, t.address, t.timeout)
	if err != nil {
		return err
	}

	t.conn = conn

	return nil
}

func (t *TelnetClient) Close() error {
	if t.conn != nil {
		return t.conn.Close()
	}

	return nil
}

func (t *TelnetClient) Receive() error {
	return t.handleMessage(t.conn, t.out)
}

func (t *TelnetClient) Send() error {
	return t.handleMessage(t.in, t.conn)
}

func (t *TelnetClient) handleMessage(from io.Reader, to io.Writer) error {
	if _, err := io.Copy(to, from); err != nil {
		return err
	}
	return nil
}
