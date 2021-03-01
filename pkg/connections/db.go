package connections

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func NewDatabaseClient(host, username, password, database, port string) (*gorm.DB, error) {
	connLogger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold: 5 * time.Second,
			LogLevel:      gormLogger.Error,
			Colorful:      true,
		},
	)

	conn, err := openPgConnection(host, username, password, database, port, &gorm.Config{
		Logger: connLogger,
	})

	if err != nil {
		return nil, err
	}


	sqlDB, err := conn.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(300)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return conn, nil
}

func openPgConnection(host string, username string, password string, database string, port string, config *gorm.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		host,
		username,
		password,
		database,
		port,
	)

	return gorm.Open(postgres.Open(dsn), config)
}
