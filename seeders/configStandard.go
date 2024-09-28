package seeders

import (
	"github.com/mquelucci/projeto-loja-virtual/models"
	"gorm.io/gorm"
)

func SeedConfig(db *gorm.DB) error {
	config := models.Config{ConfigBase: models.ConfigBase{
		NomeLoja: "Loja Virtual",
		Endereco: "",
		Numero:   0,
		Bairro:   "",
		CEP:      0,
		Cidade:   "",
		UF:       ""}}

	if err := db.Create(&config).Error; err != nil {
		return err
	}
	return nil
}
