package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/SilvanoP7/curriculum-api/models"
)

// Pong             godoc
// @Summary      Basic health check to ensure service is responding
// @Description  Responds with pong status.
// @Tags         health
// @Produce      json
// @Success      200 {object}   models.Pong
// @Router       /ping [get]
func Pong(c *gin.Context) {
        var pong models.Pong
        pong.Status = "pong"
	c.JSON(http.StatusOK, &pong)
}

// Pong             godoc
// @Summary      Basic health check to ensure service is responding
// @Description  Responds with pong status.
// @Tags         getSubjects
// @Produce      json
// @Success      200 {object}   models.Pong
// @Router       /getSubjects [get]
func GetSubjects(c *gin.Context) {
        var pong models.Pong
        pong.Status = "pong"
	c.JSON(http.StatusOK, &pong)       
}
