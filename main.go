package main

import (
	"health-system/src/initializers"
	"health-system/src/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// initializer used to init pre-server start
	initializers.Init()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(os.Getenv("PORT"))

}
