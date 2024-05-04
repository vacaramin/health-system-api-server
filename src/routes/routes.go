package routes

import (
	"health-system/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Setting up nested Routes
	userRoutes := r.Group("/users")

	//Ping Route
	r.GET("/ping", controllers.PingHandler)

	//User Routes
	userRoutes.GET("/", controllers.GetUsers)

}
