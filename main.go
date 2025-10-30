package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", PingHandler)

	log.Printf("serving requests on 0.0.0.0:5000")
	http.ListenAndServe("0.0.0.0:5000", nil)
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("ping")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; utf-8")
	w.Write([]byte("pong"))
}
