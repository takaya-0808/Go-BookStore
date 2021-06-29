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
		c.String(http.StatusOK, "heelo world")
	})
	r.Run(":8018")
}
