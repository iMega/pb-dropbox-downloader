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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	URL   string
	inner *http.Client
}

func New(url string, inner *http.Client) *Client {
	return &Client{URL: url, inner: inner}
}

func (clt *Client) GetAccessToken(
	ctx context.Context,
	args TokenParameters,
) (ResponseToken, error) {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("client_id", args.ClientID)
	data.Set("refresh_token", args.RefreshToken)

	resp, err := clt.do(ctx, data)
	if err != nil {
		return resp, fmt.Errorf("failed to get the access token, %w", err)
	}

	return resp, nil
}

func (clt *Client) GetRefreshToken(
	ctx context.Context,
	args TokenParameters,
) (ResponseToken, error) {
	data := url.Values{}
	data.Set("code", args.Code)
	data.Set("grant_type", "authorization_code")
	data.Set("code_verifier", args.CodeVerifier)
	data.Set("client_id", args.ClientID)

	resp, err := clt.do(ctx, data)
	if err != nil {
		return resp, fmt.Errorf("failed to get the refresh token, %w", err)
	}

	return resp, nil
}

type TokenParameters struct {
	Code         string
	CodeVerifier string
	ClientID     string
	RefreshToken string
}

var errReturnError = errors.New("returns error")

func (clt *Client) do(
	ctx context.Context,
	val url.Values,
) (ResponseToken, error) {
	responseToken := ResponseToken{}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		clt.URL,
		strings.NewReader(val.Encode()),
	)
	if err != nil {
		return responseToken, fmt.Errorf("failed to create a request, %w", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := clt.inner.Do(req)
	if err != nil {
		return responseToken, fmt.Errorf("failed to send a request, %w", err)
	}

	if err := json.NewDecoder(resp.Body).Decode(&responseToken); err != nil {
		return responseToken, fmt.Errorf("failed to decode a response, %w", err)
	}

	if responseToken.Error != "" {
		return responseToken,
			fmt.Errorf("%w: %s", errReturnError, responseToken.Error)
	}

	return responseToken, nil
}

type ResponseToken struct {
	AccessToken  string    `json:"access_token"`
	ExpiresIn    ExpiresIn `json:"expires_in"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	Error        string    `json:"error_description,omitempty"`
}

type ExpiresIn struct {
	Date time.Time
}

func (exp *ExpiresIn) UnmarshalJSON(data []byte) error {
	var num int

	if err := json.Unmarshal(data, &num); err != nil {
		return fmt.Errorf("failed to unmarshal ExpiresIn, %w", err)
	}

	exp.Date = time.Now().Add(14400 * time.Second)

	return nil
}
