package model

type CreateOrUpdateNews struct {
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Image    string `json:"image"`
	Author   string `json:"author"`
	Content  string `json:"content"`
	Language string `json:"language" form:"language" binding:"required"`
}

type News struct {
	Id       string `pg:"id,pk" json:"id"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Image    string `json:"image"`
	Author   string `json:"author"`
	Content  string `json:"content"`
	Language string `json:"language"`
}

// Name     string `json:"name" form:"name" binding:"required"`
// Email    string `json:"email" form:"email" binding:"required,email" `
// Password string `json:"password" form:"password" binding:"required"`
