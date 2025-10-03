package handlers

import (
	"net/http"

	"github.com/fresh132/authenticationback/logger"
	"github.com/fresh132/authenticationback/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProfile(c *gin.Context) {
	userUUID, userIDExists := c.Get("user_id")
	Email, emailExists := c.Get("user_email")

	if !userIDExists || !emailExists {
		logger.Warn.Warn("Пользователь не был авторизован",
			"Email", Email,
			"userID", userUUID,
			"Ip", c.ClientIP(),
		)
		c.JSON(401, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	var user models.User

	result := h.DB.Where("id=?", userUUID).First(&user)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Пользователь не найден"})
		return
	}

	logger.Info.Info("Профиль успешно получен",
		"Email", Email,
		"userID", userUUID,
		"Ip", c.ClientIP(),
	)

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":         user.ID,
			"mail":       user.Mail,
			"created_at": user.CreatedAt,
		},
	})

}
