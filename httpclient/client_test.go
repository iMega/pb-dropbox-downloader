package httpclient

import (
	"net/http"
	"net/http/httptest"
	"pb-dropbox-downloader/logger"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserAgent(t *testing.T) {
	testServer := httptest.NewServer(
		http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			assert.Equal(t, http.MethodGet, req.Method)
			assert.Equal(t, "Agent", req.Header.Get("User-Agent"))
		}),
	)
	defer func() { testServer.Close() }()

	conf := Config{UserAgent: "Agent"}
	client := New(conf, logger.New("", nil))

	req, err := http.NewRequest(http.MethodGet, testServer.URL, nil)
	require.NoError(t, err)

	actual, err := client.Do(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, actual.StatusCode)
}

func TestRetry(t *testing.T) {
	numberAttempts := 0
	testServer := httptest.NewServer(
		http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			numberAttempts++
			if numberAttempts < 2 {
				resp.WriteHeader(http.StatusInternalServerError)
			}
		}),
	)
	defer func() { testServer.Close() }()

	conf := Config{}
	client := New(conf, logger.New("", nil))

	req, err := http.NewRequest(http.MethodGet, testServer.URL, nil)
	require.NoError(t, err)

	_, err = client.Do(req)
	require.NoError(t, err)

	assert.Equal(t, 2, numberAttempts)
}
