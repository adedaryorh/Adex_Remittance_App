package main

import (
	"Fin-Remittance/api"
	"fmt"
)

func main() {
	fmt.Println("This is golang remittance app")
	api.NewServer(4000)
}
