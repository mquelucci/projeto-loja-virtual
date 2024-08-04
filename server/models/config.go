package models

import "gorm.io/gorm"

type Config struct {
	gorm.Model
	NomeLoja string `json:"nomeLoja"`
}
