package controllers

import (
	"fmt"
	"p2p-server/models"

	"github.com/gin-gonic/gin"
)
var(
	allClients = []models.Client{}
)

func GetAllClients(context *gin.Context){
	context.IndentedJSON(200, allClients)
}

func GetClientById(context *gin.Context){
	id := context.Params.ByName("id")

	for _, client := range(allClients){
		if client.ID==id{
			context.IndentedJSON(200, client)
			return
		}
	}

	context.JSON(404, gin.H{"msg": "Client not found!"})
}

func RegisterClient(context *gin.Context){
	client := models.Client{}
	err := context.ShouldBindJSON(&client)

	if err!=nil{
		context.JSON(400, gin.H{"msg": err})
		return
	}
	fmt.Println(client)

	for i, curr_client := range(allClients){
		if curr_client.ID == client.ID{
			if curr_client.Username == client.Username && curr_client.Password==client.Password{
				allClients[i].Address = client.Address
				allClients[i].PublicKey = client.PublicKey
				context.JSON(200, gin.H{"msg": "Details updated!"})
				return
			} else{
				context.JSON(401, gin.H{"msg": "Authentication failed!"})
				return
			}
		}
	}
	allClients = append(allClients, client)
	context.JSON(201, gin.H{"msg": "registered client"})
}

func UnregisterClient(context *gin.Context){
	id := context.Params.ByName("id")
	index := -1
	for i, client := range(allClients){
		if client.ID==id{
			index = i
			break
		}
	}
	if index!=-1{
		allClients[index] = allClients[len(allClients)-1]
    	allClients = allClients[:len(allClients)-1]
		context.JSON(200, gin.H{"msg": "client unregistered!"})
	}
	context.JSON(404, gin.H{"msg": "Client not found!"})
}