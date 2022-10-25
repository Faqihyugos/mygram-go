package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Faqihyugos/mygram-go/comment"
	"github.com/Faqihyugos/mygram-go/photo"
	"github.com/Faqihyugos/mygram-go/sosmed"
	"github.com/Faqihyugos/mygram-go/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetEnvWithKey(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}

func StartDB() *gorm.DB {
	get := GetEnvWithKey
	DB_USER := get("DB_USER")
	DB_PASS := get("DB_PASS")
	DB_HOST := get("DB_HOST")
	DB_NAME := get("DB_NAME")
	DB_PORT := get("DB_PORT")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s  dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	db.Debug().AutoMigrate(&user.User{}, &photo.Photo{}, &comment.Comment{}, &sosmed.Sosmed{})

	if err != nil {
		log.Fatal(err.Error())
	}
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}
	defer fmt.Println("Successfully Connected to Database")
	return db
}
