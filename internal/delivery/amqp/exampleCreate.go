package amqp

import (
	"app/internal/enum"
	"app/pkg/logs"
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

func (h Handler) listenExampleCreateQueue() {
	taskCreateChan, taskCreateAmqpChan, err  := h.client.GetConsumeChan(enum.ExampleCreateQueue, 10)
	if err != nil {
		h.logger.Fatal(err)
	}
	defer taskCreateAmqpChan.Close()

	// TODO: здесь нужно отменять горутину и закрывать канал, если контекст отменен

	for msg := range taskCreateChan {
		// 30 minutes timeout limit
		msgCtx, _ := context.WithDeadline(h.ctx, time.Now().Add(30 * time.Minute))
		go h.processExampleCreateMessage(msgCtx, h.logger, msg)
	}
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
