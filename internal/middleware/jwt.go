package middleware

import (
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// cookie, _ := c.Cookie("login_token")
		// fmt.Println("cookie", cookie)
		c.Next()
	}
}
