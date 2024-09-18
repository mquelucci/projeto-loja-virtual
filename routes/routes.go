package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/controllers"
	docs "github.com/mquelucci/projeto-loja-virtual/docs"
	"github.com/mquelucci/projeto-loja-virtual/middlewares"
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

	store := memstore.NewStore([]byte("lojavirtual"))
	r.Use(sessions.Sessions("lojavirtual", store))
	r.Use(cors.Default())

	// clientes := r.Group("/cliente") {

	// }

	noAuthAdmin := r.Group("/admin")
	{
		noAuthAdmin.POST("/autenticar", controllers.Autenticar)
	}

	authAdminProdutos := r.Group("/admin/produtos").Use(middlewares.Auth())
	{
		authAdminProdutos.GET("/produtos", controllers.BuscarTodosProdutos)
		authAdminProdutos.POST("/produtos/criar", controllers.CriarProduto)
		authAdminProdutos.PUT("/produtos/editar", controllers.EditarProduto)
		authAdminProdutos.DELETE("/produtos/removeImagem", controllers.RemoverImagemProduto)
		authAdminProdutos.DELETE("/produtos/deletar", controllers.DeletarProduto)
	}

	authAdminClientes := r.Group("/admin/clientes").Use(middlewares.Auth())
	{

		authAdminClientes.GET("/clientes/todos", controllers.BuscarTodosClientes)
		authAdminClientes.GET("/clientes/:cpf_cnpj", controllers.BuscarCliente)

		authAdminClientes.POST("/logout", controllers.FazerLogout)
		authAdminClientes.POST("/clientes/criar", controllers.CriarCliente)

	}

	authAdmin := r.Group("/admin").Use(middlewares.Auth())
	{
		authAdmin.GET("/configuracoes", controllers.BuscarConfiguracoes)
		authAdminClientes.POST("/configuracoes", controllers.AlterarConfiguracoes)
	}

	r.NoRoute(controllers.NotFound)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
