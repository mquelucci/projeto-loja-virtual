package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Nome  string
	Senha string
}
