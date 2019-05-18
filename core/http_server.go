package core

import (
	"log"
	"net/http"
)

func InitHTTPServer(address string, handlers map[string]func(http.ResponseWriter, *http.Request)) {
	mux := http.NewServeMux()
	for path, handler := range handlers {
		mux.HandleFunc(path, handler)
	}
	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	WaitGroup.Add(1)
	go func() {
		defer WaitGroup.Done()
		err := server.ListenAndServe()
		if err != nil {
			log.Panicln("cannot listen http with addr[", address, "], err=", err)
		}
	}()
}