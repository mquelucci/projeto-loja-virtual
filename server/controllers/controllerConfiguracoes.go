package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/database"
	"github.com/mquelucci/projeto-loja-virtual/server/models"
)

func AlterarConfiguracoes(c *gin.Context) {
	var configs models.Config
	database.DB.First(&configs)
	nomeLoja := c.PostForm("nomeLoja")
	configs.NomeLoja = nomeLoja
	database.DB.Save(&configs)
	c.HTML(http.StatusOK, "admin.html", gin.H{
		"configs": BuscarConfigs(),
	})
}
