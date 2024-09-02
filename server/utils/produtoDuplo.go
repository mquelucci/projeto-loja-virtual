package utils

import (
	"errors"

	"github.com/mquelucci/projeto-loja-virtual/server/database"
	"github.com/mquelucci/projeto-loja-virtual/server/models"
)

func ProdutoDuplo(descricao string, editando bool, produto *models.Produto) error {
	var produtos []models.Produto

	if editando {
		if descricao != produto.Descricao {
			err := database.DB.Where("descricao = ?", descricao).Find(&produtos).Error
			if err != nil {
				return err
			}
			if len(produtos) > 0 {
				return errors.New("Já existe um produto com essa descrição")
			}
		}
		return nil
	} else {
		err := database.DB.Where("descricao = ?", descricao).Find(&produtos).Error
		if err != nil {
			return err
		}
		if len(produtos) > 0 {
			return errors.New("Já existe um produto com essa descrição")
		}
		return nil
	}
}
