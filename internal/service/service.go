package service

import (
	"app/internal/model"
	"app/internal/repository"
	"app/pkg/logs"
	"context"
	//es7 "github.com/elastic/go-elasticsearch/v7"
	//"gorm.io/gorm"
)

type Example interface {
	Create(ctx context.Context, name string) (model.Example, error)
	GetAll(ctx context.Context) ([]model.Example, error)
}

type Services struct {
	Example
}

type ServicesDeps struct {
	Repos *repository.Repositories
	Logger logs.Logger
	// Connections
	//DB *gorm.DB
	//ES *es7.Client
}

func NewServices(deps ServicesDeps) *Services {
	exampleService := NewExampleService(deps.Repos.Example, deps.Logger)

	return &Services{
		Example: exampleService,
	}
}