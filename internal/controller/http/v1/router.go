package v1

import (
	"github.com/gin-gonic/gin"
	_ "github.com/kostylevdev/todo-rest-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.SignUp)
		auth.POST("sign-in", h.SignIn)
	}
	api := router.Group("/api", h.UserIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.CreateList)
			lists.GET("/", h.GetAllLists)
			lists.GET("/:listId", h.GetListById)
			lists.PUT("/:listId", h.UpdateList)
			lists.DELETE("/:listId", h.DeleteList)

			items := lists.Group("/:listId/items")
			{
				items.POST("/", h.CreateItem)
				items.GET("/", h.GetAllItems)
				items.GET("/:itemId", h.GetItemById)
				items.PUT("/:itemId", h.UpdateItem)
				items.DELETE("/:itemId", h.DeleteItem)
			}
		}
	}
	return router
}
