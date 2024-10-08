package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type ProdutoBase struct {
	Descricao  string  `json:"descricao" validate:"nonzero, nonnil" gorm:"notNull"`
	Preco      float64 `json:"preco" validate:"min=0, nonnil" gorm:"notNull"`
	Quantidade int     `json:"quantidade" validate:"min=0, nonnil" gorm:"notNull"`
	Ativo      bool    `json:"ativo" validate:"nonnil" gorm:"notNull"`
}

type Produto struct {
	gorm.Model
	ProdutoBase
	Imagem string       `json:"imagem"`
	Vendas []ItensVenda `json:"-" gorm:"foreignKey:ProdutoID"`
}

func ValidaProduto(produto *Produto) error {
	if err := validator.Validate(produto); err != nil {
		return err
	}
	return nil
}
