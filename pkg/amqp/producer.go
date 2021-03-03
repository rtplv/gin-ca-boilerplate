package amqp

import (
	"app/internal/config"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)

type Producer struct {
	Config     config.RabbitMqConfig
	Conn       *amqp.Connection
	Channel    *amqp.Channel
	Queue 	   amqp.Queue
	Tag        string
	Deliveries <-chan amqp.Delivery
	Disconnect chan error
}

func NewProducer(config config.RabbitMqConfig, exchange, queueName string, ctag string,
	parameters Parameters) (*Producer, error) {
	p := &Producer{
		Config:  config,
		Tag:     ctag,
		Disconnect: make(chan error),
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
	p.Conn, err = amqp.Dial(
		fmt.Sprintf("amqp://%s:%s@%s:%s",
			p.Config.User,
			p.Config.Password,
			p.Config.Host,
			p.Config.Port,
		))
	if err != nil {
		return nil, err
	}

	// Close connection error throwing
	go func() {
		err := <-p.Conn.NotifyClose(make(chan *amqp.Error))

		if err != nil {
			p.Disconnect<-errors.New(err.Error())
		}
	}()

	p.Channel, err = p.Conn.Channel()
	if err != nil {
		return nil, err
	}

	err = p.Channel.Qos(prefetchCount, 0, false)
	if err != nil {
		return nil, err
	}

	if err = p.Channel.ExchangeDeclare(
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

	p.Queue, err = p.Channel.QueueDeclare(
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

	if err = p.Channel.QueueBind(
		p.Queue.Name,
		routingKey,
		exchange,
		false,
		nil,
	); err != nil {
		return nil, err
	}

	p.Deliveries, err = p.Channel.Consume(
		p.Queue.Name,
		p.Tag,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Producer) Send(publishing amqp.Publishing) (err error) {
	err = p.Channel.Publish("", p.Queue.Name, false, false, publishing)

	return
}

func (p *Producer) Shutdown() error {
	// will close() the deliveries channel
	if err := p.Channel.Cancel(p.Tag, true); err != nil {
		return err
	}

	if err := p.Conn.Close(); err != nil {
		return err
	}

	// wait for handle() to exit
	return nil
}

func (p *Producer) SetDisconnectChannel(ch chan error) *Producer {
	p.Disconnect = ch

	return p
}
