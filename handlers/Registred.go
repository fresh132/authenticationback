package handlers

import (
	"net/http"
	"time"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверно набран логин или пароль"})
		return
	}

	var users models.User

	result := h.DB.Where("mail=?", input.Mail).First(&users)

	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь с таким mail уже сувщевствует"})
		return
	}

	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хешировании пароля"})
		return
	}

	create := time.Now()
	uuid, _ := uuid.NewRandom()

	user := models.User{
		ID:         uuid,
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
