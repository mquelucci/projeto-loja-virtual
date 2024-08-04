package controllers

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/database"
	"github.com/mquelucci/projeto-loja-virtual/server/models"
)

func BuscarProdutos() []models.Produto {
	produtos := []models.Produto{}
	database.DB.Order("descricao ASC").Find(&produtos)
	return produtos
}

func CriarProduto(c *gin.Context) {
	var produto models.Produto
	descricao := c.PostForm("descricao")
	preco, _ := strconv.ParseFloat(c.PostForm("preco"), 64)
	quantidade, _ := strconv.Atoi(c.PostForm("quantidade"))
	imagem, err := c.FormFile("imagem")
	if err != nil {
		log.Println("Nenhum arquivo carregado. Salvando produto no banco de dados.")
		produto.Imagem = "/assets/images/not_found.png"
	} else {
		savePath := "./client/assets/images/" + imagem.Filename
		err = c.SaveUploadedFile(imagem, savePath)
		if err != nil {
			c.HTML(http.StatusBadRequest, "novosProdutos.html", gin.H{
				"configs": BuscarConfigs(),
				"erro":    "Erro ao salvar a imagem" + err.Error(),
			})
			return
		}
		produto.Imagem = "/assets/images/" + imagem.Filename
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

	if err := models.ValidaProduto(&produto); err != nil {
		c.HTML(http.StatusBadRequest, "novosProdutos.html", gin.H{
			"configs": BuscarConfigs(),
			"erro":    err.Error(),
		})
		return
	}
	err = database.DB.Create(&produto).Error
	if err != nil {
		c.HTML(http.StatusBadRequest, "novosProdutos.html", gin.H{
			"configs": BuscarConfigs(),
			"erro":    err.Error(),
		})
	}
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

	imagem, err := c.FormFile("imagem")
	if err != nil {
		log.Println("Nenhum arquivo carregado. Mantendo o registro de imagem do produto.")
	} else {

		savePath := "./client/assets/images/" + imagem.Filename
		if savePath == produto.Imagem {
			os.Remove(savePath)
		}
		err = c.SaveUploadedFile(imagem, savePath)
		if err != nil {
			c.HTML(http.StatusBadRequest, "novosProdutos.html", gin.H{
				"configs": BuscarConfigs(),
				"erro":    "Erro ao salvar a imagem" + err.Error(),
			})
			return
		}
		produto.Imagem = "/assets/images/" + imagem.Filename
	}

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

	err = database.DB.Save(&produto).Error
	if err != nil {
		c.HTML(http.StatusBadRequest, "editarProduto.html", gin.H{
			"configs": BuscarConfigs(),
			"produto": produto,
			"erro":    err.Error(),
		})
	}
	c.HTML(http.StatusAccepted, "editarProduto.html", gin.H{
		"configs": BuscarConfigs(),
		"produto": produto,
		"message": "Produto editado com sucesso",
	})
}

func RemoverImagemProduto(c *gin.Context) {
	id := c.Query("id")
	var produto models.Produto
	database.DB.First(&produto, id)
	pathImagem := "./client" + produto.Imagem
	os.Remove(pathImagem)
	produto.Imagem = "/assets/images/not_found.png"

	err := database.DB.Save(&produto).Error
	if err != nil {
		c.HTML(http.StatusBadRequest, "editarProduto.html", gin.H{
			"configs": BuscarConfigs(),
			"produto": produto,
			"erro":    err.Error(),
		})
	}

	c.HTML(http.StatusAccepted, "editarProduto.html", gin.H{
		"configs": BuscarConfigs(),
		"produto": produto,
		"message": "Imagem removida com sucesso",
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
