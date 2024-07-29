package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/database"
	"github.com/mquelucci/projeto-loja-virtual/server/models"
)

var server = os.Getenv("SERVER")

func ExibeHTML404(c *gin.Context) {
	configs := BuscarConfigs()
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"configs": configs,
	})
}

func ExibeHTMLIndex(c *gin.Context) {
	configs := BuscarConfigs()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"configs":  configs,
		"produtos": BuscarProdutos(c),
	})
}

func ExibeHTMLAdmin(c *gin.Context) {
	configs := BuscarConfigs()
	c.HTML(http.StatusOK, "admin.html", gin.H{
		"configs": configs,
	})
}

func ExibeHTMLAdminLogin(c *gin.Context) {
	configs := BuscarConfigs()
	c.HTML(http.StatusOK, "login.html", gin.H{
		"configs": configs,
	})
}

func ExibeHTMLAdminProdutos(c *gin.Context) {
	configs := BuscarConfigs()
	var produtos []models.Produto
	database.DB.Find(&produtos)
	c.HTML(http.StatusOK, "adminProdutos.html", gin.H{
		"configs":  configs,
		"produtos": produtos,
	})
}

func ExibeHTMLAdminCadastrarProduto(c *gin.Context) {
	configs := BuscarConfigs()
	c.HTML(http.StatusOK, "novosProdutos.html", gin.H{
		"configs": configs,
	})
}

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

	c.SetCookie("auth", "true", 3600, "/", server, false, true)
	c.Redirect(http.StatusMovedPermanently, "/admin")
}

func FazerLogout(c *gin.Context) {
	c.SetCookie("auth", "", 0, "/", server, false, true)
	c.Redirect(http.StatusFound, "/admin/login")
}

func BuscarProdutos(c *gin.Context) []models.Produto {
	produtos := []models.Produto{}
	database.DB.Find(&produtos)
	return produtos
}

func BuscarConfigs() models.Config {
	var configs models.Config
	database.DB.First(&configs)
	return configs
}
