package handlers

import (
	"encoding/json"
	"go-twitter-client/auth"
	"log"
	"net/http"
	"os"
)

type tokenPostBody struct {
	OAuthToken    string `json:"oauth_token"`
	OAuthVerifier string `json:"oauth_verifier"`
}

func GetUserAccessTokens(writer http.ResponseWriter, request *http.Request) {
	var postBody tokenPostBody
	requestBodyDecoder := json.NewDecoder(request.Body)
	err := requestBodyDecoder.Decode(&postBody)

	if err != nil {
		respondWithError(writer, http.StatusBadRequest, "Could not read tokens in post body")
		return
	}

	credentials := auth.Credentials{
		ConsumerKey:    os.Getenv("OAUTH_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("OAUTH_CONSUMER_SECRET"),
		Token:          os.Getenv("OAUTH_TOKEN"),
		TokenSecret:    os.Getenv("OAUTH_TOKEN_SECRET"),
	}

	authenticator := auth.NewTwitterOAuth1(&credentials)

	token, tokenSecret, _ := authenticator.GetUserTokens(postBody.OAuthToken, postBody.OAuthVerifier)

	log.Println(token)
	log.Println(tokenSecret)
	respondWithSuccess(writer, http.StatusOK, nil)
}
