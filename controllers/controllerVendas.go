package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/controllers/responses"
	"github.com/mquelucci/projeto-loja-virtual/database"
	"github.com/mquelucci/projeto-loja-virtual/models"
)

// CriarVenda godoc
//
// @Summary		Criar uma venda
// @Description	Criar uma venda
// @Tags		vendas
// @Accept		json
// @Produce		json
// @Param		venda	body	models.VendaBase	true	"Dados da venda"
// @Success		201	{object}	models.Venda
// @Failure		401	{object}	responses.Error
// @Router		/admin/vendas/criar [post]
func CriarVenda(c *gin.Context) {
	var venda models.Venda
	type listaDeProdutos struct {
		produtoId       int
		indexItensVenda int
		quantidade      int
	}
	var sliceListaDeProdutos []listaDeProdutos

	c.ShouldBindJSON(&venda)

	for index, item := range venda.ItensVenda {

		if item.Quantidade <= 0 {
			c.JSON(http.StatusBadRequest, responses.Error{Erro: "Quantidade do item nr. " + strconv.Itoa(index+1) + " não pode ser zero!"})
			return
		}

		venda.ValorTotal += item.Preco

		// Verifica se o produto informado na lista de itens vindos da requisição já havia sido informado antes
		// Se não, registra que ele apareceu e a quantidade informada
		// Se já, soma a quantidade informada a quantidade anterior
		for index, itemDaLista := range sliceListaDeProdutos {
			if itemDaLista.produtoId == item.ProdutoID {
				sliceListaDeProdutos[index].quantidade += item.Quantidade
				break
			} else {
				sliceListaDeProdutos = append(sliceListaDeProdutos, listaDeProdutos{item.ProdutoID, index, item.Quantidade})
				break
			}
		}

	}

	for _, itemDaLista := range sliceListaDeProdutos {

		if err := database.DB.First(&venda.ItensVenda[itemDaLista.indexItensVenda].Produto, itemDaLista.produtoId).Error; err != nil {
			c.JSON(http.StatusInternalServerError, responses.Error{Erro: err.Error()})
			return
		}

		if venda.ItensVenda[itemDaLista.indexItensVenda].Produto.Quantidade < itemDaLista.quantidade {
			c.JSON(http.StatusBadRequest, responses.Error{Erro: "Quantidade insuficiente no estoque do produto " + venda.ItensVenda[itemDaLista.indexItensVenda].Produto.Descricao + " - ID: " + strconv.FormatUint(uint64(venda.ItensVenda[itemDaLista.indexItensVenda].Produto.ID), 10)})
			return
		}

		venda.ItensVenda[itemDaLista.indexItensVenda].Produto.Quantidade -= itemDaLista.quantidade

		if err := database.DB.Save(&venda.ItensVenda[itemDaLista.indexItensVenda].Produto).Error; err != nil {
			c.JSON(http.StatusInternalServerError, responses.Error{Erro: err.Error()})
			return
		}
	}

	if err := database.DB.First(&venda.Cliente, venda.ClienteID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Erro: err.Error()})
		return
	}

	if err := database.DB.Create(&venda).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Erro: "Erro na criação da venda: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, responses.Message{Message: "Venda criada com sucesso", Data: venda})
}
