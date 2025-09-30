package handlers

import (
	"net/http"

	"github.com/fresh132/authenticationback/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Login(c *gin.Context) {
	var input struct {
		Mail     string `json:"mail" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверно набран логин или пароль"})
		return
	}

	var user models.User

	result := h.DB.Where("mail=?", input.Mail).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный email"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный пароль"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Вы успешно вошли в систему"})

}
