package client

import (
	"io"
	"net"
	"errors"
)

type Message interface {

}

type MessageHandler interface {
	Handle(message Message)
}

type Qos byte

const (
	AtMostOnce Qos = 0
	AtLeastOnce = 1
	ExactlyOnce = 2
)

type Connection interface {
	Connect() error
	Disconnect() error
	Unsubscribe(topics ...string) error
	Subscribe(handler MessageHandler, topics ...string) error
	Publish(topic string, qos Qos, retained bool, payload []byte) error
}

type dummy struct {
	connected bool
	conn net.Conn
}

func (d *dummy) Connect() error {
	if d.connected {
		return errors.New("Already connected")
	}
	// open socket

	// send CONNECT packet

	// wait for CONNACK packet

	d.connected = true
	return nil
}

func (d *dummy) Disconnect() error {
	if !d.connected {
		return errors.New("Not connected")
	}

	err := disconnect(d.conn)
	err = d.conn.Close()

	d.connected = false
	return err
}

func disconnect(w io.Writer) error {
	_, err := w.Write([]byte{})

	return err
}