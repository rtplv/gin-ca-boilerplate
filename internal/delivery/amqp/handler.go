package amqp

import (
	"app/internal/config"
	"app/internal/service"
	"app/pkg/amqp"
	"app/pkg/logs"
	"context"
	"fmt"
	"time"
)

type Handler struct {
	ctx context.Context
	config config.RabbitMqConfig
	client amqp.Client
	exampleService service.Example
	logger logs.Logger
	reconnectCount int
}

func NewHandler(ctx context.Context, rmqConfig config.RabbitMqConfig, client amqp.Client, exampleService service.Example,
	logger logs.Logger) *Handler {
	return &Handler{
		ctx: ctx,
		config: rmqConfig,
		client: client,
		exampleService: exampleService,
		logger: logger,
	}
}

func (h *Handler) Consume() error {
	exampleCreateConsumer, err := h.listenExampleCreateQueue()
	//exampleCreateConsumer2, err := h.listenExampleCreateQueue()
	if err != nil {
		return err
	}

	// Reconnect feature
	select {
	// here can add many done channels
	case err = <-exampleCreateConsumer.Done:
		err = h.reconnect()
		if err != nil {
			return err
		}
	}

	return nil
}

func (h Handler) reconnect() error {
	h.reconnectCount += 1
	h.logger.Info(fmt.Sprintf("Reconnecting... Trying № %d", h.reconnectCount))

	err := h.Consume()
	if err != nil {
		if h.reconnectCount == 5 {
			return err
		} else {
			time.Sleep(60)
			return h.reconnect()
		}
	} else {
		// reset reconnect count
		h.reconnectCount = 0
	}
	return nil
}