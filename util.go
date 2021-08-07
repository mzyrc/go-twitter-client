package go_twitter_client

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

type OAuthSecrets struct {
	ConsumerSecret string
	TokenSecret    string
}

func CreateSignature(httpMethod string, baseURL string, params map[string]string, oauthCredentials *OAuthSecrets) string {
	parameterString := GetParameterString(params)
	signatureBaseString := GetSignatureBaseString(httpMethod, baseURL, parameterString)
	signingKey := GetSigningKey(oauthCredentials)

	return GetSignature(signatureBaseString, signingKey)
}

func GetParameterString(params map[string]string) string {
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

func encodeParams(s string) string {
	return strings.Replace(url.QueryEscape(s), "+", "%20", -1)
}

func GetSignatureBaseString(httpMethod string, baseURL string, parameterString string) string {
	return fmt.Sprintf("%s&%s&%s", httpMethod, encodeParams(baseURL), encodeParams(parameterString))
}

func GetSigningKey(OAuthSecrets *OAuthSecrets) string {
	return fmt.Sprintf("%s&%s", OAuthSecrets.ConsumerSecret, OAuthSecrets.TokenSecret)
}

func GetSignature(signatureBaseString string, signingKey string) string {
	mac := hmac.New(sha1.New, []byte(signingKey))
	mac.Write([]byte(signatureBaseString))

	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
