package handlers

import (
	"fmt"
	"github.com/fresh132/authenticationback/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) Registred(c *gin.Context) {
	var input struct {
		Mail     string `json:"mail"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
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

	uuidString := uuid.New().String()

	user := models.User{
		ID:       uuidString,
		Mail:     input.Mail,
		Password: string(HashedPassword),
	}

	err = h.DB.Create(&user).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении пользователя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно зарегестрирован", "id": uuidString})
}

func (h *Handler) Entrance(c *gin.Context) {
	var input struct {
		ID       string `json:"id"`
		Mail     string `json:"mail"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	var user models.User

	result := h.DB.Where("id=?", input.ID).First(&user)

	if result.Error != nil {
		fmt.Println("Пользователь не найден")
	}

	if user.Mail != input.Mail {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный логин"})
		return
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный пароль"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Вы успешно вошли в систему"})

}
