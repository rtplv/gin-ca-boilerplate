package amqp

import (
	"fmt"
	"github.com/streadway/amqp"
)

type RabbitClient struct {
	Host     string
	Port     string
	User     string
	password string
}

func NewRabbitClient(host, port, user, password string) *RabbitClient {
	return &RabbitClient{
		host,
		port,
		user,
		password,
	}
}

func (c RabbitClient) GetConnection() (*amqp.Connection, error) {
	return amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", c.User, c.password, c.Host, c.Port))
}

func (c RabbitClient) Send(queueName string, msg []byte) error {
	// Init connection
	conn, err := c.GetConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Init channel
	ch, err := conn.Channel()
	defer ch.Close()

	if err != nil {
		return err
	}

	// Declare queue
	queue, err := ch.QueueDeclare(queueName, true, false, false, false, nil)

	if err != nil {
		return err
	}

	err = ch.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Transient,
		ContentType:  "application/json",
		Body:         msg,
	})

	if err != nil {
		return err
	}

	return nil
}

func (c RabbitClient) GetConsumeChan(queueName string, prefetchCount int) (<-chan amqp.Delivery, *amqp.Channel, error) {
	conn, err := c.GetConnection()
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	// Declare queue
	queue, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return nil, nil, err
	}

	// Set QoS
	err = ch.Qos(prefetchCount, 0, false)
	if err != nil {
		return nil, nil, err
	}

	// Create message channel
	messageChannel, err := ch.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return nil, nil, err
	}

	return messageChannel, ch, nil
}
