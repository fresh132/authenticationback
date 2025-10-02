package handlers

import (
	"net/http"

	"github.com/fresh132/authenticationback/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProfile(c *gin.Context) {
	userID, exitsts := c.Get("user_id")

	if !exitsts {
		c.JSON(401, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	var user models.User

	resilt := h.DB.Where("id=?", userID).First(&user)

	if resilt.Error != nil {
		c.JSON(404, gin.H{"error": "Пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":         user.ID,
			"mail":       user.Mail,
			"created_at": user.CreatedAt,
		},
	})

}
