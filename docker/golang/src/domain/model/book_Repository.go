package model

import "time"

type Book struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Context   string    `json:"content"`
	Author    string    `json:"author"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// bookrepository
type bookRepository interface {
	FindAll() (book []*Book, err error)
	Find(id int) (book []*Book, err error)
	Create(book *Book) (*Book, error)
	Edit(id int) (*Book, error)
	Delete(id int) (*Book, error)
}
