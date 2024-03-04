package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a *application) GetUserProfile(c *gin.Context) {
	// Извлекаем ID пользователя из параметра пути
	userIDParam := c.Param("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Получаем данные пользователя
	user, err := a.users.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Получаем список книг, связанных с пользователем
	books, err := a.userBooks.GetBooksByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Отправляем профиль пользователя и список книг в ответе
	c.JSON(http.StatusOK, gin.H{"user": user, "books": books})
}
