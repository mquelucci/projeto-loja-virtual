package controllers

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/controllers/responses"
	"github.com/mquelucci/projeto-loja-virtual/server/database"
	"github.com/mquelucci/projeto-loja-virtual/server/models"
	"github.com/mquelucci/projeto-loja-virtual/server/utils"
)

// BuscarTodosProdutos godoc
// @Summary Busca todos os produtos
// @Description Busca e retorna um JSON no modelo de produtos com todos os produtos não deletados
// @Tags produtos, admin
// @Produce json
// @Success 200 {object} models.Produto
// @Failure 401 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/produtos [get]
func BuscarTodosProdutos(c *gin.Context) {
	produtos := []models.Produto{}
	err := database.DB.Order("descricao ASC").Find(&produtos).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Message: err.Error()})
	}
	c.JSON(http.StatusOK, &produtos)
}

// CriarProduto godoc
// @Summary Cria um produto
// @Description Cria um produto através dos dados recebidos via formulário do cliente
// @Tags produtos, admin
// @Accept mpfd
// @Produce json
// @Param produto formData models.ProdutoBase true "Criar produto"
// @Param imagem formData file false "Imagem do Produto"
// @Success 200 {object} models.Produto
// @Failure 401 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/produtos/criar [post]
func CriarProduto(c *gin.Context) {
	var produto models.Produto
	descricao := c.PostForm("descricao")
	err := utils.ProdutoDuplo(descricao, false, &produto)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Message: err.Error()})
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
			c.JSON(http.StatusInternalServerError, gin.H{
				"erro": "Erro no tratamento de imagem do produto" + err.Error(),
			})
			return
		}
		produto.Imagem = "/assets/images/" + imagem.Filename
	}

	preco, err := strconv.ParseFloat(c.PostForm("preco"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Erro na conversão de preço" + err.Error(),
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

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "AVISO! - Erro na validação do produto: " + err.Error(),
		})
		return
	}

	err = database.DB.Create(&produto).Error
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Erro na criação do produto: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &produto)

}

func EditarProduto(c *gin.Context) {
	var produto models.Produto
	id := c.Query("id")
	database.DB.First(&produto, id)

	descricao := c.PostForm("descricao")
	err := utils.ProdutoDuplo(descricao, true, &produto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"configs": utils.BuscarConfigs(),
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
			c.JSON(http.StatusBadRequest, gin.H{
				"configs": utils.BuscarConfigs(),
				"produto": produto,
				"erro":    err.Error(),
			})
			return
		}
		produto.Imagem = "/assets/images/" + imagem.Filename
	}

	preco, _ := strconv.ParseFloat(c.PostForm("preco"), 64)
	if preco == 0.0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"configs": utils.BuscarConfigs(),
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
			"configs": utils.BuscarConfigs(),
			"produto": produto,
			"erro":    err.Error(),
		})
		return
	}

	err = database.DB.Save(&produto).Error
	if err != nil {
		c.HTML(http.StatusBadRequest, "editarProduto.html", gin.H{
			"configs": utils.BuscarConfigs(),
			"produto": produto,
			"erro":    err.Error(),
		})
		return
	}
	c.HTML(http.StatusAccepted, "editarProduto.html", gin.H{
		"configs": utils.BuscarConfigs(),
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
				"configs": utils.BuscarConfigs(),
				"produto": produto,
				"erro":    err.Error(),
			})
			return
		}
	}
	c.HTML(http.StatusAccepted, "editarProduto.html", gin.H{
		"configs": utils.BuscarConfigs(),
		"produto": produto,
		"message": "Imagem removida com sucesso",
	})

}

func DeletarProduto(c *gin.Context) {
	session := sessions.Default(c)
	id := c.Query("id")
	var produto models.Produto
	err := database.DB.First(&produto, id).Error
	if err != nil {
		session.AddFlash("Erro ao tentar encontrar o produto na base de dados", "MsgFalha")
		session.Save()
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	err = database.DB.Delete(&produto).Error
	if err != nil {
		session.AddFlash("Erro ao tentar deletar o produto na base de dados", "MsgFalha")
		session.Save()
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	session.AddFlash("Produto deletado com sucesso", "MsgSucesso")
	session.Save()
	c.JSON(http.StatusAccepted, gin.H{
		"produto": produto,
	})
}
