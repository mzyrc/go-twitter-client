package go_twitter_client

import (
	"net/url"
	"strings"
)

func encodeParams(s string) string {
	return strings.Replace(url.QueryEscape(s), "+", "%20", -1)
}
