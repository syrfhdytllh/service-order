package routers

import (
	"service-order/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	//Create
	router.POST("/orders", controllers.CreateOrders)
	// Read All
	router.GET("/orders", controllers.GetAllOrder)
	// Read
	router.GET("/orders/:id", controllers.GetOneOrders)
	// Update
	router.PUT("/orders/:id", controllers.UpdateOrders)
	// DELETE
	router.DELETE("/orders/:id", controllers.DeleteOrders)

	return router
}
