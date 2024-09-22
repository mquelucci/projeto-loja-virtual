package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/controllers/responses"
	"github.com/mquelucci/projeto-loja-virtual/models"
)

// CriarVenda godoc
//
// @Summary		Criar uma venda
// @Description	Criar uma venda
// @Tags		vendas
// @Accept		json
// @Produce		json
// @Param		venda	body	models.Venda	true	"Dados da venda"
// @Success		201	{object}	models.Venda
// @Failure		401	{object}	responses.Error
// @Router		/admin/vendas/criar [post]
func CriarVenda(c *gin.Context) {
	var venda models.Venda

	c.ShouldBindJSON(&venda)

	c.JSON(http.StatusCreated, responses.Message{Message: "Venda criada com sucesso", Data: venda})
}
