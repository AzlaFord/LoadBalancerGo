package loader

import (
	"loadBal/internal/server"
)

type Backend struct {
	Id    int
	Url   string
	Alive bool
}

var currentIndex int = 0
var Backends []Backend

func AddBackEnd() {
	serversData := server.Servers
	for _, data := range serversData {
		Backends = append(Backends, Backend{Id: data.Id, Url: "http://localhost" + data.Port, Alive: data.Alive})
	}

}

func SelectBackEnd() Backend {
	var AliveBackends []Backend
	for _, data := range Backends {
		if data.Alive == true {
			AliveBackends = append(AliveBackends, data)
		}
	}
	backend := AliveBackends[currentIndex%len(AliveBackends)]
	currentIndex++
	return backend
}
