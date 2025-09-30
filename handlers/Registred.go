package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/fresh132/authenticationback/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Registred(c *gin.Context) {
	var input struct {
		Mail     string `json:"mail"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	var users models.User

	result := h.DB.Where("mail=?", input.Mail).First(&users)

	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь с таким mail уже сувщевствует"})
		return
	}

	if !strings.Contains(input.Mail, "@") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат email"})
		return
	}

	if len(input.Password) < 8 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Пароль должен содержать минимум 8 символов"})
		return
	}

	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хешировании пароля"})
		return
	}

	create := time.Now()

	user := models.User{
		Mail:       input.Mail,
		Password:   string(HashedPassword),
		CreateTime: int(time.Since(create)),
	}

	err = h.DB.Create(&user).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении пользователя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно зарегестрирован"})
}
