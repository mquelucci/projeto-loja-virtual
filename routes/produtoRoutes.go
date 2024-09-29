package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/controllers"
	"github.com/mquelucci/projeto-loja-virtual/middlewares"
)

func ProdutoRoutes(router *gin.Engine) {
	authAdminProdutos := router.Group("/admin/produtos").Use(middlewares.Auth())
	{
		authAdminProdutos.GET("/todos", controllers.BuscarTodosProdutos)
		authAdminProdutos.GET("/:id", controllers.BuscarProdutoPorId)
		authAdminProdutos.POST("/criar", controllers.CriarProduto)
		authAdminProdutos.PUT("/editar/:id", controllers.EditarProduto)
		authAdminProdutos.PUT("/adicionaImagem/:id", controllers.AdicionarImagemProduto)
		authAdminProdutos.DELETE("/removeImagem/:id", controllers.RemoverImagemProduto)
		authAdminProdutos.DELETE("/deletar", controllers.DeletarProduto)
	}
}
