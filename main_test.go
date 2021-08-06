package main
//
//import (
//	"log"
//	"net/http"
//	"testing"
//)
//
///*
//
//OAuth oauth_consumer_key="Xktj2isfgxFKErcj9E7t2rK7P",oauth_token="1057694577585516544-oJvMcAfXaDoTATjJPjfcgrldvkUcXN",oauth_signature_method="HMAC-SHA1",oauth_timestamp="1628254575",oauth_nonce="FtSmfwyZM3s",oauth_version="1.0",oauth_callback="http%3A%2F%2Flocalhost%3A3000%2Fcallback",oauth_signature="IddH803Kar%2FHBOgKehRPlCdWh%2FI%3D"
//
//*/
//
//func Test_ApplyOAuthSignature(t *testing.T) {
//	apiKey := "Xktj2isfgxFKErcj9E7t2rK7P"
//	apiSecretKey := "u08t3dO0LIma1pllUNeYtH4HJW4trscjla4RXKxF3SwhokgGjp"
//	accessToken := "1057694577585516544-oJvMcAfXaDoTATjJPjfcgrldvkUcXN"
//	accessTokenSecret := "O12WM5xJtaz8i7ipI47mcHftbTZEoSQU3JxTjsbKKFuSP"
//
//	credentials := OAuth{
//		Nonce:             "FtSmfwyZM3s",
//		Callback:          "http://localhost:3000/callback",
//		SignatureMethod:   "HMAC-SHA1",
//		Timestamp:         1628254575,
//		ConsumerKey:       apiKey,
//		ConsumerSecretKey: apiSecretKey,
//		Version:           "1.0",
//		Token:             accessToken,
//		TokenSecret:       accessTokenSecret,
//	}
//	//
//	//result := url.QueryEscape("include_entities=true&oauth_consumer_key=xvz1evFS4wEEPTGEFPHBog&oauth_nonce=kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg&oauth_signature_method=HMAC-SHA1&oauth_timestamp=1318622958&oauth_token=370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb&oauth_version=1.0&status=Hello Ladies + Gentlemen, a signed OAuth request!")
//	//expectedEncodedStr := "include_entities=true&oauth_consumer_key=xvz1evFS4wEEPTGEFPHBog&oauth_nonce=kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg&oauth_signature_method=HMAC-SHA1&oauth_timestamp=1318622958&oauth_token=370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb&oauth_version=1.0&status=Hello%20Ladies%20%2B%20Gentlemen%2C%20a%20signed%20OAuth%20request%21"
//
//
//	/*
//
//	include_entities=true&oauth_consumer_key=xvz1evFS4wEEPTGEFPHBog&oauth_nonce=kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg&oauth_signature_method=HMAC-SHA1&oauth_timestamp=1318622958&oauth_token=370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb&oauth_version=1.0&status=Hello%20Ladies%20%2B%20Gentlemen%2C%20a%20signed%20OAuth%20request%21 but got
//	include_entities%3Dtrue%26oauth_consumer_key%3Dxvz1evFS4wEEPTGEFPHBog%26oauth_nonce%3DkYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg%26oauth_signature_method%3DHMAC-SHA1%26oauth_timestamp%3D1318622958%26oauth_token%3D370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb%26oauth_version%3D1.0%26status%3DHello+Ladies+%2B+Gentlemen%2C+a+signed+OAuth+request%21
//	*/
//
//	//if result != expectedEncodedStr {
//	//	t.Fatalf("Expected %v but got %v", expectedEncodedStr, result)
//	//}
//
//	signature,_ := EncodeSignature(http.MethodPost, "https://api.twitter.com/oauth/request_token", &credentials, nil)
//
//	log.Println(signature)
//
//	result2 := ApplyOAuthSignature(credentials, "https://api.twitter.com/oauth/request_token")
//
//	log.Println(result2.Get("oauth_signature"))
//
//	expected := "IddH803Kar%2FHBOgKehRPlCdWh%2FI%3D"
//	actual := result2.Get("oauth_signature")
//
//	if actual != expected {
//		t.Fatalf("Expected %s but got %s", expected, actual)
//	}
//}
//
//func Test_Twitter_Apply(t *testing.T) {
//
//}