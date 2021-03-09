package amqpClient

import "github.com/streadway/amqp"

type Feature interface {
	Shutdown() error
	SetDisconnectChannel(ch chan error) *Consumer
}


type Parameters struct {
	ExchangeType string
	RoutingKey string
	PrefetchCount int
	QueueArgs amqp.Table
}