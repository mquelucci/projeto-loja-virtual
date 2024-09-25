package seeders

import (
	"github.com/mquelucci/projeto-loja-virtual/models"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) error {
	admin := models.Admin{AdminBase: models.AdminBase{Nome: "admin", Senha: "admin"}}
	if err := db.Create(&admin).Error; err != nil {
		return err
	}
	return nil
}
