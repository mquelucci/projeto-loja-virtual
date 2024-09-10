package models

import "gorm.io/gorm"

type AdminBase struct {
	Nome  string
	Senha string
}

type Admin struct {
	gorm.Model `swaggerignore:"true"`
	AdminBase
}
