package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func GetDbConn() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(GoDotEnvVariable("dsn")), &gorm.Config{})

	if err != nil {

		return nil, err
	}

	return db, nil

}
