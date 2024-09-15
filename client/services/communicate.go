package services

import (
	"io"
	"net"
)

func read(conn net.Conn) (string, error){
	message_bytes := make([]byte, 1024)
	
	size, err := conn.Read(message_bytes)

	if size == 0{
		return "", nil
	}

	if err==io.EOF{
		return string(message_bytes), nil
	}

	if err!=nil{
		panic(err)
	}

	return string(message_bytes), nil

}

func write(conn net.Conn, message string) (err error){
	message_bytes := []byte(message)

	_, err = conn.Write(message_bytes)

	if err!=nil{
		panic(err)
	}

	return
}