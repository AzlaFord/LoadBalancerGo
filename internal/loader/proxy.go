package loader

import (
	"fmt"
	"net/http"
)

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	backend := SelectBackEnd()
	req, err := http.NewRequest(r.Method, backend.Url+r.RequestURI, r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header = r.Header

}
