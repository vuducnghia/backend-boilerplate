package application

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func ConnectDatabase(config *DatabaseConfig) error {
	var logLevel logger.LogLevel
	switch config.LogLevel {
	case "1":
		logLevel = logger.Silent
	case "2":
		logLevel = logger.Error
	case "3":
		logLevel = logger.Warn
	default:
		logLevel = logger.Info
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,    // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN: fmt.Sprintf("user=%s password=%s dbname=%s port=%s",
				config.Username,
				config.Password, config.Database, config.Port),
			PreferSimpleProtocol: false, // enable implicit prepared statement usage
		}), &gorm.Config{
			Logger: newLogger,
		})
	if err != nil {
		return err
	}
	DB = db

	return nil
}
