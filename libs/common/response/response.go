package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Pagination struct {
	Total   int `json:"total"`
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

type Response[T interface{}] struct {
	Status  int    `json:"status"`
	Message string `json:"message,def"`
	Data    T      `json:"data,omitempty"`
}

type PaginatedOutputResponse[T interface{}] struct {
	Pagination `json:"pagination"`
	Items      T `json:"items"`
}

func (r *Response[T]) ActionSucceeded(ctx *gin.Context) {
	message := r.Message
	if message == "" {
		message = "Action Succeeded"
	}

	status := r.Status
	if status == 0 {
		status = http.StatusOK
	}

	webResponse := Response[T]{
		Status:  status,
		Message: message,
		Data:    r.Data,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, webResponse)
}
