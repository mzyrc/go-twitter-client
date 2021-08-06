package auth

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

type Credentials struct {
	ConsumerKey    string
	ConsumerSecret string
	Token          string
	TokenSecret    string
}

type TwitterOAuth1 struct {
	credentials *Credentials
}

func NewTwitterOAuth1(credentials *Credentials) *TwitterOAuth1 {
	return &TwitterOAuth1{credentials: credentials}
}

func (t *TwitterOAuth1) GetRequestToken() (string, error) {
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
	request.Header.Set("Authorization", t.getOAuthHeaders(OAuthParams))

	client := http.Client{}
	response, err := client.Do(request)

	if response.StatusCode != http.StatusOK {
		return "", errors.New("Could not request token")
	}

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

func (t *TwitterOAuth1) getOAuthHeaders(params map[string]string) string {
	secrets := OAuthSecrets{
		ConsumerSecret: os.Getenv("OAUTH_CONSUMER_SECRET"),
		TokenSecret:    os.Getenv("OAUTH_TOKEN_SECRET"),
	}

	signature := CreateSignature(http.MethodPost, requestTokenEndpoint, params, &secrets)

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

func (t *TwitterOAuth1) GetUserTokens(oauthToken string, oauthVerifier string) (string, string, error) {
	request, _ := http.NewRequest(http.MethodPost, accessTokenEndpoint, nil)

	query := request.URL.Query()
	query.Add("oauth_token", oauthToken)
	query.Add("oauth_verifier", oauthVerifier)

	request.URL.RawQuery = query.Encode()

	log.Println(request.URL.String())

	client := http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return "", "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", "", errors.New("Could not get user access tokens")
	}

	body, _ := ioutil.ReadAll(response.Body)

	values, err := url.ParseQuery(string(body))

	log.Println(values)

	return values.Get("oauth_token"), values.Get("oauth_token_secret"), nil
}
