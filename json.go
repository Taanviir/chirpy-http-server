package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}

	type errorResponse struct {
		Err string `json:"error"`
	}

	respondWithJSON(w, code, errorResponse{
		Err: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}

func decodeJSONBody(w http.ResponseWriter, req *http.Request, payload interface{}) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(payload)
	if err != nil {
		respondWithError(w, 500, "Failed to decode parameters", err)
	}
}
