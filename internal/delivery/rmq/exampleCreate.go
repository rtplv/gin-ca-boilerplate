package rmq

import (
	amqpPkg "app/pkg/amqp"
	"app/pkg/logs"
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

const queueName = "go:example-app/example/create"

func (h Handler) listenExampleCreateQueue(errCh chan error) (*amqpPkg.Consumer, error) {
	consumer, err := amqpPkg.NewConsumer(h.config, "default", queueName, "main", amqpPkg.Parameters{
		PrefetchCount: 10,
	})
	if err != nil {
		return nil, err
	}

	h.logger.Info("example/create consumer connection established")

	go consumer.
		SetDisconnectChannel(errCh).
		Handle(func(d amqp.Delivery) {
		fmt.Println(d)
		if err := d.Ack(false); err != nil {
			// TODO: here need failed message logging
			d.Reject(false)
		}
	})

	return consumer, err
}

type ExampleMessage struct {
	Name string `json:"name"`
}

func (h Handler) processExampleCreateMessage(ctx context.Context, logger logs.Logger, rawMessage amqp.Delivery) {
	var exampleMsg ExampleMessage
	err := json.Unmarshal(rawMessage.Body, &exampleMsg)

	createdExample, err := h.exampleService.Create(ctx, exampleMsg.Name)
	if err != nil {
		h.logger.Error(err)
		return
	}

	fmt.Println(createdExample)

	if err = rawMessage.Ack(false); err != nil {
		logger.Error(err)
		return
	}
}
