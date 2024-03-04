package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (a *application) GetUserProfile(c *gin.Context) {
	userID, _ := c.Get("id")

	id, err := strconv.Atoi(userID.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := a.users.GetByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve user profile"})
		return
	}

	response := struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Role  string `json:"role"`
	}{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, response)
}
