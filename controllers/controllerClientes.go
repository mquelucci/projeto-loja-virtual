package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/controllers/responses"
	"github.com/mquelucci/projeto-loja-virtual/database"
	"github.com/mquelucci/projeto-loja-virtual/models"
	"github.com/mquelucci/projeto-loja-virtual/utils"
)

// BuscarTodosClientes godoc
// @Summary Busca todos os clientes da loja virtual
// @Description Busca todos os clientes da loja virtual e retorna no JSON
// @Tags admin, clientes
// @Produce json
// @Success 200 {object} responses.Message{data=[]models.Cliente}
// @Failure 401 {object} responses.Error
// @Router /admin/clientes/todos [get]
func BuscarTodosClientes(c *gin.Context) {
	var clientes []models.Cliente
	database.DB.Preload("Endereco").Find(&clientes)
	c.JSON(http.StatusOK, responses.Message{Message: "Clientes encontrados", Data: clientes})
}

// BuscarCliente godoc
// @Summary Busca o cliente da loja virtual pelo seu CPF ou CNPJ
// @Description Busca o cliente da loja virtual pelo seu CPF ou CNPJ e retorna no JSON
// @Tags admin, clientes
// @Produce json
// @Param cpf_cnpj path int true "CPF_CNPJ"
// @Success 200 {object} responses.Message{data=models.Cliente}
// @Failure 401 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/clientes/{cpf_cnpj} [get]
func BuscarCliente(c *gin.Context) {
	var cliente models.Cliente
	cpfCnpj := c.Param("cpf_cnpj")

	if err := database.DB.Where("cpf_cnpj = ?", cpfCnpj).Preload("Endereco").First(&cliente).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Erro: "Não foi possível encontrar o cliente" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.Message{Message: "Cliente encontrado", Data: cliente})
}

// CriarCliente godoc
// @Summary Cria um cliente da loja virtual
// @Description Cria um cliente da loja virtual conforme informações enviadas pelo formulário
// @Tags admin, clientes
// @Produce json
// @Param cliente formData models.ClienteBase true "Dados do cliente"
// @Param endereco formData models.EnderecoBase true "Endereço do cliente"
// @Success 200 {object} responses.Message{data=models.Cliente}
// @Failure 400 {object} responses.Error
// @Failure 401 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/clientes/criar [post]
func CriarCliente(c *gin.Context) {
	var cliente models.Cliente

	nomeCliente := c.PostForm("nome")
	cpfCnpj := c.PostForm("cpf_cnpj")
	telefone := c.PostForm("telefone")
	email := c.PostForm("email")
	endereco := c.PostForm("endereco")
	numero := c.PostForm("numero")
	bairro := c.PostForm("bairro")
	cep := c.PostForm("cep")
	cidade := c.PostForm("cidade")
	uf := c.PostForm("uf")

	numeroConvertido, err := strconv.Atoi(numero)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Não foi possível converter o número informado" + err.Error()})
		return
	}

	if len(cep) != 8 {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "O CEP informado não possui 8 dígitos"})
		return
	}

	cepConvertido, err := strconv.Atoi(cep)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Não foi possível converter o CEP informado" + err.Error()})
		return
	}

	cliente.Nome = nomeCliente
	cliente.CpfCnpj = cpfCnpj
	cliente.Telefone = telefone
	cliente.Email = email

	if err := models.ValidaEndereco(&models.Endereco{Endereco: endereco, Numero: numeroConvertido, Bairro: bairro, CEP: cepConvertido, Cidade: cidade, UF: uf}); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Erro na validação do endereço: " + err.Error()})
		return
	}

	cliente.Endereco = append([]models.Endereco{}, models.Endereco{Endereco: endereco, Numero: numeroConvertido, Bairro: bairro, CEP: cepConvertido, Cidade: cidade, UF: uf})

	if err := models.ValidaCliente(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Erro na validação do cliente: " + err.Error()})
		return
	}

	if err := utils.ClienteDuplo(cpfCnpj, false, &cliente); err != nil {
		c.JSON(http.StatusConflict, responses.Error{Erro: err.Error()})
		return
	}

	if err := database.DB.Create(&cliente).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Erro: "Erro na criação do cliente: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, responses.Message{Message: "Cliente criado com sucesso", Data: cliente})
}

// DeletarCliente godoc
// @Summary Deletar um cliente da loja virtual (soft-delete)
// @Description Deleta um cliente da loja virtual conforme cpf/cnpj informadas na URL
// @Tags admin, clientes
// @Produce json
// @Param cpf_cnpj path int true "CPF_CNPJ"
// @Success 200 {object} responses.Message{data=models.Cliente}
// @Failure 401 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/clientes/deletar/{cpf_cnpj} [delete]
func DeletarCliente(c *gin.Context) {
	var cliente models.Cliente
	cpfCnpj := c.Param("cpf_cnpj")

	if err := database.DB.Where("cpf_cnpj =?", cpfCnpj).First(&cliente).Error; err != nil {
		c.JSON(http.StatusNotFound, responses.Error{Erro: "Cliente não encontrado"})
		return
	}

	if err := database.DB.Delete(&cliente).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Erro: "Erro na exclusão do cliente: " + err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, responses.Message{Message: "Cliente excluído com sucesso", Data: cliente})
}
