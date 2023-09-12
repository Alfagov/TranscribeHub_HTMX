package middlewares

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Authentication")
		if err != nil || token == "" {
			c.Set("logged_in", false)
			c.Next()
			return
		}
		c.Set("logged_in", true)
		c.Set("user_id", token)
		c.Next()
	}
}
