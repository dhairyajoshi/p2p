package services

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"net"
	"strings"
)

var (
	privateKey, _ = rsa.GenerateKey(rand.Reader, 2048)

	publicKey = privateKey.PublicKey
)

func Listen(port string){
	fmt.Printf("listening on %s\n", port)

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port));
	RegisterSelf("user1", "localhost:3001", publicKey)
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

		message, err := read(conn, *privateKey)

		if len(message) == 0{
			continue
		}

		if err!=nil{
			
		}

		if strings.TrimRight(message,"\r\n\x00") == "exit"{
			fmt.Printf("Closing Connection to %s\n", client)
			break
		}

		// write(conn, "Acknowledged!")

		fmt.Printf("received message from %s: ", client)
		fmt.Println(message)
	}
	
	
	conn.Close()
}