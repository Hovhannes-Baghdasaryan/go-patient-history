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

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,def"`
	Data    interface{} `json:"data,omitempty"`
}

func (r *Response) ActionSucceeded(ctx *gin.Context) {
	message := r.Message
	if message == "" {
		message = "Action Succeeded"
	}

	status := r.Status
	if status == 0 {
		status = http.StatusOK
	}

	webResponse := Response{
		Status:  status,
		Message: message,
		Data:    r.Data,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, webResponse)
}

func (r *Response) ActionPaginatedSucceeded(ctx *gin.Context) {
	message := r.Message
	if message == "" {
		message = "Action Succeeded"
	}

	status := r.Status
	if status == 0 {
		status = http.StatusOK
	}

	webResponse := Response{
		Status:  status,
		Message: message,
		Data:    r.Data,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, webResponse)
}
