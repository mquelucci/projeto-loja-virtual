package models

import (
	"gorm.io/gorm"
)

type TipoPag int

const (
	Credito TipoPag = iota + 1
	Debito
	Pix
	Dinheiro
)

func (tp TipoPag) String() string {
	return []string{"Credito", "Debito", "Pix", "Dinheiro"}[tp-1]
}

type VendaRequest struct {
	ClienteID uint `json:"cliente_id"`
	Itens     []struct {
		ProdutoID  uint    `json:"produto_id"`
		Quantidade int     `json:"quantidade"`
		Preco      float64 `json:"preco"`
	} `json:"itens"`
	SitPag  bool `json:"sitpag"`
	TipoPag TipoPag
}

type Venda struct {
	gorm.Model
	ClienteID  uint
	Cliente    Cliente      `gorm:"foreignKey:ClienteID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	ValorTotal float64      `gorm:"not null"`
	Itens      []ItensVenda `gorm:"foreignKey:VendaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
}
