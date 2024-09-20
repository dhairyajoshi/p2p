package services

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net"
)

func encrypt(secretMessage []byte, key rsa.PublicKey) []byte {
    label := []byte("OAEP Encrypted")
    rng := rand.Reader
    ciphertext, _ := rsa.EncryptOAEP(sha256.New(), rng, &key, secretMessage, label)
    return []byte(base64.StdEncoding.EncodeToString(ciphertext))
}

func decrypt(cipherText string, privKey rsa.PrivateKey) string {
    ct, _ := base64.StdEncoding.DecodeString(cipherText)
    label := []byte("OAEP Encrypted")
    rng := rand.Reader
    plaintext, _ := rsa.DecryptOAEP(sha256.New(), rng, &privKey, ct, label)
    
    return string(plaintext)
}

func read(conn net.Conn, privateKey rsa.PrivateKey) (string, error){
	message_bytes := make([]byte, 1024)
	
	size, err := conn.Read(message_bytes)
	if size == 0{
		return "", nil
	}

	if err==io.EOF{
		return decrypt(string(message_bytes), privateKey), nil
	}

	if err!=nil{
		panic(err)
	}
	return decrypt(string(message_bytes), privateKey), nil
}

func write(conn net.Conn, message string, publicKey rsa.PublicKey) (err error){
	
	message_bytes := []byte(message)

	_, err = conn.Write(encrypt(message_bytes, publicKey))

	if err!=nil{
		panic(err)
	}

	return
}