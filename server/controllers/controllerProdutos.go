package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/database"
	"github.com/mquelucci/projeto-loja-virtual/server/models"
)

func BuscarProdutos() []models.Produto {
	produtos := []models.Produto{}
	database.DB.Find(&produtos)
	return produtos
}

func CriarProduto(c *gin.Context) {
	// contentType := c.GetHeader("Content-Type")
	// log.Println(contentType)
	// if contentType != "" || contentType != "multipart/form-data" {
	// 	c.HTML(http.StatusBadRequest, "novosProdutos.html", gin.H{
	// 		"configs": BuscarConfigs(),
	// 		"erro":    "Content-Type não é multipart/form-data",
	// 	})
	// 	return
	// }
	var produto models.Produto
	descricao := c.PostForm("descricao")
	preco, _ := strconv.ParseFloat(c.PostForm("preco"), 64)
	quantidade, _ := strconv.Atoi(c.PostForm("quantidade"))
	imagem, err := c.FormFile("imagem")
	if err != nil {
		log.Println("Nenhum arquivo carregado. Salvando produto no banco de dados.")
	}

	savePath := "./client/assets/images/" + imagem.Filename
	err = c.SaveUploadedFile(imagem, savePath)
	if err != nil {
		c.HTML(http.StatusBadRequest, "novosProdutos.html", gin.H{
			"configs": BuscarConfigs(),
			"erro":    "Erro ao salvar a imagem" + err.Error(),
		})
		return
	}

	if preco == 0.0 {
		c.HTML(http.StatusBadRequest, "novosProdutos.html", gin.H{
			"configs": BuscarConfigs(),
			"erro":    "Preço precisa ser diferente de zero",
		})
		return
	}

	produto.Descricao = descricao
	produto.Preco = preco
	produto.Quantidade = int(quantidade)
	produto.Imagem = "/assets/images/" + imagem.Filename

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
	var produto models.Produto
	id := c.Query("id")
	database.DB.First(&produto, id)
	descricao := c.PostForm("descricao")
	preco, _ := strconv.ParseFloat(c.PostForm("preco"), 64)
	quantidade, _ := strconv.Atoi(c.PostForm("quantidade"))

	if preco == 0.0 {
		c.HTML(http.StatusBadRequest, "editarProduto.html", gin.H{
			"configs": BuscarConfigs(),
			"produto": produto,
			"erro":    "Preço precisa ser diferente de zero",
		})
		return
	}

	produto.Descricao = descricao
	produto.Preco = preco
	produto.Quantidade = int(quantidade)

	if err := models.ValidaProduto(&produto); err != nil {
		c.HTML(http.StatusBadRequest, "editarProduto.html", gin.H{
			"configs": BuscarConfigs(),
			"produto": produto,
			"erro":    err.Error(),
		})
		return
	}

	database.DB.Save(&produto)
	c.HTML(http.StatusCreated, "editarProduto.html", gin.H{
		"configs": BuscarConfigs(),
		"produto": produto,
		"message": "Produto editado com sucesso",
	})
}

func DeletarProduto(c *gin.Context) {
	id := c.Query("id")
	var produto models.Produto
	database.DB.Delete(&produto, id)
	c.HTML(http.StatusAccepted, "adminProdutos.html", gin.H{
		"configs":  BuscarConfigs(),
		"produtos": BuscarProdutos(),
	})
}
