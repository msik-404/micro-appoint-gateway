package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	ConnString = "amqp://guest:guest@rabbit:5672/"
)

type RabbitConns struct {
	Conn *amqp.Connection
}

func New() (*RabbitConns, error) {
	conns := RabbitConns{}
	conn, err := amqp.Dial(ConnString)
	if err != nil {
		return nil, err
	}
    conns.Conn = conn
    return &conns, nil
}

func (r *RabbitConns) GetRabbitConn() *amqp.Connection {
	return r.Conn
}
