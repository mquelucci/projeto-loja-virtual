package models

import "gorm.io/gorm"

type Venda struct {
	gorm.Model `swaggerignore:"true"`
	ClienteID  int          `json:"cliente_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	Cliente    Cliente      `swaggerignore:"true" gorm:"not null"`
	ValorTotal float64      `swaggerignore:"true" gorm:"not null"`
	ItensVenda []ItensVenda `json:"itens_venda" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
}
