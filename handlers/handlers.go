package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/SilvanoP7/curriculum-api/models"
)

// GetBooks responds with the list of all books as JSON.
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, '{"status":"pong"}')
}
