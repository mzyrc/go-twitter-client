package go_twitter_client

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRequestBuilder_CreateRequest(t *testing.T) {
	builder := NewRequestBuilder(&AuthConfig{
		ConsumerKey:       "sfzHizjt6RgMXuxhfPtRWHsn4",
		ConsumerKeySecret: "I6PQWTjYRe32D5MQCH6lXST7zEGDqdQh8UQV7LpYGfk5FxbVeU",
		Token:             "1057694577585516544-JFzGFsEaEJqmgQE4O2lT9UaqMJH49P",
		TokenSecret:       "ZjOqaVXgPuc39o0q7J0DkjUxfJlPop2aHSFXPO9TRMwwK",
	})

	config := RequestConfig{Method: http.MethodGet}

	url := "https://api.twitter.com/2/users/1057694577585516544"
	request, err := builder.CreateRequest(url, config)

	if err != nil {
		t.Fatalf("Test failed unexpectedly with: %v", err.Error())
	}

	assert.NotEmpty(t, request.Header.Get("Authorization"), "")
	assert.NotEmpty(t, request.Header.Get("Content-Type"), "")
	assert.Equal(t, request.Method, config.Method)
}
