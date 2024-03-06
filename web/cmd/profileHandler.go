package main

import (
	"fmt"
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
func (a *application) AddBookToUser(c *gin.Context) {
	userIdS := c.Param("id")
	userId, err := strconv.Atoi(userIdS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var requestBody struct {
		BookID   int    `json:"bookId"`
		Rated    int    `json:"rated"`
		Feedback string `json:"feedback"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(requestBody.Rated, requestBody.Feedback)
	err = a.userBooks.Add(userId, requestBody.BookID, requestBody.Rated, requestBody.Feedback)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book added to user successfully"})
}

func (a *application) UpdateFeedback(c *gin.Context) {
	bookIdS := c.Param("id")
	bookId, err := strconv.Atoi(bookIdS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format"})
		return
	}
	authHeader := c.GetHeader("Authorization")
	tokenStr := ""
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenStr = authHeader[7:]
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header format must be Bearer {token}"})
		return
	}

	userIDStr, err := ExtractUserIDFromToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse userID"})
		return
	}

	var requestBody struct {
		Rated    int    `json:"rated"`
		Feedback string `json:"feedback"`
	}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = a.userBooks.Update(bookId, userID, requestBody.Rated, requestBody.Feedback)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update feedback"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Feedback updated successfully"})
}
