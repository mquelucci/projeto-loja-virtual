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
// @Param configuracoes body models.ConfigBase true "Dados da loja virtual"
// @Success 200 {object} responses.Message{data=models.Config}
// @Failure 400 {object} responses.Error
// @Failure 401 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/configuracoes [post]
func AlterarConfiguracoes(c *gin.Context) {
	var configs models.Config
	database.DB.First(&configs)
	c.ShouldBindJSON(&configs.ConfigBase)

	if len(strconv.Itoa(configs.ConfigBase.CEP)) != 8 {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "O CEP informado é inválido"})
		return
	}

	if err := models.ValidaConfiguracoes(&configs); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Erro na validação das configurações: " + err.Error()})
		return
	}

	tx := database.DB.Begin()
	if err := tx.Save(&configs).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, responses.Error{Erro: "Erro ao tentar salvar as configurações: " + err.Error()})
		return
	}
	tx.Commit()
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
