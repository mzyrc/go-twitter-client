package go_twitter_client

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"sort"
	"time"
)

type oauth struct {
	secrets *AuthConfig
}

func NewOAuth(secrets *AuthConfig) *oauth {
	return &oauth{secrets: secrets}
}

func (o *oauth) GetAuthorizationHeader(httpMethod string, endpoint string, additionalParams map[string]string) string {
	oauthHeaders := o.getRequiredOAuthParams()

	for key, value := range additionalParams {
		oauthHeaders[key] = value
	}

	oauthHeaders["oauth_signature"] = o.createSignature(httpMethod, endpoint, oauthHeaders)
	return o.createAuthorizationHeader(oauthHeaders)
}

func (o *oauth) getRequiredOAuthParams() map[string]string {
	OAuthParams := make(map[string]string)
	OAuthParams["oauth_consumer_key"] = o.secrets.ConsumerKey
	OAuthParams["oauth_nonce"] = uuid.NewV4().String()
	OAuthParams["oauth_signature_method"] = "HMAC-SHA1"
	OAuthParams["oauth_timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	OAuthParams["oauth_token"] = o.secrets.Token
	OAuthParams["oauth_version"] = "1.0"

	return OAuthParams
}

func (o *oauth) createSignature(httpMethod string, endpoint string, params map[string]string) string {
	parameterString := o.getParameterString(params)
	signatureBaseString := o.getSignatureBaseString(httpMethod, endpoint, parameterString)
	signingKey := o.getSigningKey()

	return o.getSignature(signatureBaseString, signingKey)
}

func (o *oauth) getParameterString(params map[string]string) string {
	var finalParameterString string
	numberOfParams := len(params)

	keys := make([]string, 0, len(params))
	for key, _ := range params {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for index, key := range keys {
		finalParameterString = finalParameterString + encodeParams(key) + "=" + encodeParams(params[key])

		if index != numberOfParams-1 {
			finalParameterString = finalParameterString + "&"
		}
	}

	return finalParameterString
}

func (o *oauth) getSignatureBaseString(httpMethod string, baseURL string, parameterString string) string {
	return fmt.Sprintf("%s&%s&%s", httpMethod, encodeParams(baseURL), encodeParams(parameterString))
}

func (o *oauth) getSigningKey() string {
	return fmt.Sprintf("%s&%s", o.secrets.ConsumerKeySecret, o.secrets.TokenSecret)
}

func (o *oauth) getSignature(signatureBaseString string, signingKey string) string {
	mac := hmac.New(sha1.New, []byte(signingKey))
	mac.Write([]byte(signatureBaseString))

	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func (o *oauth) createAuthorizationHeader(OAuthParams map[string]string) string {
	header := "OAuth "

	numberOfParams := len(OAuthParams)
	counter := 1

	for key, value := range OAuthParams {
		if key == "oauth_signature" || key == "oauth_callback" {
			value = encodeParams(value)
		}

		header = fmt.Sprintf("%s%s=\"%s\"", header, key, value)

		if counter <= numberOfParams-1 {
			header = fmt.Sprintf("%s,", header)
		}

		counter = counter + 1
	}

	return header
}
