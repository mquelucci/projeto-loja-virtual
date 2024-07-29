package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth, err := c.Cookie("auth")
		log.Println(auth)
		if err != nil || auth != "true" {
			c.Redirect(http.StatusFound, "/admin/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
