package handlers

import (
	"github.com/fresh132/authenticationback/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteProfile(c *gin.Context) {
	Email, exitsts := c.Get("user_email")

	if !exitsts {
		c.JSON(401, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	var user models.User

	result := h.DB.Where("mail = ?", Email).Delete(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Ошибка при удалении профиля"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Профиль не найден"})
		return
	}

	c.JSON(200, gin.H{"message": "Профиль успешно удален", "dgmail": Email})
}
