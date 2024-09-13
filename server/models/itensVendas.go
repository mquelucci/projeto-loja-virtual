package models

import "gorm.io/gorm"

type ItensVenda struct {
	gorm.Model
	VendaID   int     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	ProdutoID int     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	Produto   Produto `gorm:"not null"`
}
