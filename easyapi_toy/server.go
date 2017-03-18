package main

import (
	"fmt"
	"os"

	"easyapi"
	"easyapi/easyapi_toy/services/service1"
)

func main() {
	app := easyapi.New()
	app.RegisterService(&service1.Service1{})

	l, err := app.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}

	defer l.Close()
	app.Serve(l)
}
