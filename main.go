package main

import (
	"os"

	authjwt "github.com/fresh132/authenticationback/authJWT"
	"github.com/fresh132/authenticationback/db"
	_ "github.com/fresh132/authenticationback/docs"
	"github.com/fresh132/authenticationback/handlers"
	"github.com/fresh132/authenticationback/logger"
	"github.com/fresh132/authenticationback/models"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Authentication Service
// @version 1.0
// @description Сервис аутентификации пользователей с использованием JWT.
// @host localhost:9091
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "9091"
	}

	logger.Info.Info("Сервер запущен на порту " + port)

	r.Run(":" + port)
}
