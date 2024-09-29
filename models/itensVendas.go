package models

import "gorm.io/gorm"

type ItensVenda struct {
	gorm.Model
	VendaID    uint
	Venda      Venda `gorm:"foreignKey:VendaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	ProdutoID  uint
	Produto    Produto `gorm:"foreignKey:ProdutoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	Quantidade int     `validate:"nonnil"`
	Preco      float64 `validate:"nonnil"`
}
