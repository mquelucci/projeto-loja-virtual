package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type ClienteBase struct {
	Nome     string     `json:"nome" validate:"nonzero, nonnil" gorm:"not null"`
	Telefone string     `json:"telefone" validade:"min=8, max=9, nonnil" gorm:"not null"`
	Email    string     `json:"email" validade:"nonzero, nonnil" gorm:"not null"`
	CpfCnpj  string     `json:"cpf_cnpj" validade:"nonzero, nonnil" gorm:"unique;not null"`
	Endereco []Endereco `json:"endereco" validate:"nonzero, nonnil" gorm:"not null"`
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
