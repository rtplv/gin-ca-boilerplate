package app

import (
	"app/internal/config"
	"app/internal/delivery/http"
	"app/internal/repository"
	"app/internal/service"
	"app/pkg/connections"
	"app/pkg/logs"
	"fmt"
)

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
	//rmq := amqp.NewRabbitClient(cfg.RMQ.Host, cfg.RMQ.Port, cfg.RMQ.User, cfg.RMQ.Password)
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
		//RMQ: rmq,
	})

	rootHandler := http.NewHandler(services.Example, logger)

	router := rootHandler.Init(cfg.GIN.Mode)

	err = router.Run(fmt.Sprintf(":%s", cfg.GIN.Port))
	if err != nil {
		logger.Fatal(err)
	}
}
