package main

import (
	"os"

	"github.com/fresh132/authenticationback/authJWT"
	"github.com/fresh132/authenticationback/db"
	"github.com/fresh132/authenticationback/handlers"
	"github.com/fresh132/authenticationback/logger"
	"github.com/fresh132/authenticationback/models"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.InitLogger()

	database := db.InitDB()

	if err := database.AutoMigrate(&models.User{}); err != nil {
		logger.Error.Error("Ошибка миграции: " + err.Error())
		panic("Ошибка миграции: " + err.Error())
	}

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "9091"
	}

	logger.Info.Info("Сервер запущен на порту " + port)

	r.Run(":" + port)
}
