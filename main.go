package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("."))))

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
		req.Header.Set("Content-Type", "text/plain")
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})

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
