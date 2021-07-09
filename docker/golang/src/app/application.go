package app

import (
	"src/service/controller"

	"github.com/gin-gonic/gin"
)

func InitRouting(router *gin.Engine) {

	// bookContoroller := controller.NewBookRepository(NewSqlHandler())

	handler := controller.BookHandler

	bookEngine := router.Group("/api")
	{
		v1 := bookEngine.Group("/v1")
		{
			v1.GET("/get_message/:name", func(c *gin.Context) { controller.GetMessage(c) })
			v1.GET("/read_message/:name", func(c *gin.Context) { controller.ReadMessage(c) })
			v1.GET("/country/:id", func(c *gin.Context) { controller.PrintID(c) })
			v1.GET("/Country/:id", func(c *gin.Context) { controller.GetCountry(c) })
			v1.POST("/Country/", func(c *gin.Context) { controller.PostCountry(c) })
			v1.POST("/loginJSON/", func(c *gin.Context) { controller.PostLogin(c) })
			v1.GET("/Book", func(c *gin.Context) { handler.View(c) })
			v1.GET("/Book/:id", func(c *gin.Context) { handler.Search(c) })
			v1.POST("/Book/:id", func(c *gin.Context) { handler.Create(c) })
			v1.PUT("/Book/:id", func(c *gin.Context) { handler.Update(c) })
			v1.DELETE("/Book/:id", func(c *gin.Context) { handler.Delete(c) })
		}
	}
	router.Run(":8018")
}
