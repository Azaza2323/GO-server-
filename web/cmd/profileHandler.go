package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a *application) GetUserProfile(c *gin.Context) {
	userIDParam := c.Param("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := a.users.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	books, err := a.userBooks.GetBooksByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "books": books})
}
