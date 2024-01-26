package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func InjectTagRouter(router *gin.Engine, clientDB *ent.Client) **gin.Engine {
	// Repo
	tagRepository := repository.InjectTagsRepositoryImpl(clientDB)

	// Service
	validate := validator.New()
	tagService := services.InjectTagsServiceImpl(tagRepository, validate)

	// Controller
	tagController := controller.InjectTagsController(tagService)

	tagsRouter := router.Group("v1/tags")
	{
		tagsRouter.GET("", tagController.FindAll)
		tagsRouter.GET("/:tagId", tagController.FindById)
		tagsRouter.POST("", tagController.Create)
		tagsRouter.PATCH("/:tagId", tagController.Update)
		tagsRouter.DELETE("/:tagId", tagController.Delete)
	}

	return &router
}
