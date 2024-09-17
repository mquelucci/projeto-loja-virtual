package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/server/controllers"
	docs "github.com/mquelucci/projeto-loja-virtual/server/docs"
	"github.com/mquelucci/projeto-loja-virtual/server/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Loja Virtual
// @version 1.0
// @description API para aplicação de loja virtual simples
// @host localhost:8080
// @BasePath /
// @schemes http

func HandleRequests() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.MaxMultipartMemory = 8 << 20

	store := cookie.NewStore([]byte("lojavirtual"))
	r.Use(sessions.Sessions("lojavirtual", store))

	// clientes := r.Group("/cliente") {

	// }

	noAuthAdmin := r.Group("/admin")
	{
		noAuthAdmin.POST("/autenticar", controllers.Autenticar)
	}

	authAdmin := r.Group("/admin").Use(middlewares.Auth())
	{
		authAdmin.POST("/configuracoes", controllers.AlterarConfiguracoes)
		authAdmin.POST("/produtos/criar", controllers.CriarProduto)
		authAdmin.POST("/produtos/editar", controllers.EditarProduto)
		authAdmin.POST("/logout", controllers.FazerLogout)
		authAdmin.POST("/clientes/criar", controllers.CriarCliente)

		authAdmin.DELETE("/produtos/removeImagem", controllers.RemoverImagemProduto)
		authAdmin.DELETE("/produtos/deletar", controllers.DeletarProduto)

		authAdmin.GET("/produtos", controllers.BuscarTodosProdutos)
		authAdmin.GET("/configuracoes", controllers.BuscarConfiguracoes)
		authAdmin.GET("/clientes/todos", controllers.BuscarTodosClientes)
		authAdmin.GET("/clientes/:cpf_cnpj", controllers.BuscarCliente)

	}

	r.NoRoute(controllers.NotFound)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
