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

	url := fmt.Sprintf("amqp://%s:%s@%s:%s",
		cfg.RMQ.User,
		cfg.RMQ.Password,
		cfg.RMQ.Host,
		cfg.RMQ.Port)

	producer, err := amqpClient.NewProducer(url, "default", "go:example-app/example/create", amqpClient.Parameters{})
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
