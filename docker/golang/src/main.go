package main

import (
	"log"
	"src/app"

	"github.com/gin-gonic/gin"
)

type Country struct {
	Code string
	Name string
}

func main() {
	log.Println("Start server ...................")
	router := gin.Default()
	app.InitRouting(router)
}
