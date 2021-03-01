package repository

import (
	"app/internal/model"
	"context"
	"gorm.io/gorm"
)

type ExampleRepo struct {
	db    *gorm.DB
	table string
}

func NewExampleRepo(db *gorm.DB) *ExampleRepo {
	return &ExampleRepo{
		db:    db,
		table: "public.example",
	}
}

func (r ExampleRepo) Create(ctx context.Context, name string) (model.Example, error) {
	creatingExample := model.Example{
		Name: name,
	}

	err := r.db.
		Table(r.table).
		WithContext(ctx).
		Create(&creatingExample).
		Error
	if err != nil {
		return model.Example{}, nil
	}

	return creatingExample, nil
}

func (r ExampleRepo) GetAll(ctx context.Context) ([]model.Example, error) {
	cities := make([]model.Example, 0)

	err := r.db.
		Table(r.table).
		WithContext(ctx).
		Find(&cities).
		Error

	if err != nil {
		return nil, err
	}

	return cities, nil
}
