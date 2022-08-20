package auth

import (
	"pb-dropbox-downloader/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAuthorizationURL(t *testing.T) {
	expected := PKCE{
		AuthorizationURL: "url-ID-huVe5Sif9SgWtYkAgGw7CvEPQ6NI0AdBuSVp1DNWPLI",
		CodeVerifier:     "MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTIzNDU2Nzg5MDE",
	}

	conf := config.Config{
		AppID:          "ID",
		DropboxAuthURL: "url-%s-%s",
		IsTest:         true,
	}

	actual := CreateAuthorizationURL(conf)

	assert.Equal(t, expected, actual)
}
