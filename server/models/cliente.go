package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type ClienteBase struct {
	Nome     string `json:"nome" validate:"nonzero, nonnil" gorm:"unique;notNull"`
	Telefone int    `json:"telefone" validade:"min=8, max=9, nonnil" gorm:"notNull"`
	Email    string `json:"email" validade:"nonzero, nonnil" gorm:"notNull"`
}

type Cliente struct {
	gorm.Model
	ClienteBase
}

func ValidaCliente(cliente *Cliente) error {
	if err := validator.Validate(cliente); err != nil {
		return err
	}
	return nil
}
