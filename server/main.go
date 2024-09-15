package main

import "github.com/gin-gonic/gin"
import "p2p-server/routes"

func main(){
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":3000")
}