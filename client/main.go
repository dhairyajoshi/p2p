package main

import (
	"os"
	"p2p-client/services"
)


func main(){
	if len(os.Args) !=3 {
		panic("2 arguments expected! [listen, connect] [port]")
	}

	port := os.Args[2]

	if os.Args[1]=="listen"{
		services.Listen(port)
	} else if os.Args[1]=="connect"{
		services.Connect(port)
	}
}