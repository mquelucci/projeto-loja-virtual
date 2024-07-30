package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/database"
	"github.com/mquelucci/projeto-loja-virtual/server/models"
)

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

func EditarProduto(c *gin.Context) {

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
