package controllers

import (
	"log"
	"net/http"
	"os"
	"strconv"

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
// @Success 200 {object} responses.Message{data=[]models.Produto}
// @Failure 401 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/produtos [get]
func BuscarTodosProdutos(c *gin.Context) {
	produtos := []models.Produto{}
	err := database.DB.Order("descricao ASC").Find(&produtos).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Erro: err.Error()})
	}
	c.JSON(http.StatusOK, responses.Message{Message: "Produtos encontrados", Data: produtos})
}

// CriarProduto godoc
// @Summary Cria um produto
// @Description Cria um produto através dos dados recebidos via formulário do cliente
// @Tags produtos, admin
// @Accept mpfd
// @Produce json
// @Param produto formData models.ProdutoBase true "Criar produto"
// @Param imagem formData file false "Imagem do Produto"
// @Success 201 {object} responses.Message{data=models.Produto}
// @Failure 400 {object} responses.Error
// @Failure 401 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/produtos/criar [post]
func CriarProduto(c *gin.Context) {
	var produto models.Produto

	// Tratamento da descrição
	descricao := c.PostForm("descricao")
	err := utils.ProdutoDuplo(descricao, false, &produto)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: err.Error()})
		return
	}
	produto.Descricao = descricao

	// Tratamento da imagem
	imagem, err := c.FormFile("imagem")
	if err != nil {
		log.Println("Nenhum arquivo carregado. Criando produto sem imagem.")
		produto.Imagem = "/assets/images/not_found.png"
	} else {
		err := utils.TratarImagemProduto(c, imagem, &produto)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Error{Erro: "Erro no tratamento de imagem do produto" + err.Error()})
			return
		}
		produto.Imagem = "/assets/images/" + imagem.Filename
	}

	// Tratamento do preço
	preco, err := strconv.ParseFloat(c.PostForm("preco"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Erro na conversão de preço" + err.Error()})
		return
	}
	produto.Preco = preco

	// Tratamento da quantidade
	quantidade, err := strconv.Atoi(c.PostForm("quantidade"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Erro na conversão de quantidade" + err.Error()})
		return
	}
	produto.Quantidade = quantidade

	// Tratamento do status ativo
	ativo := c.PostForm("ativo")
	if ativo == "on" {
		produto.Ativo = true
	} else {
		produto.Ativo = false
	}

	if err := models.ValidaProduto(&produto); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Erro na validação do produto: " + err.Error()})
		return
	}

	err = database.DB.Create(&produto).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Erro: "Erro na criação do produto: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, responses.Message{Message: "Produto criado com sucesso", Data: produto})

}

// EditarProduto godoc
// @Summary Editar um produto
// @Description Editar um produto através do id informado na url e dos dados recebidos via formulário do cliente
// @Tags produtos, admin
// @Accept mpfd
// @Produce json
// @Param id query int true "Id do produto"
// @Param produto formData models.ProdutoBase true "Dados do produto"
// @Param imagem formData file false "Imagem do Produto"
// @Success 202 {object} responses.Message{data=models.Produto}
// @Failure 400 {object} responses.Error
// @Failure 401 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/produtos/editar [post]
func EditarProduto(c *gin.Context) {
	var produto models.Produto
	id := c.Query("id")
	database.DB.First(&produto, id)

	// Tratamento da descrição
	descricao := c.PostForm("descricao")
	if err := utils.ProdutoDuplo(descricao, true, &produto); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: err.Error()})
		return
	}
	produto.Descricao = descricao

	// Tratamento da imagem
	imagem, err := c.FormFile("imagem")
	if err != nil {
		log.Println("Nenhum arquivo carregado. Mantendo o registro de imagem do produto.")
	} else {
		if err := utils.TratarImagemProduto(c, imagem, &produto); err != nil {
			c.JSON(http.StatusInternalServerError, responses.Error{Erro: "Erro no tratamento de imagem do produto" + err.Error()})
			return
		}
		produto.Imagem = "/assets/images/" + imagem.Filename
	}

	// Tratamento do preço
	preco, err := strconv.ParseFloat(c.PostForm("preco"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Erro na conversão de preço" + err.Error()})
		return
	}
	produto.Preco = preco

	// Tratamento da quantidade
	quantidade, err := strconv.Atoi(c.PostForm("quantidade"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Erro na conversão de quantidade" + err.Error()})
		return
	}
	produto.Quantidade = quantidade

	// Tratamento do status ativo
	ativo := c.PostForm("ativo")
	if ativo == "on" {
		produto.Ativo = true
	} else {
		produto.Ativo = false
	}

	if err := models.ValidaProduto(&produto); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Erro na validação do produto: " + err.Error()})
		return
	}

	if err = database.DB.Save(&produto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Erro: "Erro na edição do produto: " + err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, responses.Message{Message: "Produto editado com sucesso", Data: produto})
}

// RemoverImagemProduto godoc
// @Summary Remove a imagem de um produto
// @Description Remove a imagem de um produto específico através do Id fornecido via URL
// @Tags produtos, admin
// @Produce json
// @Param id query int true "Id do produto"
// @Success 202 {object} responses.Message{data=models.Produto}
// @Failure 400 {object} responses.Error
// @Failure 401 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/produtos/removeImagem [post]
func RemoverImagemProduto(c *gin.Context) {
	id := c.Query("id")
	var produto models.Produto
	database.DB.First(&produto, id)
	if produto.Imagem != "/assets/images/not_found.png" {
		pathImagem := "./client" + produto.Imagem
		os.Remove(pathImagem)
		produto.Imagem = "/assets/images/not_found.png"
		if err := database.DB.Save(&produto).Error; err != nil {
			c.JSON(http.StatusBadRequest, responses.Error{Erro: "Erro na validação do produto: " + err.Error()})
			return
		}
	}
	c.JSON(http.StatusAccepted, responses.Message{Message: "Imagem removida com sucesso", Data: produto})

}

// DeletarProduto godoc
// @Summary Deleta um produto
// @Description Deleta um produto específico através do Id fornecido via URL
// @Tags produtos, admin
// @Produce json
// @Param id query int true "Id do produto"
// @Success 202 {object} responses.Message{data=models.Produto}
// @Failure 401 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/produtos/deletar [post]
func DeletarProduto(c *gin.Context) {
	id := c.Query("id")
	var produto models.Produto
	err := database.DB.First(&produto, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Erro: err.Error()})
		return
	}
	if err = database.DB.Delete(&produto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Erro: err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, responses.Message{Message: "Produto exclúido com sucesso", Data: produto})
}
