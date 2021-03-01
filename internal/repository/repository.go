package repository

import (
	"app/internal/model"
	"context"
	"gorm.io/gorm"
)

type Example interface {
	Create(ctx context.Context, name string) (model.Example, error)
	GetAll(ctx context.Context) ([]model.Example, error)
}


type Repositories struct {
	Example Example
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Example: NewExampleRepo(db),
	}
}
