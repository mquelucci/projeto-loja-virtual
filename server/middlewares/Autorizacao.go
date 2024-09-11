package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/controllers/responses"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		auth := session.Get("auth")
		if auth == nil {
			c.JSON(http.StatusUnauthorized, responses.Error{Erro: "Usuário não autorizado!"})
			c.Abort()
			return
		}
		c.Next()
	}
}
