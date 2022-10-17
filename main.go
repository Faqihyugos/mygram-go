package main

import (
	"fmt"
	"log"
	"mygram/auth"
	"mygram/handler"
	"mygram/user"

	"github.com/gin-gonic/gin"
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

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable", host_db, port_db, user_db, pass_db, name_db)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	db.Debug().AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	router.POST("users/register", userHandler.RegisterUser)

	router.Run()

}
