package amqp

import (
	"app/internal/config"
	"fmt"
	"github.com/streadway/amqp"
)

type Consumer struct {
	Conn       *amqp.Connection
	Channel    *amqp.Channel
	Tag        string
	Deliveries <-chan amqp.Delivery
	Done       chan error
}

type Parameters struct {
	ExchangeType string
	RoutingKey string
	PrefetchCount int
}

func NewConsumer(config config.RabbitMqConfig, exchange, queueName string, ctag string,
	parameters Parameters) (*Consumer, error) {
	c := &Consumer{
		Conn:    nil,
		Channel: nil,
		Tag:     ctag,
		Done:    make(chan error),
	}

	// Optional parameters
	exchangeType := "direct"
	routingKey := parameters.RoutingKey
	prefetchCount := 1

	if parameters.ExchangeType != "" {
		exchangeType = parameters.ExchangeType
	}

	if parameters.PrefetchCount > 0 {
		prefetchCount = parameters.PrefetchCount
	}

	var err error

	// Open connection
	c.Conn, err = amqp.Dial(
		fmt.Sprintf("amqp://%s:%s@%s:%s",
			config.User,
			config.Password,
			config.Host,
			config.Port,
		))
	if err != nil {
		return nil, err
	}

	// Close connection error throwing
	go func() {
		err := <-c.Conn.NotifyClose(make(chan *amqp.Error))
		fmt.Println(err)
		c.Done<-err
	}()

	c.Channel, err = c.Conn.Channel()
	if err != nil {
		return nil, err
	}

	err = c.Channel.Qos(prefetchCount, 0, false)
	if err != nil {
		return nil, err
	}

	if err = c.Channel.ExchangeDeclare(
		exchange,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, err
	}

	queue, err := c.Channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	if err = c.Channel.QueueBind(
		queue.Name,
		routingKey,
		exchange,
		false,
		nil,
	); err != nil {
		return nil, err
	}

	c.Deliveries, err = c.Channel.Consume(
		queue.Name,
		c.Tag,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Consumer) Handle(callback func(d amqp.Delivery)) {
	for d := range c.Deliveries {
		callback(d)
	}

	c.Done <- nil
}

func (c *Consumer) Shutdown() error {
	// will close() the deliveries channel
	if err := c.Channel.Cancel(c.Tag, true); err != nil {
		return err
	}

	if err := c.Conn.Close(); err != nil {
		return err
	}

	// wait for handle() to exit
	return <-c.Done
}
