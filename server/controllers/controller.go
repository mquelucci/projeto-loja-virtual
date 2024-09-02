package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/utils"
)

func ExibeHTML404(c *gin.Context) {
	configs := utils.BuscarConfigs()
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"configs": configs,
	})
}

func ExibeHTMLIndex(c *gin.Context) {
	configs := utils.BuscarConfigs()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"configs":  configs,
		"produtos": utils.BuscarProdutos(),
	})
}
