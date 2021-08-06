package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/token", GetRequestTokensHandler).Methods(http.MethodGet)
	handler := cors.Default().Handler(router)

	http.ListenAndServe(":8000", handler)
}
