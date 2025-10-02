package authjwt

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		Tokenstring := c.GetHeader("Authorization")

		if Tokenstring == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Не передан токен"})
			c.Abort()
			return
		}

		if len(Tokenstring) > 7 && Tokenstring[:7] == "Bearer " {
			Tokenstring = Tokenstring[:7]
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(Tokenstring, claims, func(t *jwt.Token) (interface{}, error) { return []byte(os.Getenv("JWT_SECRET")), nil })

		if err != nil || !token.Valid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Токен не валиден"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)

		c.Next()
	}
}
