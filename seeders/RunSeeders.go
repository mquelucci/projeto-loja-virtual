package seeders

import (
	"github.com/mquelucci/projeto-loja-virtual/models"
	"gorm.io/gorm"
)

func RunSeeders(db *gorm.DB) error {
	var migration models.Migration
	if err := db.FirstOrCreate(&migration).Error; err != nil {
		return err
	}

	if !migration.FirstTime {
		if err := SeedAdmin(db); err != nil {
			return err
		}
		if err := SeedConfig(db); err != nil {
			return err
		}
		migration.FirstTime = true
		migration.ID = 1
		if err := db.Save(&migration).Error; err != nil {
			return err
		}
	}
	return nil
}
