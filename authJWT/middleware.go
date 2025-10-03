package authjwt

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fresh132/authenticationback/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		Tokenstring := c.GetHeader("Authorization")

		if Tokenstring == "" {
			logger.Warn.Warn("Не передан токен",
				"Ip", c.ClientIP(),
			)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Не передан токен"})
			c.Abort()
			return
		}

		if len(Tokenstring) > 7 && Tokenstring[:7] == "Bearer " {
			Tokenstring = Tokenstring[7:]
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(Tokenstring, claims, func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			logger.Warn.Warn("Токен не валиден",
				"Ip", c.ClientIP(),
				"Error", err.Error(),
			)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Токен не валиден"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)

		c.Next()
	}
}
