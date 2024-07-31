package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/database"
	"github.com/mquelucci/projeto-loja-virtual/server/models"
)

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

func BuscarConfigs() models.Config {
	var configs models.Config
	database.DB.First(&configs)
	return configs
}
