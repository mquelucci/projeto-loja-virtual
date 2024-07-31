package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model
	Nome     string `json:"nome" validate:"nonzero, nonnil"`
	Telefone int    `json:"telefone" validade:"min=8, max=9, nonnil"`
	Email    string `json:"email" validade:"nonzero, nonnil"`
}

func ValidaCliente(cliente *Cliente) error {
	if err := validator.Validate(cliente); err != nil {
		return err
	}
	return nil
}
