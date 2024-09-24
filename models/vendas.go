package models

import "gorm.io/gorm"

type VendaBase struct {
	ClienteID      int              `json:"cliente_id"`
	ItensVendaBase []ItensVendaBase `json:"itens_venda"`
}

type Venda struct {
	gorm.Model
	ClienteID  int          `json:"cliente_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	Cliente    Cliente      `gorm:"not null"`
	ValorTotal float64      `gorm:"not null"`
	ItensVenda []ItensVenda `json:"itens_venda" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
}
