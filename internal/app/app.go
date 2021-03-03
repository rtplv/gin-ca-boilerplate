package app

import (
	"app/internal/config"
	"app/internal/delivery/http"
	"app/internal/delivery/rmq"
	"app/internal/repository"
	"app/internal/service"
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

	// Services, Repos & API Handlers
	repos := repository.NewRepositories(db)
	services := service.NewServices(service.ServicesDeps{
		Repos: repos,
		Logger: logger,
		// TODO: OPTIONAL
		//DB: db,
		//ES: es,
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

	// amqpContext instance
	amqpContext, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		rootHandler := rmq.NewHandler(amqpContext, cfg.RMQ, services.Example, logger)
		go rootHandler.Consume()
	}()

	<-stopChan
}
