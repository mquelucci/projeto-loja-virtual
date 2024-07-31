package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/database"
	"github.com/mquelucci/projeto-loja-virtual/server/models"
)

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
	c.HTML(http.StatusOK, "adminProdutos.html", gin.H{
		"configs":  BuscarConfigs(),
		"produtos": BuscarProdutos(),
	})
}

func ExibeHTMLAdminCadastrarProduto(c *gin.Context) {
	configs := BuscarConfigs()
	c.HTML(http.StatusOK, "novosProdutos.html", gin.H{
		"configs": configs,
	})
}

func ExibeHTMLAdminEditarProduto(c *gin.Context) {
	configs := BuscarConfigs()
	id := c.Query("id")
	var produto models.Produto
	database.DB.First(&produto, id)
	c.HTML(http.StatusOK, "editarProduto.html", gin.H{
		"configs": configs,
		"produto": produto,
	})
}
