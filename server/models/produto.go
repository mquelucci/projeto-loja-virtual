package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Produto struct {
	gorm.Model
	Descricao  string  `json:"descricao" validate:"nonzero, nonnil"`
	Preco      float64 `json:"preco" validade:"nonnil"`
	Quantidade int     `json:"quantidade" validate:"nonnil"`
}

func ValidaProduto(produto *Produto) error {
	if err := validator.Validate(produto); err != nil {
		return err
	}
	return nil
}
