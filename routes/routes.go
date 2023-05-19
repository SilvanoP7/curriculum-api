package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/SilvanoP7/curriculum-api/handlers"
        "github.com/SilvanoP7/curriculum-api/middleware"
	"github.com/zsais/go-gin-prometheus"


)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(middleware.DefaultStructuredLogger()) // adds our new middleware
	router.Use(gin.Recovery())                       // adds the default recovery middleware
	p := ginprometheus.NewPrometheus("gin")
        p.Use(router)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", handlers.Pong)
                v1.GET("/db", handlers.DbTest)
		v1.GET("/getSubjects", handlers.GetSubjects)
	}
	return router
}
