package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		"produtos": BuscarProdutos(c),
	})
}

func ExibeHTMLAdminCadastrarProduto(c *gin.Context) {
	configs := BuscarConfigs()
	c.HTML(http.StatusOK, "novosProdutos.html", gin.H{
		"configs": configs,
	})
}
