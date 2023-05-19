package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/SilvanoP7/curriculum-api/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", handlers.Pong)
                v1.GET("/db", handlers.DbTest)
		v1.GET("/getSubjects", handlers.GetSubjects)
	}
	return router
}
