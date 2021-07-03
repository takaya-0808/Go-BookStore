package main

import (
	"log"
	"net/http"
	"src/service/controller"

	"github.com/gin-gonic/gin"
)

type Country struct {
	Code string
	Name string
}

func main() {
	log.Println("Start server ...................")
	router := gin.Default()
	// app.InitRouting(r)
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1 := router.Group("/v1")
	{
		v1.GET("/get_message/:name", controller.GetMessage)

		v1.GET("/read_message/:name", controller.ReadMessage)

		v1.GET("/country/:id", controller.PrintID)

		v1.GET("/Country/:id", controller.GetCountry)

		v1.POST("/Country/", controller.PostCountry)
	}
	router.Run(":8018")
}
