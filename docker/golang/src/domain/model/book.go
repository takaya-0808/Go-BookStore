package model

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
