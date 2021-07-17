package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct{}

func (bc *BookController) View(c *gin.Context) {
	ID := c.Param("id")
	bookID, _ := strconv.Atoi(ID)
	c.JSON(http.StatusOK, gin.H{"book id": bookID})
}

func (bc *BookController) Search(c *gin.Context) {}
