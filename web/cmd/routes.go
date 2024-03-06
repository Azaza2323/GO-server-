package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *application) routes() http.Handler {
	r := gin.Default()
	r.POST("/login", a.Login)
	r.POST("/register", a.Register)

	r.GET("/profile/:id", a.GetUserProfile)

	r.GET("/", a.getAllBooks)
	r.GET("/:id", a.getBookByID)
	r.POST("/create", a.InsertBook)

	r.GET("/users/:id", a.GetUser)
	r.DELETE("/users/:id", a.DeleteUser)

	r.POST("/admin/create", a.InsertBook)
	r.DELETE("/admin/delete/:id", a.DeleteBook)

	r.GET("/category/:category", a.getBooksByCategory)

	r.POST("/read/:bookID", a.AddBookToUser)

	return r
}
