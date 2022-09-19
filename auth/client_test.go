package auth_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"pb-dropbox-downloader/auth"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetAccessToken_ReturnsToken(t *testing.T) {
	args := auth.TokenParameters{
		Code:         "code",
		CodeVerifier: "code-verifier",
		ClientID:     "client-id",
	}

	testServer := httptest.NewServer(
		http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			expected := url.Values{
				"grant_type":    []string{"refresh_token"},
				"client_id":     []string{args.ClientID},
				"refresh_token": []string{args.RefreshToken},
			}

			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t,
				req.Header.Get("Content-Type"),
				"application/x-www-form-urlencoded",
			)
			require.NoError(t, req.ParseForm())
			require.Equal(t, expected, req.Form)

			res.Write([]byte(`{
				"access_token": "sl.BNolEfsF",
				"token_type": "bearer",
				"expires_in": 14400
			  }`))
		}),
	)
	defer func() { testServer.Close() }()

	client := auth.New(testServer.URL, http.DefaultClient)

	ctx := context.Background()
	actual, err := client.GetAccessToken(ctx, args)
	require.NoError(t, err)

	expected := auth.ResponseToken{
		AccessToken: "sl.BNolEfsF",
		ExpiresIn: auth.ExpiresIn{
			Date: time.Now().Add(14400 * time.Second).Truncate(time.Millisecond),
		},
	}

	actual.ExpiresIn.Date = actual.ExpiresIn.Date.Truncate(time.Millisecond)

	assert.Equal(t, expected, actual)
}

func Test_GetAccessToken_ReturnsError(t *testing.T) {
	testServer := httptest.NewServer(
		http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t,
				req.Header.Get("Content-Type"),
				"application/x-www-form-urlencoded",
			)

			res.Write([]byte(`{
				"error": "invalid_grant",
				"error_description": "refresh token is malformed"
			  }`))
		}),
	)
	defer func() { testServer.Close() }()

	client := auth.New(testServer.URL, http.DefaultClient)
	_, err := client.GetAccessToken(
		context.Background(),
		auth.TokenParameters{},
	)

	require.ErrorIs(t, err, auth.ErrReturnError)
}

func Test_GetRefreshToken_ReturnsToken(t *testing.T) {
	args := auth.TokenParameters{
		Code:         "code",
		CodeVerifier: "code-verifier",
		ClientID:     "client-id",
	}

	testServer := httptest.NewServer(
		http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			expected := url.Values{
				"grant_type":    []string{"authorization_code"},
				"client_id":     []string{args.ClientID},
				"code":          []string{args.Code},
				"code_verifier": []string{args.CodeVerifier},
			}

			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t,
				req.Header.Get("Content-Type"),
				"application/x-www-form-urlencoded",
			)
			require.NoError(t, req.ParseForm())
			require.Equal(t, expected, req.Form)

			res.Write([]byte(`{
				"access_token": "sl.BNrSdT9WP",
				"token_type": "bearer",
				"expires_in": 14400,
				"refresh_token": "uD_1XUXxy3Q",
				"scope": "account_info.read",
				"uid": "5472",
				"account_id": "dbid:AACLJHBBHU"
			  }`))
		}),
	)
	defer func() { testServer.Close() }()

	client := auth.New(testServer.URL, http.DefaultClient)

	ctx := context.Background()

	actual, err := client.GetRefreshToken(ctx, args)
	require.NoError(t, err)

	expected := auth.ResponseToken{
		AccessToken: "sl.BNrSdT9WP",
		ExpiresIn: auth.ExpiresIn{
			Date: time.Now().Add(14400 * time.Second).Truncate(time.Millisecond),
		},
		RefreshToken: "uD_1XUXxy3Q",
	}

	actual.ExpiresIn.Date = actual.ExpiresIn.Date.Truncate(time.Millisecond)

	assert.Equal(t, expected, actual)
}

func Test_GetRefreshToken_ReturnsError(t *testing.T) {
	testServer := httptest.NewServer(
		http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t,
				req.Header.Get("Content-Type"),
				"application/x-www-form-urlencoded",
			)

			res.Write([]byte(`{
				"error": "invalid_grant",
				"error_description": "code has already been used"
			  }`))
		}),
	)
	defer func() { testServer.Close() }()

	client := auth.New(testServer.URL, http.DefaultClient)
	_, err := client.GetRefreshToken(context.Background(), auth.TokenParameters{})

	require.ErrorIs(t, err, auth.ErrReturnError)
}
