package main

import (
	"log"
	"bookstore/app"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Start server ...................")
	r := gin.Default()
	app.InitRouting(r)
}