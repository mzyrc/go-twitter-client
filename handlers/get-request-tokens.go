package handlers

import (
	"go-twitter-client/auth"
	"log"
	"net/http"
	"os"
)

func GetRequestTokensHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("Fetching tokens")

	credentials := auth.Credentials{
		ConsumerKey:    os.Getenv("OAUTH_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("OAUTH_CONSUMER_SECRET"),
		Token:          os.Getenv("OAUTH_TOKEN"),
		TokenSecret:    os.Getenv("OAUTH_TOKEN_SECRET"),
	}

	authenticator := auth.NewTwitterOAuth1(&credentials)

	tokens, err := authenticator.GetRequestToken()

	if err != nil {
		respondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}

	respondWithSuccess(writer, http.StatusOK, tokens)
}
