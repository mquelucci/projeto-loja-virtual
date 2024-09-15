package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type ClienteBase struct {
	Nome     string `json:"nome" validate:"nonzero, nonnil" gorm:"unique;not null"`
	Telefone int    `json:"telefone" validade:"min=8, max=9, nonnil" gorm:"not null"`
	Email    string `json:"email" validade:"nonzero, nonnil" gorm:"not null"`
	Endereco string `json:"endereco" validate:"nonzero, nonnil" gorm:"not null"`
	Bairro   string `json:"bairro" validate:"nonzero, nonnil" gorm:"not null"`
	CEP      string `json:"cep" validate:"nonzero, nonnil" gorm:"not null"`
	Cidade   string `json:"cidade" validate:"nonzero, nonnil" gorm:"not null"`
	UF       string `json:"uf" validate:"nonzero, nonnil" gorm:"not null"`
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
