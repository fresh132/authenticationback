package main

import (
	"github.com/fresh132/authenticationback/db"
	"github.com/fresh132/authenticationback/handlers"
	"github.com/fresh132/authenticationback/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDB()

	database.AutoMigrate(&models.User{})

	handler := handlers.NewHandler(database)

	r := gin.Default()

	r.POST("/register", handler.Registred)
	r.POST("/enter", handler.Entrance)

	r.Run(":9091")
}
