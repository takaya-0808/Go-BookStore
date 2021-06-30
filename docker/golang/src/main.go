package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		v1.GET("/get_message/:name", func(c *gin.Context) {
			name := c.Param("name")
			message := "name is " + name
			c.JSON(http.StatusOK, gin.H{
				"name": message,
			})
		})
		// v1.GET("/read_message/:name", func(c *gin.Context){})
	}
	router.Run(":8018")
}
