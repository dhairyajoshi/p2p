package services

import (
	"fmt"
	"net"
	"strings"
)

func Listen(port string){
	fmt.Printf("listening on %s\n", port)

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port));
	
	if err!=nil{
		msg:= fmt.Sprintf("Error occured: %s", err)
		fmt.Println(msg)
	}

	for{
		conn, err:= l.Accept();
		if err!=nil{
			panic(err)
		}
		
		go startConnection(conn)
	}
	
}

func startConnection(conn net.Conn){
	client := conn.RemoteAddr()

	for{

		message, err := read(conn)

		if len(message) == 0{
			continue
		}

		if err!=nil{
			
		}

		if strings.TrimRight(message,"\r\n\x00") == "exit"{
			fmt.Printf("Closing Connection to %s\n", client)
			break
		}

		write(conn, "Acknowledged!")

		fmt.Printf("received message from %s: ", client)
		fmt.Println(message)
	}
	
	
	conn.Close()
}