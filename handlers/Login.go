package handlers

import (
	"fmt"
	"net/http"

	"github.com/fresh132/authenticationback/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Login(c *gin.Context) {
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

	result := h.DB.Where("mail=?", input.Mail).First(&user)

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
