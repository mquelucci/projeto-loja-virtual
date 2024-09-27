package models

import (
	"time"

	"gopkg.in/validator.v2"
)

type ConfigBase struct {
	NomeLoja string `json:"nomeLoja"`
	Endereco string `json:"endereco" validate:"nonzero, nonnil" gorm:"not null"`
	Numero   int    `json:"numero" validate:"nonzero, nonnil" gorm:"not null"`
	Bairro   string `json:"bairro" validate:"nonzero, nonnil" gorm:"not null"`
	CEP      int    `json:"cep" validate:"nonzero, nonnil" gorm:"not null"`
	Cidade   string `json:"cidade" validate:"nonzero, nonnil" gorm:"not null"`
	UF       string `json:"uf" validate:"nonzero, nonnil" gorm:"not null"`
}

type Config struct {
	ID        uint      `json:"-" gorm:"primary_key"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time
	ConfigBase
}

func ValidaConfiguracoes(config *Config) error {
	if err := validator.Validate(config); err != nil {
		return err
	}
	return nil
}
