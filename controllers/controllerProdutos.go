package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/controllers/responses"
	"github.com/mquelucci/projeto-loja-virtual/database"
	"github.com/mquelucci/projeto-loja-virtual/models"
	"github.com/mquelucci/projeto-loja-virtual/utils"
)

// BuscarTodosProdutos godoc
// @Summary Busca todos os produtos
// @Description Busca e retorna um JSON no modelo de produtos com todos os produtos não deletados
// @Tags produtos
// @Produce json
// @Success 200 {object} responses.Message{data=[]models.Produto}
// @Failure 401 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/produtos/todos [get]
func BuscarTodosProdutos(c *gin.Context) {
	produtos := []models.Produto{}
	err := database.DB.Order("descricao ASC").Find(&produtos).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Erro: err.Error()})
	}
	c.JSON(http.StatusOK, responses.Message{Message: "Produtos encontrados", Data: produtos})
}

// BuscarProdutoPorId godoc
// @Summary Busca um produto pelo seu ID
// @Description Busca e retorna um JSON no modelo de produtos com o produto que possui o ID informado
// @Tags produtos
// @Produce json
// @Param id path int true "ID do produto"
// @Success 200 {object} responses.Message{data=[]models.Produto}
// @Failure 401 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Router /admin/produtos/{id} [get]
func BuscarProdutoPorId(c *gin.Context) {
	var produto models.Produto
	id := c.Param("id")

	if err := database.DB.First(&produto, id).Error; err != nil {
		c.JSON(http.StatusNotFound, responses.Error{Erro: "Produto não encontrado"})
		return
	}

	c.JSON(http.StatusOK, responses.Message{Message: "Produto encontrado", Data: produto})
}

// CriarProduto godoc
// @Summary Cria um produto
// @Description Cria um produto através dos dados recebidos via formulário do cliente
// @Tags produtos
// @Produce json
// @Param produto body models.ProdutoBase true "Criar produto"
// @Success 201 {object} responses.Message{data=models.Produto}
// @Failure 400 {object} responses.Error
// @Failure 401 {object} responses.Error
// @Failure 409 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/produtos/criar [post]
func CriarProduto(c *gin.Context) {
	var produtoBase models.ProdutoBase
	var produto models.Produto

	if err := c.ShouldBindJSON(&produtoBase); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Não foi possível converter o JSON para o modelo de Produtos [" + err.Error() + "]"})
		return
	}

	err := utils.ProdutoDuplo(produtoBase.Descricao, false, &produtoBase)
	if err != nil {
		c.JSON(http.StatusConflict, responses.Error{Erro: err.Error()})
		return
	}

	produto.Imagem = "/assets/images/not_found.png"
	produto.ProdutoBase = produtoBase

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
// @Tags produtos
// @Produce json
// @Param id query int true "Id do produto"
// @Param produto body models.ProdutoBase true "Dados do produto"
// @Success 202 {object} responses.Message{data=models.Produto}
// @Failure 400 {object} responses.Error
// @Failure 401 {object} responses.Error
// @Failure 409 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/produtos/editar [put]
func EditarProduto(c *gin.Context) {
	var produto models.Produto
	var produtoBase models.ProdutoBase
	id := c.Query("id")
	database.DB.First(&produto, id)

	if err := c.ShouldBindJSON(&produtoBase); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Não foi possível converter o JSON para o modelo de Produtos [" + err.Error() + "]"})
		return
	}

	err := utils.ProdutoDuplo(produtoBase.Descricao, true, &produtoBase)
	if err != nil {
		c.JSON(http.StatusConflict, responses.Error{Erro: err.Error()})
		return
	}

	produto.Imagem = "/assets/images/not_found.png"
	produto.ProdutoBase = produtoBase

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

// AdicionarImagemProduto godoc
// @Summary Adiciona a imagem de um produto
// @Description Adiciona a imagem de um produto através do id informado na url e da imagem enviada via formulário
// @Tags produtos
// @Produce json
// @Param id path int true "Id do produto"
// @Param imagem formData file false "Imagem do Produto"
// @Success 202 {object} responses.Message{data=models.Produto}
// @Failure 401 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/produtos/adicionaImagem/{id} [put]
func AdicionarImagemProduto(c *gin.Context) {
	var produto models.Produto
	id := c.Param("id")
	database.DB.First(&produto, id)

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

	if err = database.DB.Save(&produto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{
			Erro: "Erro ao tentar salvar a imagem do produto no banco de dados: [" + err.Error() + "]",
		})
	}

	c.JSON(http.StatusAccepted, responses.Message{Message: "Imagem adicionada com sucesso", Data: produto})
}

// RemoverImagemProduto godoc
// @Summary Remove a imagem de um produto
// @Description Remove a imagem de um produto específico através do Id fornecido via URL
// @Tags produtos
// @Produce json
// @Param id path int true "Id do produto"
// @Success 202 {object} responses.Message{data=models.Produto}
// @Failure 400 {object} responses.Error
// @Failure 401 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/produtos/removeImagem/{id} [delete]
func RemoverImagemProduto(c *gin.Context) {
	id := c.Param("id")
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
// @Tags produtos
// @Produce json
// @Param id query int true "Id do produto"
// @Success 202 {object} responses.Message{data=models.Produto}
// @Failure 401 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/produtos/deletar [delete]
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
