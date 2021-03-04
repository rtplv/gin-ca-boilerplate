package main

import (
	"app/internal/config"
	"app/internal/model"
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

	err = db.Migrator().AutoMigrate(
		&model.Example{},
	)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Migrate successful")
}
