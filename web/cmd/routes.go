package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *application) routes() http.Handler {
	r := gin.Default()

	r.POST("/login", a.Login)
	r.POST("/register", a.Register)
	r.GET("/", a.getAllBooks)
	r.GET("/:id", a.getBookByID)
	r.DELETE("/:id", a.deleteBookByID)
	r.POST("/create", a.InsertBook)

	return r
}
