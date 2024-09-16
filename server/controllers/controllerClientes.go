package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/controllers/responses"
	"github.com/mquelucci/projeto-loja-virtual/server/database"
	"github.com/mquelucci/projeto-loja-virtual/server/models"
)

func BuscarClientes(c *gin.Context) {
	var clientes []models.Cliente
	database.DB.Preload("Endereco").Find(&clientes)
	c.JSON(http.StatusOK, responses.Message{Message: "Clientes encontrados", Data: clientes})
}

func CriarCliente(c *gin.Context) {
	var cliente models.Cliente

	nomeCliente := c.PostForm("nome")
	telefone := c.PostForm("telefone")
	email := c.PostForm("email")

	endereco := c.PostForm("endereco")

	numero := c.PostForm("numero")
	numeroConvertido, err := strconv.Atoi(numero)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Não foi possível converter o número informado" + err.Error()})
		return
	}

	bairro := c.PostForm("bairro")

	cep := c.PostForm("cep")
	cepConvertido, err := strconv.Atoi(cep)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Não foi possível converter o CEP informado" + err.Error()})
		return
	}

	cidade := c.PostForm("cidade")
	uf := c.PostForm("uf")

	cliente.Nome = nomeCliente
	cliente.Telefone = telefone
	cliente.Email = email
	cliente.Endereco = append([]models.Endereco{}, models.Endereco{Endereco: endereco, Numero: numeroConvertido, Bairro: bairro, CEP: cepConvertido, Cidade: cidade, UF: uf})

	if err := models.ValidaCliente(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Erro: "Erro na validação do cliente: " + err.Error()})
		return
	}

	if err := database.DB.Create(&cliente).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Erro: "Erro na criação do cliente: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, responses.Message{Message: "Cliente criado com sucesso", Data: cliente})
}
