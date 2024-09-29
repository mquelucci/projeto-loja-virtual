package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/mquelucci/projeto-loja-virtual/controllers"
	docs "github.com/mquelucci/projeto-loja-virtual/docs"
	"github.com/mquelucci/projeto-loja-virtual/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	r *gin.Engine
)

// @title Loja Virtual
// @version 1.0
// @description API para aplicação de loja virtual simples
// @host localhost:8080
// @BasePath /
// @schemes http

func HandleRequests() {
	gin.SetMode(gin.DebugMode)
	r = gin.Default()

	r.MaxMultipartMemory = 8 << 20 //Define que o tamanho máximo de upload do arquivo é 8 * 2^20 bytes (ou 8388608 bytes, ou 8 megabytes)

	store := memstore.NewStore([]byte("lojavirtual"))
	r.Use(sessions.Sessions("lojavirtual", store))
	r.Use(cors.Default())

	routes.AdminRoutes(r)
	routes.ProdutoRoutes(r)
	routes.ClienteRoutes(r)
	routes.VendaRoutes(r)

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.NoRoute(controllers.NotFound)
	r.Run(":80")
}
