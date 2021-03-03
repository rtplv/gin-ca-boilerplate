package main

import (
	"app/internal/config"
	amqpPkg "app/pkg/amqp"
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	producer, err := amqpPkg.NewProducer(cfg.RMQ, "default", "go:example-app/example/create", "default", amqpPkg.Parameters{})
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
