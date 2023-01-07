package db

import (
	"fmt"
	"os"

	"todo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
  
func ConnectToDataBase()(db *gorm.DB) {
	var err error
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Status:", err)
	} else {
		fmt.Println("Database connection successfully opened")
	}
	return DB
}

func SyncDB() {
	DB.AutoMigrate(&models.Todo{})
}