package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/database"
	"github.com/mquelucci/projeto-loja-virtual/server/models"
)

func ExibeHTMLIndex(c *gin.Context) {
	var configs models.Config
	database.DB.First(&configs)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"configs":  configs,
		"produtos": BuscarProdutos(c),
	})
}

func ExibeHTMLAdmin(c *gin.Context) {
	var configs models.Config
	database.DB.First(&configs)
	c.HTML(http.StatusOK, "admin.html", gin.H{
		"configs": configs,
	})
}

func ExibeHTMLAdminLogin(c *gin.Context) {
	var configs models.Config
	database.DB.First(&configs)
	c.HTML(http.StatusOK, "adminLogin", gin.H{
		"configs": configs,
	})
}

func FazerLogin(c *gin.Context) {
	usuario := c.PostForm("usuario")
	senha := c.PostForm("senha")

	var admin models.Admin
	if err := database.DB.Where("nome = ? AND senha = ?", usuario, senha).First(&admin).Error; err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"Error": "Invalid credentials",
		})
		return
	}
}

func BuscarProdutos(c *gin.Context) []models.Produto {
	produtos := []models.Produto{}
	database.DB.Find(&produtos)
	return produtos
}

func CriarProduto(c *gin.Context) {

}
