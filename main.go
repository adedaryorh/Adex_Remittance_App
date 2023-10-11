package main

import (
	"Fin-Remittance/api"
)

func main() {
	server := api.NewServer(".")
	server.Start(3000)
}
