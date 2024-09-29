package models

import "gorm.io/gorm"

type AdminBase struct {
	Nome  string `json:"nome"`
	Senha string `json:"senha"`
}

type Admin struct {
	gorm.Model `json="-" swaggerignore:"true"`
	AdminBase
}
