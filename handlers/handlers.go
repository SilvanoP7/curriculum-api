package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/SilvanoP7/curriculum-api/models"
)

// GetBooks responds with the list of all books as JSON.
// Pong             godoc
// @Summary      Basic health check to ensure service is respnding
// @Description  Responds with pong status.
// @Tags         hralth
// @Produce      json
// @Success      200  {array}  models.Pong
// @Router       /ping [get]
func Pong(c *gin.Context) {
        var pong models.Pong
        pong.Status = "pong"
	c.JSON(http.StatusOK, &pong)
}
