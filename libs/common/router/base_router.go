package common

import (
	"github.com/gin-gonic/gin"
	"go-patient-history/libs/common/logger/middleware"
)

func BaseRouter() *gin.Engine {
	router := gin.New()

	router.Use(middleware.SetUpLoggerMiddleware())

	router.Use(gin.Recovery())

	return router
}
