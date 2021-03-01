package config

import (
	"github.com/joho/godotenv"
	"os"
)

type (
	Config struct {
		GIN GinConfig
		DB  DatabaseConfig
		ES  ElasticConfig
		RMQ RabbitMqConfig
	}

	GinConfig struct {
		Mode string
		Port string
	}

	DatabaseConfig struct {
		Host     string
		Port     string
		Database string
		Username string
		Password string
	}

	ElasticConfig struct {
		Addresses []string
	}

	RabbitMqConfig struct {
		Host     string
		Port     string
		User     string
		Password string
	}
)

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		GIN: GinConfig{
			Mode: os.Getenv("GIN_MODE"),
			Port: os.Getenv("GIN_PORT"),
		},
		DB: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_DATABASE"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
		},
		ES: ElasticConfig{
			Addresses: []string{
				os.Getenv("ES_URL"),
			},
		},
		RMQ: RabbitMqConfig{
			Host:     os.Getenv("RABBITMQ_HOST"),
			Port:     os.Getenv("RABBITMQ_PORT"),
			User:     os.Getenv("RABBITMQ_USER"),
			Password: os.Getenv("RABBITMQ_PASSWORD"),
		},
	}, nil
}
