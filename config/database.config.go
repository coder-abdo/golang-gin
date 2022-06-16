package config

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	envErr := godotenv.Load()
	if envErr != nil {
		panic("couldn't load env file")
	}
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp($s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("faild to connect database")
	}
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	mySqlDB, err := db.DB()

	if err != nil {
		panic("faild to close database connections")
	}
	mySqlDB.Close()
}
