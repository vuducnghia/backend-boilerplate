package routes

import (
	"backend-boilerplate/src/api/middleware"
	"backend-boilerplate/src/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter() *gin.Engine {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Boilerplate API"
	docs.SwaggerInfo.Description = "This is a boilerplate golang server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:9000"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "PATCH"},
		AllowHeaders:    []string{"Accept", "Accept-Encoding", "Authorization", "Content-Type", "Content-Length", "Origin", "X-CSRF-Token"},
	}))

	router.Use(middleware.Timeout)
	router.Use(middleware.ErrorHandler)

	// setup router groups
	NoAuthApi := router.Group("/api")

	addHeartbeat(NoAuthApi)
	addAuthRoutes(NoAuthApi)
	addUserRoutes(NoAuthApi)

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
