package loader

import (
	"loadBal/internal/server"
)

type Backend struct {
	Id    int
	Url   string
	Alive bool
}

var Backends []Backend

func AddBackEnd() {
	serversData := server.Servers
	for _, data := range serversData {
		Backends = append(Backends, Backend{Id: data.Id, Url: "http://localhost" + data.Port, Alive: data.Alive})
	}

}
