package controllers

import (
	"Go-BookStore/app/pkg/model"
	"Go-BookStore/app/pkg/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	controller usecase.Bookusecase
}

func NewBookController(controller usecase.Bookusecase) BookController {
	bu := BookController{controller: controller}
	return bu
}

func (bc *BookController) View(c *gin.Context) {
	ID := c.Param("id")
	bookID, _ := strconv.Atoi(ID)
	model, err := bc.controller.BookByID(bookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model)
	}
	c.JSON(http.StatusOK, model)
}

func (bc *BookController) Search(c *gin.Context) {}

func (bc *BookController) Add(c *gin.Context) {
	var books model.BOOK
	c.BindJSON(&books)
	err := bc.controller.Add(&books)
	c.JSON(http.StatusOK, err)
}
