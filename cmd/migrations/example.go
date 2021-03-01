package main

import (
	"app/internal/model"
	"app/pkg/logs"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

const tableName = "public.example"

type ExampleMigration struct {
	DB *gorm.DB
	Logger *logs.Logger
}

func NewExampleMigration(db *gorm.DB, logger *logs.Logger) *ExampleMigration {
	return &ExampleMigration{
		DB: db,
		Logger: logger,
	}
}

func (m ExampleMigration) Up()  {
	targetModel := model.Example{}
	tableExist := m.DB.Table(tableName).Migrator().HasTable(targetModel)

	if tableExist {
		m.Logger.Fatal(errors.New(fmt.Sprintf("Table %s already exist", tableName)))
	}

	err := m.DB.Table(tableName).Migrator().AutoMigrate(&targetModel)
	if err != nil {
		m.Logger.Fatal(err)
	}

	m.Logger.Info("Successfully migrate!")
}

func (m ExampleMigration) Down()  {
	targetModel := model.Example{}
	tableExist := m.DB.Table(tableName).Migrator().HasTable(targetModel)

	if !tableExist {
		m.Logger.Fatal(errors.New(fmt.Sprintf("Table %s does not exist", tableName)))
	}

	err := m.DB.Table(tableName).Migrator().DropTable(&targetModel)
	if err != nil {
		m.Logger.Fatal(err)
	}

	m.Logger.Info("Successfully migrate down!")
}
