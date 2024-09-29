package utils

import (
	"errors"

	"github.com/mquelucci/projeto-loja-virtual/database"
	"github.com/mquelucci/projeto-loja-virtual/models"
)

func ClienteDuplo(cpfCnpj string, editando bool, cliente *models.ClienteBase) error {
	var clientes []models.Cliente

	if editando && cpfCnpj == cliente.CpfCnpj {
		return nil
	} else {
		err := database.DB.Where("cpf_cnpj = ?", cpfCnpj).Find(&clientes).Error
		if err != nil {
			return err
		}
		if len(clientes) > 0 {
			return errors.New("Já existe um cliente com esse CPF/CNPJ " + cpfCnpj)
		}
		return nil
	}
}
