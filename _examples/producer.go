package main

import (
	"app/internal/config"
	"app/pkg/amqpClient"
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	producer, err := amqpClient.NewProducer(cfg.RMQ, "default", "go:example-app/example/create", "default", amqpClient.Parameters{})
	if err != nil {
		fmt.Println(err)
	}
	defer producer.Shutdown()

	err = producer.Send(amqp.Publishing{
		ContentType:     "application/json",
		DeliveryMode:    amqp.Transient,
		Body:            []byte(`{ "name": "testName" }`),
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Message sended success")
}
