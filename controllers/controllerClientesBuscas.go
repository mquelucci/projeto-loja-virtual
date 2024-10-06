package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/controllers/responses"
	"github.com/mquelucci/projeto-loja-virtual/database"
	"github.com/mquelucci/projeto-loja-virtual/models"
)

// BuscarTodosClientes godoc
// @Summary Busca todos os clientes da loja virtual
// @Description Busca todos os clientes da loja virtual e retorna no JSON
// @Tags clientes
// @Produce json
// @Success 200 {object} responses.Message{data=[]models.Cliente}
// @Failure 401 {object} responses.Error
// @Router /admin/clientes/todos [get]
func BuscarTodosClientes(c *gin.Context) {
	var clientes []models.Cliente
	database.DB.Find(&clientes)
	c.JSON(http.StatusOK, responses.Message{Message: "Clientes encontrados", Data: clientes})
}

// BuscarClientePorCpfCnpj godoc
// @Summary Busca o cliente da loja virtual pelo seu CPF ou CNPJ
// @Description Busca o cliente da loja virtual pelo seu CPF ou CNPJ e retorna no JSON
// @Tags clientes
// @Produce json
// @Param cpf_cnpj path int true "CPF_CNPJ"
// @Success 200 {object} responses.Message{data=models.Cliente}
// @Failure 400 {object} responses.Error
// @Failure 401 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Router /admin/clientes/{cpf_cnpj} [get]
func BuscarClientePorCpfCnpj(c *gin.Context) {
	var cliente models.Cliente
	cpfCnpj := c.Param("cpf_cnpj")
	if cpfCnpj == "" {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "O CPF/CNPJ não foi informado"})
	}
	if err := database.DB.Where("cpf_cnpj = ?", cpfCnpj).First(&cliente).Error; err != nil {
		c.JSON(http.StatusNotFound, responses.Error{Erro: "Não foi possível encontrar o cliente"})
		return
	}

	c.JSON(http.StatusOK, responses.Message{Message: "Cliente encontrado", Data: cliente})
}

// BuscarClientesPorCep godoc
// @Summary Busca os clientes da loja virtual que residem no CEP informado
// @Description Busca os clientes da loja virtual que residem no CEP informado e retorna no JSON
// @Tags clientes
// @Produce json
// @Param cep path int true "CEP"
// @Success 200 {object} responses.Message{data=[]models.Cliente}
// @Failure 400 {object} responses.Error
// @Failure 401 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Router /admin/clientes/cep/{cep} [get]
func BuscarClientesPorCep(c *gin.Context) {
	var clientes []models.Cliente
	cep := c.Param("cep")
	if cep == "" {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "O CEP não foi informado"})
	}
	if err := database.DB.Where("cep", cep).Find(&clientes).Error; err != nil {
		c.JSON(http.StatusNotFound, responses.Error{Erro: "Não foi possível encontrar os clientes desse CEP"})
		return
	}

	c.JSON(http.StatusOK, responses.Message{Message: "Clientes encontrados", Data: clientes})

}

// BuscarClientesPorEmpresa godoc
// @Summary Busca os clientes da loja virtual que tem a empresa informada no cadastro
// @Description Busca os clientes da loja virtual que tem a empresa informada no cadastro e retorna no JSON
// @Tags clientes
// @Produce json
// @Param empresa path string true "EMPRESA"
// @Success 200 {object} responses.Message{data=[]models.Cliente}
// @Failure 400 {object} responses.Error
// @Failure 401 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Router /admin/clientes/empresa/{empresa} [get]
func BuscarClientesPorEmpresa(c *gin.Context) {
	var clientes []models.Cliente
	empresa := c.Param("empresa")
	if empresa == "" {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "A empresa não foi informada"})
	}
	if err := database.DB.Where("empresa LIKE ?", "%"+empresa+"%").Find(&clientes).Error; err != nil {
		c.JSON(http.StatusNotFound, responses.Error{Erro: "Não foi possível encontrar o clientes dessa empresa"})
		return
	}

	c.JSON(http.StatusOK, responses.Message{Message: "Clientes encontrados", Data: clientes})

}

// BuscarClientesPorNome godoc
// @Summary Busca os clientes da loja virtual que possuem o nome informado
// @Description Busca os clientes da loja virtual que possuem o nome informado no cadastro e retorna no JSON
// @Tags clientes
// @Produce json
// @Param nome path string true "NOME"
// @Success 200 {object} responses.Message{data=[]models.Cliente}
// @Failure 400 {object} responses.Error
// @Failure 401 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Router /admin/clientes/nome/{nome} [get]
func BuscarClientesPorNome(c *gin.Context) {
	var clientes []models.Cliente
	nome := c.Param("nome")
	if nome == "" {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "A nome não foi informada"})
	}
	if err := database.DB.Where("nome LIKE ?", "%"+nome+"%").Find(&clientes).Error; err != nil {
		c.JSON(http.StatusNotFound, responses.Error{Erro: "Não foi possível encontrar o clientes dessa nome"})
		return
	}

	c.JSON(http.StatusOK, responses.Message{Message: "Clientes encontrados", Data: clientes})

}
