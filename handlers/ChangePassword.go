package handlers

import (
	"net/http"
	"strings"

	"github.com/fresh132/authenticationback/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) ChangePassword(c *gin.Context) {
	var input struct {
		Mail        string `json:"mail"`
		OldPassword string `json:"oldpassword"`
		NewPassword string `json:"newpassword"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	if !strings.Contains(input.Mail, "@") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат email"})
		return
	}

	if len(input.NewPassword) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пароль должен содержать минимум 8 символов"})
		return
	}

	var user models.User

	result := h.DB.Where("mail=?", input.Mail).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не найден!"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.OldPassword)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный пароль"})
		return
	}

	if err := h.DB.Where("mail=?", input.NewPassword).First(&user).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Новый пароль уже занят"})
		return
	}

	user.Password = input.NewPassword

	if err := h.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Ошибка при обновлении пароля, попробуйте еще раз"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пароль обновлен успешно!"})

}
