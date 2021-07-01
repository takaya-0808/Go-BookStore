package main

import (
	"log"
	"net/http"

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
		v1.GET("/get_message/:name", func(c *gin.Context) {
			name := c.Param("name")
			message := "name is " + name
			c.JSON(http.StatusOK, gin.H{
				"name": message,
			})
		})

		v1.GET("/read_message/:name", func(c *gin.Context) {
			var msg struct {
				Name    string `json:"user"`
				Message string
				Number  int
			}
			msg.Name = c.Param("name")
			msg.Message = "これは構造体をJSONで返すためのテストです。"
			msg.Number = 1111
			c.JSON(http.StatusOK, msg)
		})

		v1.GET("/country/:id", func(c *gin.Context) {
			id := c.Param("id")
			log.Println(id)
		})

		v1.GET("/Country", func(c *gin.Context) {
			var country Country
			id := c.Query("country")
			country.Code = id
			country.Name = "japan"
			c.JSON(http.StatusOK, country)
		})
	}
	router.Run(":8018")
}
