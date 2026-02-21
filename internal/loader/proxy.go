package loader

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	backend := SelectBackEnd()
	req, err := http.NewRequest(r.Method, backend.Url+r.RequestURI, r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header = r.Header
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, _ := io.ReadAll(resp.Body)
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}
