package services

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func RegisterSelf(id string, address string, public_key rsa.PublicKey){
	payload := map[string] string{
		"id": id,
		"address": address,
		"public_key": fmt.Sprintf("%d %d", public_key.N, public_key.E),
	}
	payload_bytes, _ := json.Marshal(payload)

	reader := bytes.NewReader(payload_bytes)

	res , err := http.Post("http://localhost:3000/clients", "application/json", reader)

	if err!=nil{
		panic(err)
	}

	fmt.Println(res)
}

func GetClientDetails(id string) (string, string){
	res, err := http.Get(fmt.Sprintf("http://localhost:3000/clients/%s", id))

	if err!=nil{
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err!=nil{
		panic(err)
	}

	json_body := make(map[string] string)

	json.Unmarshal(body, &json_body)

	return json_body["address"], json_body["public_key"]
}