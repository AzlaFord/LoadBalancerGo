package server

import (
	"io"
	"net/http"
)

type BackEndServer struct {
	Id    int
	Port  string
	Alive bool
}

var Servers = []BackEndServer{
	{Id: 0, Port: ":8080", Alive: true},
	{Id: 1, Port: ":6060", Alive: true},
	{Id: 2, Port: ":4040", Alive: true},
}

func Server(id int, port string) {

	server := http.NewServeMux()
	server.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello from healthHandler #"+port+"\n")
	})
	server.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		// time.Sleep(3 * time.Second) test
		io.WriteString(w, "Hello from apiHandler #"+port+"\n")
	})

	http.ListenAndServe(port, server)
}
