package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))

	server := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	fmt.Printf("server is listening on %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
	}
}
