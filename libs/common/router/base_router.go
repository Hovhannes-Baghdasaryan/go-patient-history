package common

import (
	"github.com/Hovhannes-Baghdasaryan/go-patient-history/libs/common/logger/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func BaseRouter() *gin.Engine {
	router := gin.New()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(middleware.SetUpLoggerMiddleware())

	router.Use(gin.Recovery())

	return router
}
