package handlers

import (
	"net/http"

	"github.com/fresh132/authenticationback/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) ChangePassword(c *gin.Context) {

	Email, exitsts := c.Get("user_email")

	if !exitsts {
		c.JSON(401, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	var input struct {
		NewPassword string `json:"newpassword" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверно набран логин или пароль"})
		return
	}

	var user models.User

	result := h.DB.Where("mail=?", Email).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не найден"})
		return
	}

	hachPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при хешировании пароля"})
		return
	}

	user.Password = string(hachPassword)

	if err := h.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Ошибка при обновлении пароля, попробуйте еще раз"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пароль обновлен успешно!"})

}
