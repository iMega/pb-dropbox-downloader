package auth_test

import (
	"fmt"
	"pb-dropbox-downloader/auth"
	"pb-dropbox-downloader/config"
	"testing"
)

func TestCreateAuthorizationURL(t *testing.T) {
	// expected := auth.PKCE{
	// 	AuthorizationURL: "url-ID-huVe5Sif9SgWtYkAgGw7CvEPQ6NI0AdBuSVp1DNWPLI",
	// 	CodeVerifier:     "MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTIzNDU2Nzg5MDE",
	// }

	conf := &config.Config{
		AppID:          "fuwgr8q8src9lk8",
		DropboxAuthURL: "https://www.dropbox.com/oauth2/authorize?response_type=code&code_challenge_method=S256&token_access_type=offline&redirect_uri=https://pb.imega.ru&client_id=%s&code_challenge=%s",
		IsTest:         false,
	}

	actual := auth.CreateAuthorizationURL(conf)

	fmt.Println(actual.AuthorizationURL)

	fmt.Println(`curl -s https://api.dropbox.com/oauth2/token \`)
	fmt.Println(`-d client_id=fuwgr8q8src9lk8 \`)
	fmt.Println(`-d redirect_uri=https://pb.imega.ru \`)
	fmt.Println(`-d grant_type=authorization_code \`)
	fmt.Println("-d code_verifier=" + actual.CodeVerifier + ` \`)
	fmt.Println(`-d code=`)
	fmt.Println(actual.CodeVerifier)

	// assert.Equal(t, expected, actual)
}
