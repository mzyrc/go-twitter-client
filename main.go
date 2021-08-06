package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"go-twitter-client/handlers"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/oauth/request_token", handlers.GetRequestTokensHandler).Methods(http.MethodGet)
	router.HandleFunc("/oauth/access_token", handlers.GetUserAccessTokens).Methods(http.MethodPost)

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8000", handler))
}
