package handlers

import (
	"net/http"

	"github.com/fresh132/authenticationback/logger"
	"github.com/fresh132/authenticationback/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Create ChangePassword
// @Security ApiKeyAuth
// @Description Позволяет авторизованному пользователю сменить пароль
// @Tags user
// @ID ChangePassword
// @Accept  json
// @Produce json
// @Param   input  body  models.PasswordChangeRequest  true  "Новый пароль (минимум 8 символов)"
// @Success 200  {object}  map[string]string  "Пароль обновлен успешно!"
// @Failure 400  {object}  map[string]string  "Неверно набран логин или пароль / Пользователь не найден / Ошибка при хешировании пароля"
// @Failure 401  {object}  map[string]string  "Пользователь не авторизован"
// @Failure 409  {object}  map[string]string  "Ошибка при обновлении пароля, попробуйте еще раз"
// @Router /auth/update [put]
func (h *Handler) ChangePassword(c *gin.Context) {
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

	var input struct {
		NewPassword string `json:"newpassword" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Warn.Warn("Неверный ввод данных при смене пароля",
			"Email", Email,
			"Ip", c.ClientIP(),
			"Error", err.Error(),
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверно набран логин или пароль"})
		return
	}

	var user models.User

	result := h.DB.Where("mail=?", Email).First(&user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.NewPassword)); err == nil {
		logger.Warn.Warn("Попытка смены пароля на текущий",
			"Email", Email,
			"Ip", c.ClientIP(),
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Новый пароль не должен совпадать с текущим"})
		return
	}

	if result.Error != nil {
		logger.Error.Error("Ошибка при получении пользователя из базы данных",
			"Email", Email,
			"Ip", c.ClientIP(),
			"Error", result.Error.Error(),
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не найден"})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.Error.Error("Ошибка при хешировании пароля",
			"Email", Email,
			"Ip", c.ClientIP(),
			"Error", err.Error(),
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при хешировании пароля"})
		return
	}

	user.Password = string(hashPassword)

	if err := h.DB.Save(&user).Error; err != nil {
		logger.Error.Error("Ошибка при обновлении пароля в базе данных",
			"Email", Email,
			"Ip", c.ClientIP(),
			"Error", err.Error(),
		)
		c.JSON(http.StatusConflict, gin.H{"error": "Ошибка при обновлении пароля, попробуйте еще раз"})
		return
	}

	logger.Info.Info("Пароль успешно изменен",
		"Email", Email,
		"Ip", c.ClientIP(),
	)

	c.JSON(http.StatusOK, gin.H{"message": "Пароль обновлен успешно!"})

}
