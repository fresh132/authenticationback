package handlers

import (
	"net/http"
	"time"

	"github.com/fresh132/authenticationback/logger"
	"github.com/fresh132/authenticationback/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Registred(c *gin.Context) {
	var input struct {
		Mail     string `json:"mail" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Warn.Warn("Неверный ввод данных при регистрации",
			"Email", input.Mail,
			"Ip", c.ClientIP(),
			"Error", err.Error(),
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверно набран логин или пароль"})
		return
	}

	var users models.User

	result := h.DB.Where("mail=?", input.Mail).First(&users)

	if result.Error == nil {
		logger.Warn.Warn("Попытка регистрации с уже существующим mail",
			"Email", input.Mail,
			"Ip", c.ClientIP(),
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь с таким mail уже существует"})
		return
	}

	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		logger.Error.Error("Ошибка при хешировании пароля",
			"Email", input.Mail,
			"Ip", c.ClientIP(),
			"Error", err.Error(),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хешировании пароля"})
		return
	}

	nowtime := time.Now()
	uuid, _ := uuid.NewRandom()

	user := models.User{
		ID:         uuid,
		Mail:       input.Mail,
		Password:   string(HashedPassword),
		CreateTime: nowtime,
	}

	err = h.DB.Create(&user).Error

	if err != nil {
		logger.Error.Error("Ошибка при сохранении пользователя",
			"Email", input.Mail,
			"Ip", c.ClientIP(),
			"Error", err.Error(),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении пользователя"})
		return
	}

	logger.Info.Info("Пользователь успешно зарегистрирован",
		"Email", input.Mail,
		"Ip", c.ClientIP(),
	)

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно зарегистрирован"})
}
