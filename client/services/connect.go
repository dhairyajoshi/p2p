package services

import (
	"bufio"
	"crypto/rsa"
	"fmt"
	"math/big"
	"net"
	"os"
	"strconv"
	"strings"
)

func Connect(port string){
	fmt.Printf("connecting to %s\n", port)

	address, public_key := GetClientDetails(port)

	key_values := strings.Split(public_key, " ")
	
	new_conn, err := net.Dial("tcp", address)
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
		var int_val, _ = new(big.Int).SetString(key_values[0], 10)
		E, _ := strconv.ParseInt(key_values[1], 10, 64)
		write(new_conn, send, rsa.PublicKey{N: int_val, E: int(E)})

		if send=="exit"{
			break
		}

	}

}