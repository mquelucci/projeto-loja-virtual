package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/database"
	"github.com/mquelucci/projeto-loja-virtual/server/models"
	_ "github.com/mquelucci/projeto-loja-virtual/server/responses"
)

// Autenticar Faz a autenticação do usuário
// @Summary Faz a autenticação do usuário
// @Description Através dos dados fornecidos via formulário HTML,
// @Description compara com o banco de dados para autenticar ou rejeitar
// @Tags auth,user
// @Success 301
// @Failure 404 {object} responses.Error
// @Router /autenticar [post]
func Autenticar(c *gin.Context) {
	usuario := c.PostForm("usuario")
	senha := c.PostForm("senha")
	var admin models.Admin
	if err := database.DB.Where("nome = ? AND senha = ?", usuario, senha).First(&admin).Error; err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"Error": "Credenciais inválidas",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("auth", true)
	err := session.Save()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/admin")
}

func FazerLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("auth")
	err := session.Save()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Redirect(http.StatusFound, "/admin/login")
}
