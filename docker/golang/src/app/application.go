package app

import (
	"src/service/controller"

	"github.com/gin-gonic/gin"
)

func InitRouting(router *gin.Engine) {

	// bookContoroller := controller.NewBookRepository(NewSqlHandler())

	handler := controller.BookHandler{}

	bookEngine := router.Group("/api")
	{
		v1 := bookEngine.Group("/v1")
		{
			v1.GET("/Book", func(c *gin.Context) { handler.View(c) })
			v1.GET("/Book/:id", func(c *gin.Context) { handler.Search(c) })
			v1.POST("/Book/:id", func(c *gin.Context) { handler.Create(c) })
			v1.PUT("/Book/:id", func(c *gin.Context) { handler.Update(c) })
			v1.DELETE("/Book/:id", func(c *gin.Context) { handler.Delete(c) })
		}
	}
	router.Run(":8018")
}
