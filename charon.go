package main

import (
	"charon/methods"
	"net/http"
	"sync"
)

func initHTTPServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/webdav/", methods.RedirectWebdav)

	server := &http.Server{
		Addr:    "0.0.0.0:8081",
		Handler: mux,
	}

	server.ListenAndServe()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go initHTTPServer()

	wg.Wait()
}
