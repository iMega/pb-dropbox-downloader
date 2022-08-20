// Copyright © 2022 Dmitry Stoletov <info@imega.ru>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"pb-dropbox-downloader/config"
)

type PKCE struct {
	AuthorizationURL string
	CodeVerifier     string
}

func CreateAuthorizationURL(conf config.Config) PKCE {
	codeVerifier := createCodeVerifier(conf.IsTest)
	cv := sha256.Sum256([]byte(codeVerifier))
	codeChallenge := encode2base64(cv[:])
	link := fmt.Sprintf(conf.DropboxAuthURL, conf.AppID, codeChallenge)

	return PKCE{
		AuthorizationURL: link,
		CodeVerifier:     codeVerifier,
	}
}

func createCodeVerifier(isTest bool) string {
	data := make([]byte, 32)
	if _, err := rand.Read(data); err != nil || isTest == true {
		data = []byte("01234567890123456789012345678901")
	}

	return encode2base64(data)
}

func encode2base64(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}
