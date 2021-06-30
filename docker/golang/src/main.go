package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Start server ...................")
	r := gin.Default()
	// app.InitRouting(r)
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("/v1")
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
	r.Run(":8018")
}
