package controller

import (
	"log"
	"net/http"
	"src/domain/model"

	"github.com/gin-gonic/gin"
)

func GetMessage(c *gin.Context) {
	name := c.Param("name")
	message := "name is " + name
	c.JSON(http.StatusOK, gin.H{"name": message})
}

func ReadMessage(c *gin.Context) {
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
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

// type BookController struct {
// 	controller usecase.BookUsecase
// }

// func NewBookRepository(sqlHandler mysql.SqlHandler) *BookController {

// 	return &BookController {
// 		controller: usecase.BookUsecase {
// 			BookRepository: &mysql.BookRepository {
// 				SqlHandler: sqlHandler,
// 			},
// 		},
// 	}
// }

// func (control *BookController) Create(c Context) {

// 	b := model.Book{}
// 	c.Bind(&b)
// 	book, err := control.controller.Add(b)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, NewError(err))
// 		return
// 	}
// 	c.JSON(http.StatusCreated, book)
// }

// func (control *BookController) Index(c Context) {

// 	books, err := control.controller.View()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, NewError(err))
// 		return
// 	}
// 	c.JSON(http.StatusOK, books)
// }

// func (control *BookController) Show(c Context) {

// 	id, _ := strconv.Atoi(c.Param("id"))
// 	book, err := control.controller.Search(id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, NewError(err))
// 		return
// 	}
// 	c.JSON(http.StatusOK, book)
// }

// func (control *BookController) Renewal(c Context)  {}

// func (control *BookController) Remove(c Context) {}
