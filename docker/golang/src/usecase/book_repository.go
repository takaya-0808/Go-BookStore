package usecase

import (
	"src/domain/model"
)

type BookRepository interface {
	FindAll() (model.Books, error)
	FindById(id int) (model.Book, error)
	Add(book *model.Book) (int, error)
	Update(book *model.Book) (int, error)
	Delete(id int) (int, error)
}
