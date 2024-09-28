package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type ConfigBase struct {
	NomeLoja string `json:"nomeLoja" validate:"nonzero, nonnil" gorm:"not null"`
	Endereco string `json:"endereco" validate:"nonzero, nonnil" gorm:"not null"`
	Numero   int    `json:"numero" validate:"nonnil" gorm:"not null"`
	Bairro   string `json:"bairro" validate:"nonzero, nonnil" gorm:"not null"`
	CEP      int    `json:"cep" validate:"nonzero, nonnil" gorm:"not null"`
	Cidade   string `json:"cidade" validate:"nonzero, nonnil" gorm:"not null"`
	UF       string `json:"uf" validate:"nonzero, nonnil, len=2" gorm:"not null"`
}

type Config struct {
	gorm.Model `json:"-"`
	ConfigBase
}

func ValidaConfiguracoes(config *Config) error {
	if err := validator.Validate(config); err != nil {
		return err
	}
	return nil
}
