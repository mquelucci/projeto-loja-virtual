package utils

import (
	"errors"

	"github.com/mquelucci/projeto-loja-virtual/database"
	"github.com/mquelucci/projeto-loja-virtual/models"
)

func ProdutoDuplo(descricao string, editando bool, produto *models.ProdutoBase) error {
	var produtos []models.ProdutoBase

	if editando && descricao == produto.Descricao {
		return nil
	} else {
		err := database.DB.Where("descricao = ?", descricao).Find(&produtos).Error
		if err != nil {
			return err
		}
		if len(produtos) > 0 {
			return errors.New("Já existe um produto com essa descrição " + descricao)
		}
		return nil
	}
}
