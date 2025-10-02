package main

import (
	"github.com/fresh132/authenticationback/authJWT"
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
	r.POST("/enter", handler.Login)

	auth := r.Group("/auth")
	auth.Use(authjwt.AuthMiddleware())
	{
		auth.GET("/user", handler.GetProfile)
		auth.DELETE("/delete", handler.DeleteProfile)
		auth.PUT("/update", handler.ChangePassword)
	}
	r.Run(":9091")
}
