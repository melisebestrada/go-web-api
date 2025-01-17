package main

import (
	"fmt"
	"log"

	"github.com/melisebestrada/go-web-api/cmd/server"
)

func main() {
	conf := &server.ConfigServer{
		ServerAddress: ":8080",
		DataFilePath:  "data/products.json",
	}

	log.Println("Server running on: 8080")
	app := server.NewServer(conf)
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}

}
