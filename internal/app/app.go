package app

import (
	"app/internal/config"
	"app/internal/delivery/http"
	"app/internal/delivery/rmq"
	"app/internal/repository"
	"app/internal/service"
	amqpPkg "app/pkg/amqp"
	"app/pkg/connections"
	"app/pkg/logs"
	"context"
	"fmt"
)

var ctx = context.Background()

func Run() {
	logger := logs.NewLogger()

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatal(err)
	}

	// Deps

	db, err := connections.NewDatabaseClient(
		cfg.DB.Host,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Database,
		cfg.DB.Port,
	)
	if err != nil {
		logger.Fatal(err)
	}

	// TODO: OPTIONAL
	//es, err := connections.NewElasticSearchClient(cfg.ES.Addresses)
	//if err != nil {
	//	logger.Fatal(err)
	//}

	// TODO: OPTIONAL
	rmqClient := amqpPkg.NewRabbitClient(cfg.RMQ.Host, cfg.RMQ.Port, cfg.RMQ.User, cfg.RMQ.Password)
	rmqConn, err := rmqClient.GetConnection()
	if err != nil {
		logger.Fatal(err)
	}
	defer rmqConn.Close()

	// Services, Repos & API Handlers
	repos := repository.NewRepositories(db)
	services := service.NewServices(service.ServicesDeps{
		Repos: repos,
		Logger: logger,
		// TODO: OPTIONAL
		//DB: db,
		//ES: es,
		//RMQ: rmq,
	})

	stopChan := make(chan bool)

	// HTTP Transport
	go func() {
		rootHandler := http.NewHandler(services.Example, logger)
		router := rootHandler.Init(cfg.GIN.Mode)
		err = router.Run(fmt.Sprintf(":%s", cfg.GIN.Port))
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// AMQP Transport
	go func() {
		amqpContext, cancel := context.WithCancel(ctx)

		rootHandler := rmq.NewHandler(amqpContext, cfg.RMQ, rmqClient, services.Example, logger)
		go rootHandler.Consume()

		cancel()
	}()

	<-stopChan
}
