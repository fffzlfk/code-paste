package database

import (
	"code-paste/model"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init(logFile *os.File, host, user, password string) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s",
		host,
		user,
		password,
        "paste", 
	)
	newLogger := logger.New(
		log.New(logFile, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      false,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	db.AutoMigrate(model.Paste{})
}
