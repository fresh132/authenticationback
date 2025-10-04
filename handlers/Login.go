package handlers

import (
	"net/http"

	authjwt "github.com/fresh132/authenticationback/authJWT"
	"github.com/fresh132/authenticationback/logger"
	"github.com/fresh132/authenticationback/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Summary      Login
// @Description  Вход по email и паролю, возвращает JWT-токен
// @ID           Login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body  models.PasswordMailRequest  true  "Данные для входа"
// @Success      200  {object}  map[string]string  "Вы успешно вошли в систему + токен"
// @Failure      400  {object}  map[string]string  "Неверный email или пароль"
// @Failure      500  {object}  map[string]string  "Ошибка генерации токена"
// @Router       /enter [post]
func (h *Handler) Login(c *gin.Context) {
	var input struct {
		Mail     string `json:"mail" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Warn.Warn("Неверный ввод данных при входе в систему",
			"Email", input.Mail,
			"Ip", c.ClientIP(),
			"Error", err.Error(),
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверно набран логин или пароль"})
		return
	}

	var user models.User

	result := h.DB.Where("mail=?", input.Mail).First(&user)

	if result.Error != nil {
		logger.Error.Error("Ошибка при получении пользователя из базы данных",
			"Email", input.Mail,
			"Ip", c.ClientIP(),
			"Error", result.Error.Error(),
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный email"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		logger.Warn.Warn("Неверный пароль при попытке входа в систему",
			"Email", input.Mail,
			"Ip", c.ClientIP(),
			"Error", err.Error(),
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный пароль"})
		return
	}

	tokenString, err := authjwt.GenerateToken(user.ID.String(), user.Mail)
	if err != nil {
		logger.Error.Error("Ошибка при генерации токена",
			"Email", input.Mail,
			"Ip", c.ClientIP(),
			"Error", err.Error(),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка генерации токена"})
		return
	}

	logger.Info.Info("Пользователь успешно вошел в систему",
		"Email", input.Mail,
		"Ip", c.ClientIP(),
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Вы успешно вошли в систему",
		"token":   tokenString,
	})

}
