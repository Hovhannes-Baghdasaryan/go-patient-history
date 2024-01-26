package repository

import (
	"github.com/gin-gonic/gin"
	"go-patient-history/libs/common/response"
)

type BaseRepository interface {
	FindPaginated(ctx *gin.Context, page int, perPage int) (PaginatedOutputResponse[interface{}], error)
}

type PaginatedOutputResponse[T interface{}] struct {
	Pagination response.Pagination `json:"pagination"`
	Items      T                   `json:"items"`
}
