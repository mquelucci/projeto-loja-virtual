package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/controllers/responses"
	"github.com/mquelucci/projeto-loja-virtual/database"
	"github.com/mquelucci/projeto-loja-virtual/models"
)

// Autenticar godoc
// @Summary Faz a autenticação do usuário
// @Description Através dos dados fornecidos via formulário HTML, compara com o banco de dados para autenticar ou rejeitar
// @Tags auth, admin
// @Accept mpfd
// @Produce json
// @Param usuario formData models.AdminBase true "Dados do usuário"
// @Success 200 {object} responses.Message
// @Failure 401 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /admin/autenticar [post]
func Autenticar(c *gin.Context) {
	usuario := c.PostForm("nome")
	senha := c.PostForm("senha")
	var admin models.Admin
	if err := database.DB.Where("nome = ? AND senha = ?", usuario, senha).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, responses.Error{Erro: err.Error()})
		return
	}
	session := sessions.Default(c)
	session.Set("auth", true)
	err := session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Erro: err.Error()})
		return
	}
	c.JSON(http.StatusOK, responses.Message{Message: "Usuário autenticado"})
}

// FazerLogout godoc
// @Summary Faz o logout do usuário
// @Description Remove a autenticação da sessão do usuário atual
// @Tags auth, admin
// @Produce json
// @Success 200 {object} responses.Message
// @Failure 500 {object} responses.Error
// @Router /admin/logout [post]
func FazerLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("auth")
	err := session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Erro: err.Error()})
		return
	}
	c.JSON(http.StatusOK, responses.Message{Message: "Usuário deslogado"})
}
