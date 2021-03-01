package main

import (
	"app/internal/config"
	"app/internal/enum"
	amqpPkg "app/pkg/amqp"
	"app/pkg/logs"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"os"
	"time"
)

var ctx = context.Background()

func main()  {
	logger := logs.NewLogger()

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatal(err)
	}


	// TODO: Optional
	//db, err := connections.NewDatabaseClient(cfg.DB.Host, cfg.DB.Username, cfg.DB.Password, cfg.DB.Database, cfg.DB.Port)
	//if err != nil {
	//	logger.Fatalln(err)
	//}

	rmq := amqpPkg.NewRabbitClient(cfg.RMQ.Host, cfg.RMQ.Port, cfg.RMQ.User, cfg.RMQ.Password)
	rmqConn, err := rmq.GetConnection()
	if err != nil {
		logger.Fatal(err)
	}

	// TODO: Optional
	// Init app features
	//repos := repository.NewRepositories(db)
	//services := service.NewServices(service.ServicesDeps{
	//	Repos: repos,
	//  Logger: logger,
	//	//DB: db,
	//	// TODO: Optional
	//	//RMQ: rmq,
	//})

	// rmq channel
	messageChannel, rmqChannel, err := rmq.GetConsumeChan(enum.ExampleQueue, 10)
	if err != nil {
		logger.Fatal(err)
	}
	defer rmqChannel.Close()
	defer rmqConn.Close()


	logger.Info(fmt.Sprintf("Consumer read, PID: %d", os.Getpid()))

	stopChan := make(chan bool)

	// listen channel
	go func() {
		for m := range messageChannel {
			// Create context with deadline, for cancelling async reqs
			msgCtx, cancel := context.WithDeadline(ctx, time.Now().Add(30 * time.Minute))

			go processMessage(msgCtx, logger, m)

			cancel()
		}
	}()

	// Check connection status
	go func() {
		ticker := time.Tick(30 * time.Second)

		for now := range ticker {
			if rmqConn.IsClosed() {
				logger.Fatal(errors.New(fmt.Sprintf("Connection refused. %d", now.Unix())))
			}
		}
	}()

	<-stopChan
}

type ExampleMessage struct {
	Text string `json:"text"`
}

func processMessage(ctx context.Context, logger *logs.Logger, rawMessage amqp.Delivery) {
	var report ExampleMessage
	err := json.Unmarshal(rawMessage.Body, &report)

	if err = rawMessage.Ack(false); err != nil {
		logger.Error(err)
		return
	}
}
