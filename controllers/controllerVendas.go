package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/controllers/responses"
	"github.com/mquelucci/projeto-loja-virtual/database"
	"github.com/mquelucci/projeto-loja-virtual/models"
	"gorm.io/gorm"
)

type VendasSummary struct {
	gorm.Model
	ClienteID      uint
	CpfCnpjCliente string
	NomeCliente    string
	ValorTotal     float64
	ItensVenda     []struct {
		ProdutoID  uint
		Quantidade int
		Preco      float64
		Descricao  string
	}
}

// CriarVenda godoc
//
// @Summary		Criar uma venda
// @Description	Criar uma venda
// @Tags		vendas
// @Accept		json
// @Produce		json
// @Param		venda	body	models.VendaRequest	true	"Dados da venda"
// @Success		201	{object}	models.VendaRequest
// @Failure		401	{object}	responses.Error
// @Failure		404	{object}	responses.Error
// @Failure		422	{object}	responses.Error
// @Failure		500	{object}	responses.Error
// @Router		/admin/vendas/criar [post]
func CriarVenda(c *gin.Context) {
	var vendaJson models.VendaRequest

	if err := c.ShouldBindJSON(&vendaJson); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Erro ao interpretar o JSON - [" + err.Error() + "]"})
		return
	}

	//Agrupar produtos iguais e somar as quantidades
	produtoQuantidades := make(map[uint]int)
	for _, item := range vendaJson.Itens {
		produtoQuantidades[item.ProdutoID] += item.Quantidade
	}

	// Início de transação para garantir a consistência dos dados
	tx := database.DB.Begin()

	// Verificar a disponibilidade de estoque para cada produto
	for produtoID, quantidadeSolicitada := range produtoQuantidades {
		var produto models.Produto
		if err := tx.Where("id =?", produtoID).First(&produto).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusNotFound, responses.Error{Erro: "Erro ao buscar o produto com ID" + strconv.FormatUint(uint64(produtoID), 10) + "[" + err.Error() + "]"})
			return
		}

		if produto.Quantidade < quantidadeSolicitada {
			tx.Rollback()
			c.JSON(http.StatusUnprocessableEntity, responses.Error{Erro: "Quantidade insuficiente em estoque para o produto ID " + strconv.FormatUint(uint64(produtoID), 10)})
			return
		}

		// Atualizar a quantidade em estoque
		produto.Quantidade -= quantidadeSolicitada
		if err := tx.Save(&produto).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, responses.Error{Erro: "Erro ao tentar atualizar a quantidade em estoque do produto ID " + strconv.FormatUint(uint64(produtoID), 10) + "[" + err.Error() + "]"})
			return
		}
	}

	// Criar a venda
	venda := models.Venda{
		ClienteID: vendaJson.ClienteID,
	}

	// Inserir os itens na venda
	vendaItens := make([]models.ItensVenda, len(vendaJson.Itens))
	for i, item := range vendaJson.Itens {
		vendaItens[i] = models.ItensVenda{
			ProdutoID:  item.ProdutoID,
			Quantidade: item.Quantidade,
			Preco:      item.Preco,
		}
		venda.ValorTotal += (item.Preco * float64(item.Quantidade))
	}
	venda.Itens = vendaItens

	if err := tx.Create(&venda).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, responses.Error{Erro: "Erro ao tentar salvar a venda [" + err.Error() + "]"})
		return
	}

	// Confirmar a transação
	tx.Commit()

	c.JSON(http.StatusCreated, responses.Message{Message: "Venda criada com sucesso", Data: vendaJson})
}

// BuscarVendaPorId godoc
//
// @Summary		Busca uma venda por Id
// @Description	Busca uma venda por Id
// @Tags		vendas
// @Produce		json
// @Param		id	path	int	true	"ID da venda"
// @Success		200	{object}	VendasSummary
// @Failure		401	{object}	responses.Error
// @Failure		404	{object}	responses.Error
// @Router		/admin/vendas/buscar/{id} [get]
func BuscarVendaPorId(c *gin.Context) {
	var venda models.Venda
	var vendaSummary VendasSummary
	id := c.Param("id")

	if err := database.DB.Preload("Itens.Produto").Preload("Cliente").First(&venda, id).Error; err != nil {
		c.JSON(http.StatusNotFound, responses.Error{Erro: "Venda não encontrada [" + err.Error() + "]"})
		return
	}

	vendaSummary.ID = venda.ID
	vendaSummary.CreatedAt = venda.CreatedAt
	vendaSummary.UpdatedAt = venda.UpdatedAt
	vendaSummary.DeletedAt = venda.DeletedAt
	vendaSummary.ClienteID = venda.ClienteID
	vendaSummary.CpfCnpjCliente = venda.Cliente.CpfCnpj
	vendaSummary.NomeCliente = venda.Cliente.Nome
	vendaSummary.ValorTotal = venda.ValorTotal
	for _, item := range venda.Itens {
		vendaSummary.ItensVenda = append(vendaSummary.ItensVenda, struct {
			ProdutoID  uint
			Quantidade int
			Preco      float64
			Descricao  string
		}{
			item.ProdutoID,
			item.Quantidade,
			item.Preco,
			item.Produto.Descricao,
		})
	}

	c.JSON(http.StatusOK, responses.Message{Message: "Venda encontrada", Data: vendaSummary})
}
