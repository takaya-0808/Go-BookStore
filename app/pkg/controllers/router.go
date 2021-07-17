package controllers

import (
	"github.com/gin-gonic/gin"
)

func SetUp(Router *gin.RouterGroup) {
	books := Router.Group("/book")
	{
		bc := BookController{}
		books.GET("", bc.View)
		books.GET("/:id", bc.Search)
		// books.POST("", bc.Add)
		// books.PUT("/:id")
		// books.DELETE("/:id")
	}
}
