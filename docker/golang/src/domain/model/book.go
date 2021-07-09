package model

import (
	"time"
)

type Books []Book

type Book struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Context   string    `json:"content"`
	Author    string    `json:"author"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Msg struct {
	Name    string `json:"user"`
	Message string
	Number  int
}

type Country struct {
	Code string
	Name string
}

type UserModel struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Email    string `json:"email"`
}
