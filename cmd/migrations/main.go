package main

import (
	"app/internal/config"
	"app/pkg/connections"
	"app/pkg/logs"
)

func main() {
	logger := logs.NewLogger()
	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatal(err)
	}

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

	NewExampleMigration(db, logger).Up()
}
