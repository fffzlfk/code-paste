package database

import (
	"code-paste/model"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init(logFile *os.File, host, user, password string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/paste?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
	)
	newLogger := logger.New(
		log.New(logFile, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	db.AutoMigrate(model.Paste{})
}
