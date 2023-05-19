package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/SilvanoP7/curriculum-api/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", handlers.Pong)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/getSubjects", handlers.getSubjects)
	}
	return router
}
