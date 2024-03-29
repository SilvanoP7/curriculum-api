package main

import ( 
        _ "github.com/SilvanoP7/curriculum-api/docs"
	"github.com/SilvanoP7/curriculum-api/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           curriculum-api
// @version         1.0
// @BasePath  /api/v1

func main() {
	router := routes.SetupRouter()
        router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
