package middleware

import (
	logger "crud-go-api/libs/common/logger/main"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func SetUpLoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		log := logger.SetupLogger()

		infoData := fmt.Sprintf("Status=%#v IP=%#v - [%#v] Method=%#v Path=%#v Proto=%#v Latency=%#v UserAgent=%#v \n",
			param.StatusCode,
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.Latency,
			param.Request.UserAgent(),
		)

		if param.StatusCode >= 400 {
			log.Error(infoData)
		} else {
			log.Info(infoData)
		}

		return ""
	})
}
