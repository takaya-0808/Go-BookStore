package app

import (
	"src/service/controller"

	"github.com/gin-gonic/gin"
)

func InitRouting(router *gin.Engine) {

	// bookContoroller := controller.NewBookRepository(NewSqlHandler())

	bookEngine := router.Group("/api")
	{
		v1 := bookEngine.Group("/v1")
		{
			v1.GET("/get_message/:name", func(c *gin.Context) { controller.GetMessage(c) })
			v1.GET("/read_message/:name", func(c *gin.Context) { controller.ReadMessage(c) })
			v1.GET("/country/:id", func(c *gin.Context) { controller.PrintID(c) })
			v1.GET("/Country/:id", func(c *gin.Context) { controller.GetCountry(c) })
			v1.POST("/Country/", func(c *gin.Context) { controller.PostCountry(c) })
		}
	}
	router.Run(":8018")
}
