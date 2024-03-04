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
	// The routes below were modified or added to resolve the conflict
	r.DELETE("/:id", a.deleteBookByID)   // Kept from HEAD
	r.GET("/users/:id", a.GetUser)       // Kept from HEAD
	r.DELETE("/users/:id", a.DeleteUser) // Kept from HEAD

	r.POST("/admin/create", a.InsertBook)       // Added from origin/azamat
	r.DELETE("/admin/delete/:id", a.DeleteBook) // Added from origin/azamat
	return r
}
