package common

import (
	"crud-go-api/libs/common/logger/middleware"
	"github.com/gin-gonic/gin"
)

func BaseRouter() *gin.Engine {
	router := gin.New()

	router.Use(middleware.SetUpLoggerMiddleware())

	router.Use(gin.Recovery())

	return router
}
