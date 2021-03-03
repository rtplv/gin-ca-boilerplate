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

	credentials := amqpClient.Credentials{
		User:     cfg.RMQ.User,
		Password: cfg.RMQ.Password,
		Host:     cfg.RMQ.Host,
		Port:     cfg.RMQ.Port,
	}

	producer, err := amqpClient.NewProducer(credentials, "default", "go:example-app/example/create", "default", amqpClient.Parameters{})
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
