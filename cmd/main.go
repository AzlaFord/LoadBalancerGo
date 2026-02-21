package main

import (
	"loadBal/internal/server"
)

func main() {
	Server_data := server.Servers
	for _, data := range Server_data {
		go server.Server(data.Id, data.Port)
	}
	for {

	}
}
