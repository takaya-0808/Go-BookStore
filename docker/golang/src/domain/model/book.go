package model

type Books []Book

type Book struct {
	ID     int    `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type User struct {
	ID    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

type CreateBook struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type UpdateBook struct {
	ID     int    `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type msg struct {
	Name    string `json:"user"`
	Message string
	Number  int
}

type Country struct {
	Code string
	Name string
}

type UserModel struct {
	UserName string
	PassWord string
	Email    string
}
