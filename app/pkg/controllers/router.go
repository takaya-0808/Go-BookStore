package controllers

import (
	"github.com/gin-gonic/gin"
)

func SetUp(Router *gin.RouterGroup) {
	books := Router.Group("/book")
	{
		bc := BookController{}
		books.GET("", bc.Search)
		books.GET("/:id", bc.View)
		// books.POST("")
		// books.PUT("/:id")
		// books.DELETE("/:id")
	}
}
