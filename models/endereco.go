package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type EnderecoBase struct {
	Endereco string `json:"endereco" validate:"nonzero, nonnil" gorm:"not null"`
	Numero   int    `json:"numero" validate:"nonnil" gorm:"not null"`
	Bairro   string `json:"bairro" validate:"nonzero, nonnil" gorm:"not null"`
	CEP      int    `json:"cep" validate:"nonzero, nonnil" gorm:"not null"`
	Cidade   string `json:"cidade" validate:"nonzero, nonnil" gorm:"not null"`
	UF       string `json:"uf" validate:"nonzero, nonnil, len=2" gorm:"not null"`
}

type Endereco struct {
	gorm.Model
	ClienteID uint   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	Endereco  string `json:"endereco" validate:"nonzero, nonnil" gorm:"not null"`
	Numero    int    `json:"numero" validate:"nonnil" gorm:"not null"`
	Bairro    string `json:"bairro" validate:"nonzero, nonnil" gorm:"not null"`
	CEP       int    `json:"cep" validate:"nonzero, nonnil" gorm:"not null"`
	Cidade    string `json:"cidade" validate:"nonzero, nonnil" gorm:"not null"`
	UF        string `json:"uf" validate:"nonzero, nonnil, len=2" gorm:"not null"`
}

func ValidaEndereco(endereco *Endereco) error {
	if err := validator.Validate(endereco); err != nil {
		return err
	}
	return nil
}
