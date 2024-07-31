package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/database"
	"github.com/mquelucci/projeto-loja-virtual/server/models"
)

func FazerLogin(c *gin.Context) {
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
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/admin")
}

func FazerLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("auth")
	session.Save()
	c.Redirect(http.StatusFound, "/admin/login")
}
