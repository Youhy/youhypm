package router

import (
	"github.com/gin-gonic/gin"
	"github.com/youhy/youhypm/models"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1") // localhost:5000/api/v1
	{
		users := main.Group("users")
		{
			users.POST("/", models.AddUser)
			users.GET("/:id", models.GetUser)
			users.GET("/", models.AllUsers)
			users.DELETE("/:id", models.Delete)
			users.PUT("/:id", models.UpdateUser)
		}
		books := main.Group("books")
		{
			books.GET("/:id", models.GetTask)
			books.GET("/", models.AllTasks)
			books.POST("/", models.AddTask)
			books.DELETE("/:id", models.Delete)
		}
	}
	return router
}
