package services

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Connect(port string){
	fmt.Printf("connecting to %s\n", port)
	new_conn, err := net.Dial("tcp", fmt.Sprintf(":%s", port))
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer new_conn.Close()

	for {

		fmt.Print("Enter the message: ")
		in := bufio.NewReader(os.Stdin)
		send, _ := in.ReadString('\n')
		send = strings.TrimRight(send,"\r\n")
		write(new_conn, send)

		if send=="exit"{
			break
		}

		message, _ := read(new_conn)

		fmt.Println(message)

	}

}