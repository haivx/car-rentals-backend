package model

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
