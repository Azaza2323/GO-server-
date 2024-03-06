package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *application) routes() http.Handler {
	r := gin.Default()

	r.POST("/login", a.Login)
	r.POST("/register", a.Register)
	auth := r.Group("/").Use(isAuth())
	{
		auth.GET("/profile/:id", a.GetUserProfile)
		auth.GET("/reviews/:id", a.GetReviews)
		auth.GET("/", a.getAllBooks)
		auth.GET("/:id", a.getBookByID)
		auth.PUT("/profile/:id", a.UpdateFeedback)
		auth.POST("/add/:id", a.AddBookToUser)
		auth.POST("/create", isAdmin(), a.InsertBook)
		auth.DELETE("/:id", isAdmin(), a.deleteBookByID)
		auth.POST("/admin/create", isAdmin(), a.InsertBook)
		auth.DELETE("/admin/delete/:id", isAdmin(), a.DeleteBook)
		auth.GET("/category/:category", a.getBooksByCategory)
	}

	return r
}
