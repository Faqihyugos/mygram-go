package config

import (
	"fmt"
	"log"

	"github.com/Faqihyugos/mygram-go/comment"
	"github.com/Faqihyugos/mygram-go/photo"
	"github.com/Faqihyugos/mygram-go/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host_db = "localhost"
	port_db = 5432
	user_db = "postgres"
	name_db = "mygram"
	pass_db = "postgres"
)

func StartDB() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable", host_db, port_db, user_db, pass_db, name_db)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	db.Debug().AutoMigrate(&user.User{}, &photo.Photo{}, &comment.Comment{})

	if err != nil {
		log.Fatal(err.Error())
	}
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}
	defer fmt.Println("Successfully Connected to Database")
	return db
}
