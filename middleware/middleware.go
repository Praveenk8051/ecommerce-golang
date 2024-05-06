package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	token "github.com/praveen/ecommerce/tokens"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientToken := c.Request.Header.Get("token")
		if ClientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No authorisation provided"})
			c.Abort()
			return
		}
		claims, err := token.ValidateTokens(ClientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}
