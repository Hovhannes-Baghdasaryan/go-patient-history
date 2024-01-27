package common

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-patient-history/libs/common/logger/middleware"
)

func BaseRouter() *gin.Engine {
	router := gin.New()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(middleware.SetUpLoggerMiddleware())

	router.Use(gin.Recovery())

	return router
}
