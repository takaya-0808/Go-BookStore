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
