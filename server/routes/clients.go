package routes

import "p2p-server/controllers"
import "github.com/gin-gonic/gin"

func registerClientRoutes(server *gin.Engine){
	router := server.Group("/clients")

	router.GET("", controllers.GetAllClients)

	router.GET("/:id", controllers.GetClientById)

	router.POST("", controllers.RegisterClient)

	router.DELETE("/:id", controllers.UnregisterClient)
}