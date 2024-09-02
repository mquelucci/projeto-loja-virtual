package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/database"
	"github.com/mquelucci/projeto-loja-virtual/server/models"
	"github.com/mquelucci/projeto-loja-virtual/server/utils"
)

func ExibeHTMLAdmin(c *gin.Context) {
	configs := utils.BuscarConfigs()
	c.HTML(http.StatusOK, "admin.html", gin.H{
		"configs": configs,
	})
}

func ExibeHTMLAdminLogin(c *gin.Context) {
	configs := utils.BuscarConfigs()
	c.HTML(http.StatusOK, "login.html", gin.H{
		"configs": configs,
	})
}

func ExibeHTMLAdminProdutos(c *gin.Context) {
	session := sessions.Default(c)
	msgSucesso := session.Flashes("MsgSucesso")
	msgFalha := session.Flashes("MsgFalha")
	msgInfo := session.Flashes("MsgInfo")
	session.Flashes()
	session.Save()
	c.HTML(http.StatusOK, "adminProdutos.html", gin.H{
		"configs":    utils.BuscarConfigs(),
		"produtos":   utils.BuscarProdutos(),
		"MsgSucesso": msgSucesso,
		"MsgInfo":    msgInfo,
		"MsgFalha":   msgFalha,
	})
}

func ExibeHTMLAdminCadastrarProduto(c *gin.Context) {
	session := sessions.Default(c)
	msgSucesso := session.Flashes("MsgSucesso")
	msgFalha := session.Flashes("MsgFalha")
	msgInfo := session.Flashes("MsgInfo")
	session.Flashes()
	session.Save()
	c.HTML(http.StatusOK, "novosProdutos.html", gin.H{
		"configs":    utils.BuscarConfigs(),
		"MsgSucesso": msgSucesso,
		"MsgInfo":    msgInfo,
		"MsgFalha":   msgFalha,
	})
}

func ExibeHTMLAdminEditarProduto(c *gin.Context) {
	session := sessions.Default(c)
	msgSucesso := session.Flashes("MsgSucesso")
	msgFalha := session.Flashes("MsgFalha")
	msgInfo := session.Flashes("MsgInfo")
	session.Flashes()
	session.Save()

	configs := utils.BuscarConfigs()
	id := c.Query("id")
	var produto models.Produto
	database.DB.First(&produto, id)

	c.HTML(http.StatusOK, "editarProduto.html", gin.H{
		"configs":    configs,
		"produto":    produto,
		"MsgSucesso": msgSucesso,
		"MsgInfo":    msgInfo,
		"MsgFalha":   msgFalha,
	})
}
