package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://gizug0-vd.myshopify.com","https://peterjacksons.com"},
		AllowCredentials: true,
		Debug: true,
	})

	r := mux.NewRouter()
	r.HandleFunc("*", CatchAllHandler)
	r.HandleFunc("/ping", PingHandler).Methods(http.MethodGet)
	r.HandleFunc("/ping", PostPingHandler).Methods(http.MethodPost)

	log.Printf("serving requests on 0.0.0.0:5000")
	http.ListenAndServe("0.0.0.0:5000", c.Handler(r))
}

func CatchAllHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("ping all")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; utf-8")
	w.Write([]byte("pong all"))
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("ping")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; utf-8")
	w.Write([]byte("pong"))
}

func PostPingHandler(w http.ResponseWriter, r *http.Request) {
	decoded, err := json.MarshalIndent(r.Header, "", "  ")

	if err != nil {
		log.Printf("unable to decode headers: %v", err)
		internalServerError(w)
	}

	log.Printf("headers: %s", decoded)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; utf-8")
	w.Write([]byte("pong"))
}	

func internalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	if _, err := w.Write([]byte(`{"message": "internal server error"}`)); err != nil{
		log.Printf("unable to respond to request: %v", err)
		return
	}
}
