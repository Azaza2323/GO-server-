package main

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	models "server/mysql"
	"strconv"
)

func (a *application) getAllBooks(c *gin.Context) {
	books, err := a.books.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (a *application) getBookByID(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	book, err := a.books.GetID(bookID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (a *application) deleteBookByID(c *gin.Context) {
	id := c.Query("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	err = a.books.Delete(bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (a *application) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	bookId, _ := strconv.Atoi(id)
	if err := a.books.Delete(bookId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (a *application) InsertBook(c *gin.Context) {
	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := a.books.Insert(book.Name, book.Author, book.Description, book.Category, book.Image); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
