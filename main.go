package main

import "github.com/SilvanoP7/curriculum-api/routes"

func main() {
	router := routes.SetupRouter()

	router.Run(":8080")
}
