package model

type Cake struct {
	ID          int    `json:"id"`
	Title       string `json:"title" valid:"required~Title is required,type(string)~Title must be a string"`
	Description string `json:"description" valid:"required~Description is required,type(string)~Description must be a string"`
	Rating      int    `json:"rating" valid:"required~Rating is required,type(int)~Rating must be an integer,range(0|10)~Range must be between 0-10"`
	Image       string `json:"image" valid:"required~Image is required,type(string)~Image must be a string"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
