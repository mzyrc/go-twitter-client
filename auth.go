package go_twitter_client

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	requestTokenEndpoint = "https://api.twitter.com/oauth/request_token"
	accessTokenEndpoint  = "https://api.twitter.com/oauth/access_token"
)

type AuthConfig struct {
	ConsumerKey       string `json:"consumer_key"`
	ConsumerKeySecret string `json:"consumer_key_secret"`
	Token             string `json:"token"`
	TokenSecret       string `json:"token_secret"`
}

type UserIdentity struct {
	Token       string `json:"token"`
	TokenSecret string `json:"token_secret"`
	UserId      string `json:"user_id"`
}

type OAuth1Strategy struct {
	config *AuthConfig
}

func NewOAuth1Strategy(config *AuthConfig) *OAuth1Strategy {
	return &OAuth1Strategy{config: config}
}

func (oas *OAuth1Strategy) GetRequestToken() (string, error) {
	builder := NewRequestBuilder(oas.config)

	additionalAuthParams := make(map[string]string)
	additionalAuthParams["oauth_callback"] = "http://localhost:3000/callback"

	request, err := builder.CreateRequest(requestTokenEndpoint, RequestConfig{
		Method:            http.MethodPost,
		CustomOAuthParams: additionalAuthParams,
	})

	if err != nil {
		return "", err
	}

	httpClient := http.Client{}
	response, err := httpClient.Do(request)

	if response.StatusCode != http.StatusOK {
		return "", errors.New("could not request token")
	}

	if err != nil {
		log.Println(err)
	}

	body, _ := ioutil.ReadAll(response.Body)

	values, err := url.ParseQuery(string(body))

	if err != nil {
		return "", errors.New("error occurred when parsing tokens")
	}

	return values.Get("oauth_token"), nil
}

func (oas *OAuth1Strategy) GetUserAccessTokens(oauthToken string, oauthVerifier string) (*UserIdentity, error) {
	request, _ := http.NewRequest(http.MethodPost, accessTokenEndpoint, nil)

	query := request.URL.Query()
	query.Add("oauth_token", oauthToken)
	query.Add("oauth_verifier", oauthVerifier)

	request.URL.RawQuery = query.Encode()
	client := http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("could not get user access tokens")
	}

	body, _ := ioutil.ReadAll(response.Body)
	values, err := url.ParseQuery(string(body))

	return &UserIdentity{
		Token:       values.Get("oauth_token"),
		TokenSecret: values.Get("oauth_token_secret"),
		UserId:      values.Get("user_id"),
	}, nil
}
