package main

import (
	"loadBal/internal/loader"
	"loadBal/internal/server"
	"net/http"
)

func main() {
	loader.AddBackEnd()
	go loader.HealthCheck()
	Server_data := server.Servers
	http.HandleFunc("/api", loader.ProxyHandler)
	for _, data := range Server_data {
		go server.Server(data.Id, data.Port)
	}

	http.ListenAndServe(":9090", nil)
}
