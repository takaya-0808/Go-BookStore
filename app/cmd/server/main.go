package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"Go-BookStore/app/pkg/connecter"
	"Go-BookStore/app/pkg/controllers"
)

func main() {

	fmt.Print("hello world")
	db := connecter.SetConnection()
	defer db.Close()
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"hoge": "fuga",
		})
	})
	r := router.Group("/api/v1")
	controllers.SetUp(r)
	router.Run(":8019")
}
