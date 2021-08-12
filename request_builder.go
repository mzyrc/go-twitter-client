package go_twitter_client

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

type AuthConfig struct {
	ConsumerKey       string
	ConsumerKeySecret string
	Token             string
	TokenSecret       string
}

type RequestConfig struct {
	Method            string
	CustomOAuthParams map[string]string
}

type RequestBuilder struct {
	authConfig *AuthConfig
}

func NewRequestBuilder(config *AuthConfig) *RequestBuilder {
	return &RequestBuilder{authConfig: config}
}

func (rb *RequestBuilder) CreateRequest(url string, config RequestConfig) (*http.Request, error) {
	if config.Method == "" {
		return nil, errors.New("missing http method in config")
	}

	request, err := http.NewRequest(config.Method, url, nil)

	oauth := NewOAuth(rb.authConfig)
	oauthHeader := oauth.GetAuthorizationHeader(config.Method, url, config.CustomOAuthParams)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", oauthHeader)

	if err != nil {
		return nil, err
	}

	return request, nil
}

func (rb *RequestBuilder) getRequiredOAuthParams() map[string]string {
	OAuthParams := make(map[string]string)
	OAuthParams["oauth_consumer_key"] = rb.authConfig.ConsumerKey
	OAuthParams["oauth_nonce"] = uuid.NewV4().String()
	OAuthParams["oauth_signature_method"] = "HMAC-SHA1"
	OAuthParams["oauth_timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	OAuthParams["oauth_token"] = rb.authConfig.Token
	OAuthParams["oauth_version"] = "1.0"

	return OAuthParams
}
