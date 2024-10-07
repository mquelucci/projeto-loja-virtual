package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type ClienteBase struct {
	Nome     string  `json:"nome" validate:"nonzero, nonnil" gorm:"not null"`
	Empresa  *string `json:"empresa"`
	Telefone string  `json:"telefone" validade:"min=8, max=9, nonnil" gorm:"not null"`
	Email    string  `json:"email" validade:"nonzero, nonnil" gorm:"not null"`
	CpfCnpj  string  `json:"cpf_cnpj" validade:"nonzero, nonnil, min=11, max=14" gorm:"not null; index:idx_cpf_cnpj"`
	Endereco string  `json:"endereco" validate:"nonzero, nonnil" gorm:"not null"`
	Numero   int     `json:"numero" validate:"nonnil" gorm:"not null"`
	Bairro   string  `json:"bairro" validate:"nonzero, nonnil" gorm:"not null"`
	CEP      int     `json:"cep" validate:"nonzero, nonnil" gorm:"not null"`
	Cidade   string  `json:"cidade" validate:"nonzero, nonnil" gorm:"not null"`
	UF       string  `json:"uf" validate:"nonzero, nonnil, len=2" gorm:"not null"`
}

type Cliente struct {
	gorm.Model
	ClienteBase
	Vendas []Venda `json:"-" gorm:"foreignKey:ClienteID"`
}

func ValidaCliente(clienteBase *ClienteBase) error {
	if err := validator.Validate(clienteBase); err != nil {
		return err
	}
	return nil
}
