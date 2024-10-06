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

// CriarCliente godoc
// @Summary Cria um cliente da loja virtual
// @Description Cria um cliente da loja virtual conforme informações enviadas pelo formulário
// @Tags clientes
// @Produce json
// @Param cliente body models.ClienteBase true "Dados do cliente"
// @Success 200 {object} responses.Message{data=models.Cliente}
// @Failure 400 {object} responses.Error
// @Failure 401 {object} responses.Error
// @Failure 409 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/clientes/criar [post]
func CriarCliente(c *gin.Context) {
	var clienteBase models.ClienteBase

	if err := c.ShouldBindJSON(&clienteBase); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Não foi possível converter o JSON para o modelo de Cliente [" + err.Error() + "]"})
		return
	}

	if err := utils.ClienteDuplo(clienteBase.CpfCnpj, false, &clienteBase); err != nil {
		c.JSON(http.StatusConflict, responses.Error{Erro: err.Error()})
		return
	}

	if len(strconv.Itoa(clienteBase.CEP)) != 8 {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "O CEP não tem 8 dígitos"})
		return
	}

	if err := models.ValidaCliente(&clienteBase); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Erro na validação do cliente: " + err.Error()})
		return
	}

	cliente := models.Cliente{ClienteBase: clienteBase}

	if err := database.DB.Create(&cliente).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Erro: "Erro na criação do cliente: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, responses.Message{Message: "Cliente criado com sucesso", Data: cliente})
}

// EditarCliente godoc
// @Summary Edita um cliente da loja virtual
// @Description Edita um cliente da loja virtual conforme o JSON e cpf_cnpj informados
// @Tags clientes
// @Produce json
// @Param cpf_cnpj path int true "CPF_CNPJ"
// @Param cliente body models.ClienteBase true "Dados do cliente"
// @Success 200 {object} responses.Message{data=models.Cliente}
// @Failure 400 {object} responses.Error
// @Failure 401 {object} responses.Error
// @Failure 409 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/clientes/editar/{cpf_cnpj} [put]
func EditarCliente(c *gin.Context) {
	var cliente models.Cliente
	var clienteBase models.ClienteBase

	cpfCnpj := c.Param("cpf_cnpj")

	if err := c.ShouldBindJSON(&clienteBase); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Não foi possível converter o JSON para o modelo de ClienteBase [" + err.Error() + "]"})
		return
	}

	if err := utils.ClienteDuplo(clienteBase.CpfCnpj, true, &clienteBase); err != nil {
		c.JSON(http.StatusConflict, responses.Error{Erro: err.Error()})
		return
	}

	if len(strconv.Itoa(clienteBase.CEP)) != 8 {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "O CEP não tem 8 dígitos"})
		return
	}

	if err := models.ValidaCliente(&clienteBase); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Erro na validação do cliente: " + err.Error()})
		return
	}

	if err := database.DB.Where("cpf_cnpj =?", cpfCnpj).First(&cliente).Error; err != nil {
		c.JSON(http.StatusNotFound, responses.Error{Erro: "Cliente não encontrado"})
		return
	}

	cliente.ClienteBase = clienteBase

	if err := database.DB.Save(&cliente).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Erro: "Erro durante a edição do cliente na base dados [" + err.Error() + "]"})
		return
	}

	c.JSON(http.StatusOK, responses.Message{Message: "Cliente editado com sucesso", Data: cliente})
}

// DeletarCliente godoc
// @Summary Deletar um cliente da loja virtual (soft-delete)
// @Description Deleta um cliente da loja virtual conforme cpf/cnpj informadas na URL
// @Tags clientes
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
