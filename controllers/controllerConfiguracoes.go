package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/controllers/responses"
	"github.com/mquelucci/projeto-loja-virtual/database"
	"github.com/mquelucci/projeto-loja-virtual/models"
)

// AlterarConfiguracoes godoc
// @Summary Altera as configurações da loja virtual
// @Description Altera as configurações da loja virtual conforme informações enviadas pelo formulário
// @Tags admin
// @Produce json
// @Param configuracoes formData models.ConfigBase true "Dados da loja virtual"
// @Success 200 {object} responses.Message{data=models.Config}
// @Failure 400 {object} responses.Error
// @Failure 401 {object} responses.Error
// @Router /admin/configuracoes [post]
func AlterarConfiguracoes(c *gin.Context) {
	var configs models.Config
	database.DB.First(&configs)
	nomeLoja := c.PostForm("nomeLoja")
	endereco := c.PostForm("endereco")
	numero := c.PostForm("numero")
	bairro := c.PostForm("bairro")
	cep := c.PostForm("cep")
	cidade := c.PostForm("cidade")
	uf := c.PostForm("uf")
	configs.NomeLoja = nomeLoja
	configs.Endereco = endereco
	configs.Bairro = bairro

	cepConvertido, err := strconv.Atoi(cep)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Não foi possível converter o CEP informado!"})
	}
	configs.CEP = cepConvertido

	numeroConvertido, err := strconv.Atoi(numero)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Não foi possível converter o número informado!"})
	}
	configs.Numero = numeroConvertido

	configs.Cidade = cidade
	configs.UF = uf

	if err := models.ValidaConfiguracoes(&configs); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Erro na validação das configurações: " + err.Error()})
		return
	}

	database.DB.Save(&configs)
	c.JSON(http.StatusOK, responses.Message{Message: "Configurações alteradas com sucesso", Data: configs})
}

// BuscarConfiguracoes godoc
// @Summary Busca as configurações da loja virtual
// @Description Busca as configurações da loja virtual e retorna no JSON
// @Tags admin
// @Produce json
// @Success 200 {object} responses.Message{data=models.Config}
// @Failure 401 {object} responses.Error
// @Router /admin/configuracoes [get]
func BuscarConfiguracoes(c *gin.Context) {
	var configs models.Config
	database.DB.Find(&configs)
	c.JSON(http.StatusOK, responses.Message{Message: "Configurações encontradas", Data: configs})
}
