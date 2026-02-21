package main

import (
	"loadBal/internal/loader"
	"loadBal/internal/server"
)

func main() {
	loader.AddBackEnd()
	Server_data := server.Servers
	for _, data := range Server_data {
		go server.Server(data.Id, data.Port)
	}
	for {

	}
}
