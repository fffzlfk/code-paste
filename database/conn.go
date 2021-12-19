package database

import (
	"code-paste/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(host, user, password string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/paste?charset=utf8mb4&parseTime=True&loc=Local", 
		user,
		password,
		host,
	)
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s",
	//	host,
	//	user,
	//	password,
	//	"paste",
	//)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	db.AutoMigrate(model.Paste{})
}
