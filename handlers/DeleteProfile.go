package handlers

import (
	"github.com/fresh132/authenticationback/logger"
	"github.com/fresh132/authenticationback/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteProfile(c *gin.Context) {
	userUUID, _ := c.Get("user_id")

	Email, exitsts := c.Get("user_email")

	if !exitsts {
		logger.Warn.Warn("Пользователь не был авторизован",
			"Email", Email,
			"userID", userUUID,
			"Ip", c.ClientIP(),
		)
		c.JSON(401, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	result := h.DB.Unscoped().Delete(&models.User{}, "id = ?", userUUID)

	if result.Error != nil {
		logger.Error.Error("Ошибка при удалении профиля",
			"Email", Email,
			"Ip", c.ClientIP(),
			"Error", result.Error.Error(),
		)
		c.JSON(400, gin.H{"error": "Ошибка при удалении профиля"})
		return
	}

	if result.RowsAffected == 0 {
		logger.Warn.Warn("Попытка удаления несуществующего профиля",
			"Email", Email,
			"Ip", c.ClientIP(),
		)
		c.JSON(404, gin.H{"error": "Профиль не найден"})
		return
	}

	logger.Info.Info("Профиль успешно удален",
		"Email", Email,
		"Ip", c.ClientIP(),
	)

	c.JSON(200, gin.H{"message": "Профиль успешно удален", "dgmail": Email})
}
