package models

import (
	"gorm.io/gorm"
)

type VendaRequest struct {
	ClienteID uint `json:"cliente_id"`
	Itens     []struct {
		ProdutoID  uint    `json:"produto_id"`
		Quantidade int     `json:"quantidade"`
		Preco      float64 `json:"preco"`
	} `json:"itens"`
}

type Venda struct {
	gorm.Model
	ClienteID  uint
	Cliente    Cliente      `gorm:"foreignKey:ClienteID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	ValorTotal float64      `gorm:"not null"`
	Itens      []ItensVenda `gorm:"foreignKey:VendaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
}
