package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		auth := session.Get("auth")
		if auth == nil {
			c.Redirect(http.StatusFound, "/admin/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
