package models

import "gorm.io/gorm"

type Venda struct {
	gorm.Model
	ClienteID  int          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	Cliente    Cliente      `gorm:"not null"`
	ValorTotal float64      `gorm:"not null"`
	ItensVenda []ItensVenda `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
}
