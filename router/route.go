package router

import (
	"rakamin-submission/controllers"
	"rakamin-submission/database"

	"github.com/gin-gonic/gin"
)

func StartRoute() *gin.Engine {
	route := gin.Default()

	route.GET("/users/allUser", controllers.GetAllUser)
	route.GET("/users/login", controllers.GetAllUser)
	route.POST("/users/:userId", controllers.Register)
	route.PUT("/users/:userId", controllers.UpdateUser)
	route.DELETE("/users/:userId", controllers.DeleteUser)
	route.POST("/photos", controllers.CreatePhoto)
	// route.GET("/photos", controllers.GetPhoto)
	// route.PUT("/photoId", controllers.GetPhoto)
	// route.DELETE("/:photoId", controllers.GetPhoto)

	database.StartDB()

	return route
}
