package usecase

import (
	"Go-BookStore/app/pkg/model"
)

type Bookusecase interface {
	BookByID(id int) (books []*model.BOOK, err error)
	FindAll()
	Add(*model.BOOK) (err error)
	Edit(id int)
	Delete(id int)
}
