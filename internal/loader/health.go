package loader

import (
	"net/http"
	"time"
)

func HealthCheck() {
	for {
		for i, b := range Backends {
			resp, err := http.Get(b.Url)
			if err != nil || resp.StatusCode != 200 {
				Backends[i].Alive = false
			} else {
				Backends[i].Alive = true
			}
			if resp != nil {
				resp.Body.Close()
			}
		}
		time.Sleep(5 * time.Second)
	}
}
