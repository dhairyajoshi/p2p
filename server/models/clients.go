package models

type Client struct{
	ID string `json:"id"`
	ConnectionString string `json:"connection_string"`
}