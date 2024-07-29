package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Produto struct {
	gorm.Model
	Descricao  string  `json:"descricao" validate:"nonzero, nonnil"`
	Preco      float64 `json:"preco" validade:"gte=0,  nonnil"`
	Quantidade int     `json:"quantidade" validate:"nonzero, nonnil"`
}

func ValidaProduto(produto *Produto) error {
	if err := validator.Validate(produto); err != nil {
		return err
	}
	return nil
}
