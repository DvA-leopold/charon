package main

import (
	"fmt"
	"net/http"
	"sync"
)

func test(http.ResponseWriter, *http.Request) {
	fmt.Println("hello server")
}

func auth(out http.ResponseWriter, req *http.Request) {
	fmt.Println("hello auth")
}

func initTCPServer() {

}

func initHTTPServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", test)
	mux.HandleFunc("/login", auth)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go initHTTPServer()
	go initTCPServer()

	wg.Wait()
}
