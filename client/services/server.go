package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterSelf(id string, address string){
	payload := map[string] string{
		"id": id,
		"connection_address": address,
	}
	payload_bytes, _ := json.Marshal(payload)

	reader := bytes.NewReader(payload_bytes)

	res , err := http.Post("http://localhost:3000/clients", "application/json", reader)

	if err!=nil{
		panic(err)
	}

	fmt.Println(res)
}