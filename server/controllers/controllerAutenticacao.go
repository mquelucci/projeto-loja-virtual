package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/database"
	"github.com/mquelucci/projeto-loja-virtual/server/models"
)

func Autenticar(c *gin.Context) {
	//Grava usuario e senha recebidos do formulário
	usuario := c.PostForm("usuario")
	senha := c.PostForm("senha")

	//Cria uma variável do modelo Admin para receber os dados
	//de acesso de administrador cadastrados no banco,
	//caso o usuário e senha informados coincidam com o que tem no banco de dados
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
