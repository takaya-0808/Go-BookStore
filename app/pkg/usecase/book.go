package usecase

import (
	"Go-BookStore/app/pkg/model"
)

type BookUsecase struct {
	bookrepo model.BookDomainRepository
}

func NewBookUsecase(bookrepo model.BookDomainRepository) *BookUsecase {
	bookUsecase := BookUsecase{bookrepo: bookrepo}
	return &bookUsecase
}

func (bu *BookUsecase) BookByID(id int) (books []*model.BOOK, err error) {
	books, err = bu.bookrepo.Find(id)
	return
}

func (bu *BookUsecase) FindAll() {}

func (bu *BookUsecase) Add(book *model.BOOK) (err error) {
	_, err = bu.bookrepo.Create(book)
	return err
}

func (bu *BookUsecase) Edit(id int) {}

func (bu *BookUsecase) Delete(id int) {}
