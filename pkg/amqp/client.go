package amqp

import "github.com/streadway/amqp"

type Client interface {
	GetConnection() (*amqp.Connection, error)

	Send(queueName string, msg []byte) error
	GetConsumeChan(queueName string, prefetchCount int) (<-chan amqp.Delivery, *amqp.Channel, error)
}
