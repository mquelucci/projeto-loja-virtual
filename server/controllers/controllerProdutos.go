package controllers

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/database"
	"github.com/mquelucci/projeto-loja-virtual/server/models"
	"github.com/mquelucci/projeto-loja-virtual/server/utils"
)

func BuscarProdutos() []models.Produto {
	produtos := []models.Produto{}
	database.DB.Order("descricao ASC").Find(&produtos)
	return produtos
}

func CriarProduto(c *gin.Context) {
	var produto models.Produto

	descricao := c.PostForm("descricao")
	err := utils.ProdutoDuplo(descricao, false, &produto)
	if err != nil {
		c.HTML(http.StatusBadRequest, "novosProdutos.html", gin.H{
			"configs": BuscarConfigs(),
			"erro":    err.Error(),
		})
		return
	}
	produto.Descricao = descricao

	imagem, err := c.FormFile("imagem")
	if err != nil {
		log.Println("Nenhum arquivo carregado. Criando produto sem imagem.")
		produto.Imagem = "/assets/images/not_found.png"
	} else {
		err := utils.TratarImagemProduto(c, imagem, &produto)
		if err != nil {
			c.HTML(http.StatusBadRequest, "novosProdutos.html", gin.H{
				"configs": BuscarConfigs(),
				"erro":    err.Error(),
			})
			return
		}
		produto.Imagem = "/assets/images/" + imagem.Filename
	}

	preco, _ := strconv.ParseFloat(c.PostForm("preco"), 64)
	if preco == 0.0 {
		c.HTML(http.StatusBadRequest, "novosProdutos.html", gin.H{
			"configs": BuscarConfigs(),
			"erro":    "Preço precisa ser diferente de zero",
		})
		return
	}
	produto.Preco = preco

	quantidade, _ := strconv.Atoi(c.PostForm("quantidade"))
	produto.Quantidade = int(quantidade)

	ativo := c.PostForm("ativo")
	if ativo == "on" {
		produto.Ativo = true
	} else {
		produto.Ativo = false
	}

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
		return
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
	err := utils.ProdutoDuplo(descricao, true, &produto)
	if err != nil {
		c.HTML(http.StatusBadRequest, "editarProduto.html", gin.H{
			"configs": BuscarConfigs(),
			"produto": produto,
			"erro":    err.Error(),
		})
		return
	}
	produto.Descricao = descricao

	imagem, err := c.FormFile("imagem")
	if err != nil {
		log.Println("Nenhum arquivo carregado. Mantendo o registro de imagem do produto.")
	} else {
		err := utils.TratarImagemProduto(c, imagem, &produto)
		if err != nil {
			c.HTML(http.StatusBadRequest, "editarProduto.html", gin.H{
				"configs": BuscarConfigs(),
				"produto": produto,
				"erro":    err.Error(),
			})
			return
		}
		produto.Imagem = "/assets/images/" + imagem.Filename
	}

	preco, _ := strconv.ParseFloat(c.PostForm("preco"), 64)
	if preco == 0.0 {
		c.HTML(http.StatusBadRequest, "editarProduto.html", gin.H{
			"configs": BuscarConfigs(),
			"produto": produto,
			"erro":    "Preço precisa ser diferente de zero",
		})
		return
	}
	produto.Preco = preco

	quantidade, _ := strconv.Atoi(c.PostForm("quantidade"))
	produto.Quantidade = int(quantidade)

	ativo := c.PostForm("ativo")
	if ativo == "on" {
		produto.Ativo = true
	} else {
		produto.Ativo = false
	}

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
		return
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
	if produto.Imagem != "/assets/images/not_found.png" {
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
			return
		}
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
	err := database.DB.Delete(&produto, id).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{})
}
