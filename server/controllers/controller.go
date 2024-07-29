package controllers

import (
	"net/http"
	"os"
	"strconv"

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

func CriarProduto(c *gin.Context) {
	var produto models.Produto
	descricao := c.PostForm("descricao")
	preco, _ := strconv.ParseFloat(c.PostForm("preco"), 64)
	quantidade, _ := strconv.Atoi(c.PostForm("quantidade"))

	produto.Descricao = descricao
	produto.Preco = preco
	produto.Quantidade = int(quantidade)

	if err := models.ValidaProduto(&produto); err != nil {
		c.HTML(http.StatusBadRequest, "novosProdutos.html", gin.H{
			"configs": BuscarConfigs(),
			"erro":    err.Error(),
		})
		return
	}
	database.DB.Create(&produto)
	c.HTML(http.StatusCreated, "novosProdutos.html", gin.H{
		"configs": BuscarConfigs(),
		"message": "Produto criado com sucesso",
	})

}

func DeletarProduto(c *gin.Context) {
	id := c.Query("id")
	var produto models.Produto
	database.DB.Delete(&produto, id)
	c.HTML(http.StatusAccepted, "adminProdutos.html", gin.H{
		"configs":  BuscarConfigs(),
		"produtos": BuscarProdutos(c),
	})
}

func BuscarConfigs() models.Config {
	var configs models.Config
	database.DB.First(&configs)
	return configs
}
