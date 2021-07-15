package controller

import (
	"net/http"
	//	"src/domain/model"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookUsecase usecase.bookUsecase
}

func (handler *BookHandler) View(c *gin.Context) {
	books, err := handler.bookUsecase.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, books)
		return
	}
	c.JSON(http.StatusOK, books)
	return
}

func (handler *BookHandler) Search(c *gin.Context) {}

func (handler *BookHandler) Create(c *gin.Context) {}

func (handler *BookHandler) Update(c *gin.Context) {}

func (handler *BookHandler) Delete(c *gin.Context) {}
