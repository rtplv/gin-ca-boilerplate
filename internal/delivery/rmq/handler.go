package rmq

import (
	"app/internal/config"
	"app/internal/service"
	"app/pkg/amqp"
	"app/pkg/logs"
	"context"
	"errors"
	"fmt"
	"time"
)

type Handler struct {
	ctx context.Context
	config config.RabbitMqConfig
	client amqp.Client
	exampleService service.Example
	logger logs.Logger
	reconnectAttempts int
	maxReconnectAttempts int
}

func NewHandler(ctx context.Context, rmqConfig config.RabbitMqConfig, client amqp.Client, exampleService service.Example,
	logger logs.Logger) *Handler {
	return &Handler{
		ctx: ctx,
		config: rmqConfig,
		client: client,
		exampleService: exampleService,
		logger: logger,
		maxReconnectAttempts: 10,
	}
}

func (h *Handler) Consume() {
	errCh := make(chan error)

	_, err := h.listenExampleCreateQueue(errCh)
	_, err = h.listenExampleCreateQueue(errCh)
	_, err = h.listenExampleCreateQueue(errCh)
	_, err = h.listenExampleCreateQueue(errCh)
	_, err = h.listenExampleCreateQueue(errCh)
	if err != nil {
		h.logger.Error(err)
		go h.reconnect()
		return
	}

	// Reconnect in case if consumer down. If attempt execute error, reconnect attempts count not reset
	select {
	// here can add many disconnect channels
	case err = <-errCh:
		h.reconnectAttempts = 0
		h.logger.Error(err)
		go h.reconnect()
	}
}

func (h *Handler) reconnect() {
	if h.reconnectAttempts == h.maxReconnectAttempts {
		h.logger.Error(errors.New("reconnect failed. Max attempts count achieved"))
		return
	}

	time.Sleep(1 * time.Second)

	h.reconnectAttempts += 1
	h.logger.Info(fmt.Sprintf("Reconnecting... Attempt № %d", h.reconnectAttempts))

	go h.Consume()
}