package mysql

import (
	"errors"
)

var (
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
type Book struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Image       string `json:"image"`
	Rated       int    `json:"rated,omitempty"`
	Feedback    string `json:"feedback,omitempty"`
}

type UserBooks struct {
	UserID int `json:"user_id"`
	BookID int `json:"book_id"`
}
type UserBookDetail struct {
	Book
	Rated    int    `json:"rated"`
	Feedback string `json:"feedback"`
}
type ReviewResponse struct {
	Username string `json:"username"`
	Rated    int    `json:"rated"`
	Feedback string `json:"feedback"`
}
