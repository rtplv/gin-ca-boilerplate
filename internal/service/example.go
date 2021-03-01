package service

import (
	"app/internal/model"
	"app/internal/repository"
	"app/pkg/logs"
	"context"
)

type ExampleService struct {
	repo repository.Example
	logger *logs.Logger
}

func NewExampleService(repo repository.Example, logger *logs.Logger) *ExampleService {
	return &ExampleService{
		repo,
		logger,
	}
}

func (s ExampleService) Create(ctx context.Context, name string) (model.Example, error) {
	return s.repo.Create(ctx, name)
}

func (s ExampleService) GetAll(ctx context.Context) ([]model.Example, error) {
	return s.repo.GetAll(ctx)
}
