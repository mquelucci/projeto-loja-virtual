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

	noAuth := r.Group("/admin")
	{
		noAuth.POST("/autenticar", controllers.Autenticar)
	}

	auth := r.Group("/admin").Use(middlewares.Auth())
	{
		auth.POST("/configuracoes", controllers.AlterarConfiguracoes)
		auth.POST("/produtos/criar", controllers.CriarProduto)
		auth.POST("/produtos/editar", controllers.EditarProduto)
		auth.POST("/logout", controllers.FazerLogout)
		auth.DELETE("/produtos/removeImagem", controllers.RemoverImagemProduto)
		auth.DELETE("/produtos/delete", controllers.DeletarProduto)
		auth.GET("/produtos", controllers.BuscarTodosProdutos)

	}

	r.NoRoute(controllers.ExibeHTML404)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
