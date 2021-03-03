package rmq

import (
	amqpPkg "app/pkg/amqp"
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"time"
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
			go h.processExampleCreateMessage(d)
		})

	return consumer, err
}

type ExampleMessage struct {
	Name string `json:"name"`
}

func (h Handler) processExampleCreateMessage(rawMessage amqp.Delivery) {
	ctx, cancel := context.WithTimeout(h.ctx, 30 * time.Minute)
	defer cancel()

	var exampleMsg ExampleMessage
	err := json.Unmarshal(rawMessage.Body, &exampleMsg)
	if err != nil {
		h.logger.Error(err)
		return
	}

	createdExample, err := h.exampleService.Create(ctx, exampleMsg.Name)
	if err != nil {
		h.logger.Error(err)
		return
	}

	fmt.Println(createdExample)

	if err = rawMessage.Ack(false); err != nil {
		h.logger.Error(err)
		return
	}
}
