package controller

import (
	"log"
	"net/http"
	"src/domain/model"
	"src/usecase"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookUsecase usecase.BookRepository
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

func GetMessage(c *gin.Context) {
	name := c.Param("name")
	message := "name is " + name
	c.JSON(http.StatusOK, gin.H{"name": message})
}

func ReadMessage(c *gin.Context) {
	var msg model.Msg
	msg.Name = c.Param("name")
	msg.Message = "これは構造体をJSONで返すためのテストです。"
	msg.Number = 1111
	c.JSON(http.StatusOK, msg)
}

func PrintID(c *gin.Context) {
	id := c.Param("id")
	log.Println(id)
}

func GetCountry(c *gin.Context) {
	var country model.Country
	id := c.Param("id")
	country.Code = id
	country.Name = "japan"
	c.JSON(http.StatusOK, country)
}

func PostCountry(c *gin.Context) {
	var country model.Country
	if err := c.BindJSON(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	log.Println(country.Code)
	c.JSON(http.StatusCreated, gin.H{"mesaage": "ok"})
}

func PostLogin(c *gin.Context) {
	var usermodel model.UserModel
	if c.BindJSON(&usermodel) == nil {
		if usermodel.UserName == "hoge" && usermodel.PassWord == "password" {
			c.JSON(http.StatusOK, gin.H{
				"status": "you are logged in",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
		}
	}
}
