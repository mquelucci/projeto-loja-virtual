package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/controllers"
	"github.com/mquelucci/projeto-loja-virtual/middlewares"
)

func ClienteRoutes(router *gin.Engine) {
	authAdminClientes := router.Group("/admin/clientes").Use(middlewares.Auth())
	{

		authAdminClientes.GET("/todos", controllers.BuscarTodosClientes)
		authAdminClientes.GET("/:cpf_cnpj", controllers.BuscarCliente)
		authAdminClientes.POST("/criar", controllers.CriarCliente)
		authAdminClientes.PUT("/editar/:cpf_cnpj", controllers.EditarCliente)
		authAdminClientes.DELETE("/deletar/:cpf_cnpj", controllers.DeletarCliente)

	}
}
