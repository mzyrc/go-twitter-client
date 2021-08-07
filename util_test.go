package go_twitter_client

import (
	"net/http"
	"testing"
)

func TestGetParameterString(t *testing.T) {
	params := make(map[string]string)

	params["include_entities"] = "true"
	params["oauth_consumer_key"] = "xvz1evFS4wEEPTGEFPHBog"
	params["oauth_nonce"] = "kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg"
	params["oauth_signature_method"] = "HMAC-SHA1"
	params["oauth_timestamp"] = "1318622958"
	params["oauth_token"] = "370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb"
	params["oauth_version"] = "1.0"
	params["status"] = "Hello Ladies + Gentlemen, a signed OAuth request!"

	actual := GetParameterString(params)
	expected := "include_entities=true&oauth_consumer_key=xvz1evFS4wEEPTGEFPHBog&oauth_nonce=kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg&oauth_signature_method=HMAC-SHA1&oauth_timestamp=1318622958&oauth_token=370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb&oauth_version=1.0&status=Hello%20Ladies%20%2B%20Gentlemen%2C%20a%20signed%20OAuth%20request%21"

	if actual != expected {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}

func TestGetSignatureBaseString(t *testing.T) {
	parameterString := "include_entities=true&oauth_consumer_key=xvz1evFS4wEEPTGEFPHBog&oauth_nonce=kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg&oauth_signature_method=HMAC-SHA1&oauth_timestamp=1318622958&oauth_token=370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb&oauth_version=1.0&status=Hello%20Ladies%20%2B%20Gentlemen%2C%20a%20signed%20OAuth%20request%21"

	actual := GetSignatureBaseString(http.MethodPost, "https://api.twitter.com/1.1/statuses/update.json", parameterString)
	expected := "POST&https%3A%2F%2Fapi.twitter.com%2F1.1%2Fstatuses%2Fupdate.json&include_entities%3Dtrue%26oauth_consumer_key%3Dxvz1evFS4wEEPTGEFPHBog%26oauth_nonce%3DkYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg%26oauth_signature_method%3DHMAC-SHA1%26oauth_timestamp%3D1318622958%26oauth_token%3D370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb%26oauth_version%3D1.0%26status%3DHello%2520Ladies%2520%252B%2520Gentlemen%252C%2520a%2520signed%2520OAuth%2520request%2521"

	if actual != expected {
		t.Fatalf("Expected: %v but got: %v", expected, actual)
	}
}

func TestGetSigningKey(t *testing.T) {
	secrets := OAuthSecrets{
		ConsumerSecret: "kAcSOqF21Fu85e7zjz7ZN2U4ZRhfV3WpwPAoE3Z7kBw",
		TokenSecret:    "LswwdoUaIvS8ltyTt5jkRh4J50vUPVVHtR2YPi5kE",
	}

	actual := GetSigningKey(&secrets)
	expected := "kAcSOqF21Fu85e7zjz7ZN2U4ZRhfV3WpwPAoE3Z7kBw&LswwdoUaIvS8ltyTt5jkRh4J50vUPVVHtR2YPi5kE"

	if actual != expected {
		t.Fatalf("Expected: %v but got: %v", expected, actual)
	}
}

func TestGetSignature(t *testing.T) {
	signatureBaseString := "POST&https%3A%2F%2Fapi.twitter.com%2F1.1%2Fstatuses%2Fupdate.json&include_entities%3Dtrue%26oauth_consumer_key%3Dxvz1evFS4wEEPTGEFPHBog%26oauth_nonce%3DkYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg%26oauth_signature_method%3DHMAC-SHA1%26oauth_timestamp%3D1318622958%26oauth_token%3D370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb%26oauth_version%3D1.0%26status%3DHello%2520Ladies%2520%252B%2520Gentlemen%252C%2520a%2520signed%2520OAuth%2520request%2521"
	signingKey := "kAcSOqF21Fu85e7zjz7ZN2U4ZRhfV3WpwPAoE3Z7kBw&LswwdoUaIvS8ltyTt5jkRh4J50vUPVVHtR2YPi5kE"

	actual := GetSignature(signatureBaseString, signingKey)
	expected := "hCtSmYh+iHYCEqBWrE7C7hYmtUk="

	if actual != expected {
		t.Fatalf("Expected: %v but got: %v", expected, actual)
	}
}

func Test_Integration_Twitter_Docs_Example(t *testing.T) {
	params := make(map[string]string)
	params["include_entities"] = "true"
	params["oauth_consumer_key"] = "xvz1evFS4wEEPTGEFPHBog"
	params["oauth_nonce"] = "kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg"
	params["oauth_signature_method"] = "HMAC-SHA1"
	params["oauth_timestamp"] = "1318622958"
	params["oauth_token"] = "370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb"
	params["oauth_version"] = "1.0"
	params["status"] = "Hello Ladies + Gentlemen, a signed OAuth request!"
	
	secrets := OAuthSecrets{
		ConsumerSecret: "kAcSOqF21Fu85e7zjz7ZN2U4ZRhfV3WpwPAoE3Z7kBw",
		TokenSecret:    "LswwdoUaIvS8ltyTt5jkRh4J50vUPVVHtR2YPi5kE",
	}
	
	
	actual := CreateSignature(http.MethodPost, "https://api.twitter.com/1.1/statuses/update.json", params, &secrets)
	expected := "hCtSmYh+iHYCEqBWrE7C7hYmtUk="

	if actual != expected {
		t.Fatalf("Expected: %v but got: %v", expected, actual)
	}
}

func Test_Integration_Postman_Scenario(t *testing.T) {
	params := make(map[string]string)
	params["oauth_callback"] = "http://localhost:3000/callback"
	params["oauth_consumer_key"] = "Xktj2isfgxFKErcj9E7t2rK7P"
	params["oauth_nonce"] = "kD2XXnZyM9J"
	params["oauth_signature_method"] = "HMAC-SHA1"
	params["oauth_timestamp"] = "1628262814"
	params["oauth_token"] = "1057694577585516544-oJvMcAfXaDoTATjJPjfcgrldvkUcXN"
	params["oauth_version"] = "1.0"

	secrets := OAuthSecrets{
		ConsumerSecret: "u08t3dO0LIma1pllUNeYtH4HJW4trscjla4RXKxF3SwhokgGjp",
		TokenSecret:    "O12WM5xJtaz8i7ipI47mcHftbTZEoSQU3JxTjsbKKFuSP",
	}
	
	actual := encodeParams(CreateSignature(http.MethodPost, "https://api.twitter.com/oauth/request_token", params, &secrets))
	expected := "NZY3uzdzBIAfwOxzctOvAvQyYc8%3D"

	if actual != expected {
		t.Fatalf("Expected: %v but got: %v", expected, actual)
	}
}
