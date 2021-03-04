package repository

import (
	"app/internal/model"
	"context"
	"gorm.io/gorm"
)

type ExampleRepo struct {
	db    *gorm.DB
}

func NewExampleRepo(db *gorm.DB) *ExampleRepo {
	return &ExampleRepo{
		db:    db,
	}
}

func (r ExampleRepo) Create(ctx context.Context, name string) (model.Example, error) {
	creatingExample := model.Example{
		Name: name,
	}

	err := r.db.
		WithContext(ctx).
		Create(&creatingExample).
		Error
	if err != nil {
		return model.Example{}, nil
	}

	return creatingExample, nil
}

func (r ExampleRepo) GetAll(ctx context.Context) ([]model.Example, error) {
	examples := make([]model.Example, 0)

	err := r.db.
		WithContext(ctx).
		Find(&examples).
		Error

	if err != nil {
		return nil, err
	}

	return examples, nil
}
