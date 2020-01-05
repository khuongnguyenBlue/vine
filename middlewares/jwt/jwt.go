package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/khuongnguyenBlue/vine/utils"
	"net/http"
	"strings"
)

type JWTHeader struct {
	Token string `header:"Authorization" binding:"required"`
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		msg := "Invalid token"
		var jwtHeader JWTHeader
		if err := c.ShouldBindHeader(&jwtHeader); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": msg})
			return
		}

		token := strings.TrimPrefix(jwtHeader.Token, "Bearer ")
		claims, err := utils.ParseToken(token)
		if err != nil {
			if err.(*jwt.ValidationError).Errors == jwt.ValidationErrorExpired {
				msg = "Token has been expired"
			}

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": msg})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
	}
}
