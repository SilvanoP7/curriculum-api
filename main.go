package main

import ( 
        _ "github.com/SilvanoP7/curriculum-api/docs"
	"github.com/SilvanoP7/curriculum-api/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           curriculum-api
// @version         1.0

func main() {
	router := routes.SetupRouter()

	router.Run(":8080")
}
