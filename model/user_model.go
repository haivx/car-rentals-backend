package model

import (
	"time"
)

type User struct {
	Id        string    `pg:"id,pk" json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Password  string    `json:"-"`
	Token     string    `json:"-"` // @TODO: need to remove
}

type CreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
