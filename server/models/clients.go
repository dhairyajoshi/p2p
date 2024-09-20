package models

type Client struct{
	ID string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	PublicKey string `json:"public_key"`
	Address string `json:"address"`
	Online bool `json:"online"`
	LastSeen string `json:"last_seen"`
}