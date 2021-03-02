package amqp

import (
	"app/internal/service"
	"app/pkg/amqp"
	"app/pkg/logs"
	"context"
)

type Handler struct {
	ctx context.Context
	client amqp.Client
	exampleService service.Example
	logger logs.Logger
}

func NewHandler(ctx context.Context, client amqp.Client, exampleService service.Example, logger logs.Logger) *Handler {
	return &Handler{
		ctx: ctx,
		client: client,
		exampleService: exampleService,
		logger: logger,
	}
}

func (h Handler) Consume() {
	go h.listenExampleCreateQueue()
}