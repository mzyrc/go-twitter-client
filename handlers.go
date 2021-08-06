package main

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func GetRequestTokensHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("Fetching tokens")

	tokens, err := getRequestTokens()

	if err != nil {
		respondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}

	respondWithSuccess(writer, http.StatusOK, tokens)
}

func getRequestTokens() (string, error) {
	request, _ := http.NewRequest(http.MethodPost, "https://api.twitter.com/oauth/request_token", nil)

	OAuthParams := make(map[string]string)
	OAuthParams["oauth_callback"] = os.Getenv("OAUTH_CALLBACK")
	OAuthParams["oauth_consumer_key"] = os.Getenv("OAUTH_CONSUMER_KEY")
	OAuthParams["oauth_nonce"] = uuid.NewV4().String()
	OAuthParams["oauth_signature_method"] = "HMAC-SHA1"
	OAuthParams["oauth_timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	OAuthParams["oauth_token"] = os.Getenv("OAUTH_TOKEN")
	OAuthParams["oauth_version"] = "1.0"

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", getOAuthHeaders(OAuthParams))

	client := http.Client{}
	response, err := client.Do(request)

	if response.StatusCode != http.StatusOK {
		return "", errors.New("Could not request token")
	}

	log.Println(response.Body)

	if err != nil {
		log.Println(err)
	}

	body, _ := ioutil.ReadAll(response.Body)

	values, err := url.ParseQuery(string(body))

	if err != nil {
		return "", errors.New("Error occured when parsing tokens")
	}

	values.Get("oauth_token")

	return values.Get("oauth_token"), nil
}

func getOAuthHeaders(params map[string]string) string {
	secrets := OAuthSecrets{
		ConsumerSecret: os.Getenv("OAUTH_CONSUMER_SECRET"),
		TokenSecret:    os.Getenv("OAUTH_TOKEN_SECRET"),
	}

	signature := CreateSignature(http.MethodPost, "https://api.twitter.com/oauth/request_token", params, &secrets)

	nonceHeader := fmt.Sprintf("oauth_nonce=\"%s\"", params["oauth_nonce"])
	callbackHeader := fmt.Sprintf("oauth_callback=\"%s\"", encodeParams(params["oauth_callback"]))
	signatureMethodHeader := fmt.Sprintf("oauth_signature_method=\"%s\"", params["oauth_signature_method"])
	signatureHeader := fmt.Sprintf("oauth_signature=\"%s\"", encodeParams(signature))
	oauthTimestampHeader := fmt.Sprintf("oauth_timestamp=\"%s\"", params["oauth_timestamp"])
	consumerKeyHeader := fmt.Sprintf("oauth_consumer_key=\"%s\"", params["oauth_consumer_key"])
	versionHeader := fmt.Sprintf("oauth_version=\"%s\"", params["oauth_version"])
	oauthTokenHeader := fmt.Sprintf("oauth_token=\"%s\"", params["oauth_token"])

	return fmt.Sprintf("OAuth %s,%s,%s,%s,%s,%s,%s,%s", consumerKeyHeader, oauthTokenHeader, signatureMethodHeader, oauthTimestampHeader, nonceHeader, versionHeader, callbackHeader, signatureHeader)
}
