package models

import "gorm.io/gorm"

type ItensVenda struct {
	gorm.Model `swaggerignore:"true"`
	VendaID    int     `swaggerignore:"true" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	ProdutoID  int     `json:"produto_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	Produto    Produto `swaggerignore:"true" gorm:"not null"`
	Quantidade int     `json:"quantidade" validate:"nonnil" gorm:"not null"`
	Preco      float64 `json:"preco" validate:"nonnull" gorm:"not null"`
}
