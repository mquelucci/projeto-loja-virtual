package utils

import (
	"github.com/mquelucci/projeto-loja-virtual/server/database"
	"github.com/mquelucci/projeto-loja-virtual/server/models"
)

func BuscarProdutos() []models.Produto {
	produtos := []models.Produto{}
	database.DB.Order("descricao ASC").Find(&produtos)
	return produtos
}
