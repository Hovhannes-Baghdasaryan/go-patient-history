package exception

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (err *Error) BadRequestException(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, Error{
		Status:  http.StatusBadRequest,
		Message: err.Message,
	})
}

func (err *Error) NotFoundException(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, Error{
		Status:  http.StatusNotFound,
		Message: err.Message,
	})
}

func (err *Error) ForbiddenException(ctx *gin.Context) {
	ctx.JSON(http.StatusForbidden, Error{
		Status:  http.StatusForbidden,
		Message: err.Message,
	})
}

func (err *Error) NotAuthorizedException(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, Error{
		Status:  http.StatusUnauthorized,
		Message: err.Message,
	})
}

func (err *Error) InternalException(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, Error{
		Status:  http.StatusInternalServerError,
		Message: err.Message,
	})
}
