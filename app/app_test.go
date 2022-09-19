package app_test

import (
	"context"
	"errors"
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"pb-dropbox-downloader/api"
	"pb-dropbox-downloader/app"
	"pb-dropbox-downloader/logger"
	"pb-dropbox-downloader/translations"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type CheckList struct {
	CheckConnection       bool
	DownloadConfigFile    bool
	DownloadTranslateFile bool
	ExchangePinToCode     bool
	GetCurrentAccount     bool
	GetListFolder         bool
	DownloadBook          bool
}

func Test_FirstLaunch(t *testing.T) {
	expected := CheckList{
		CheckConnection:       true,
		DownloadConfigFile:    false,
		DownloadTranslateFile: true,
		ExchangePinToCode:     true,
		GetCurrentAccount:     true,
		GetListFolder:         true,
		DownloadBook:          true,
	}
	actual := CheckList{}

	rootDir := helperPrepareRootDir(t)
	// fmt.Println(rootDir)
	defer os.RemoveAll(rootDir)

	ctx, cancel := context.WithCancel(context.Background())

	testServer := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			fmt.Println(r.RequestURI)
			if r.RequestURI == "/healthcheck" {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			if r.RequestURI == "/generate_204" {
				actual.CheckConnection = true
				w.WriteHeader(http.StatusNoContent)
				return
			}

			if strings.Contains(r.RequestURI, app.FileConfig) {
				actual.DownloadConfigFile = true
				w.Write([]byte("AppID=0000\n"))
				return
			}

			if strings.Contains(r.RequestURI, "/translations/en.ftl") {
				actual.DownloadTranslateFile = true
				w.Write([]byte(translations.Fallback()))
				return
			}

			if r.RequestURI == "/oauth2/token" {
				w.Write([]byte(`{"access_token":"AccessToken","expires_in":14000,"refresh_token":"RefreshToken"}`))
				return
			}

			if r.RequestURI == "/?pin=0000" {
				actual.ExchangePinToCode = true
				w.Write([]byte(`code`))
				return
			}

			if strings.Contains(r.RequestURI, "users/get_current_account") {
				actual.GetCurrentAccount = true
				w.Write([]byte(`{
						"account_id": "account_id",
						"name": {
							"given_name":"given_name",
							"surname":"surname",
							"familiar_name":"familiar_name",
							"display_name":"display_name"
						},
						"email":"email@example.com",
						"locale":"locale",
						"referral_link":"referral_link",
						"is_paired":true,
						"account_type": {".tag":"tag"},
						"country":"country"
					}`))
				return
			}

			if strings.Contains(r.RequestURI, "files/list_folder") {
				actual.GetListFolder = true
				w.Write([]byte(`{
						"entries": [
						  {
							".tag": "folder",
							"name": "Приложения",
							"path_lower": "/приложения",
							"path_display": "/Приложения",
							"id": "id:oaKCdDCZFMAAAAAAAAADJg"
						  },
						  {
							".tag": "folder",
							"name": "pb-downloader",
							"path_lower": "/приложения/pb-downloader",
							"path_display": "/Приложения/pb-downloader",
							"id": "id:oaKCdDCZFMAAAAAAAAADKA"
						  },
						  {
							".tag": "file",
							"name": "Том ДеМарко - Deadline.fb2.epub",
							"path_lower": "/deadline.fb2.epub",
							"path_display": "/Том ДеМарко - Deadline.fb2.epub",
							"id": "id:oaKCdDCZFMAAAAAAAAACIw",
							"client_modified": "2022-06-22T18:27:25Z",
							"server_modified": "2022-07-31T21:44:44Z",
							"rev": "5e520ca94a4d1461c7734",
							"size": 643393,
							"is_downloadable": true,
							"content_hash": "ac02b13f8c7dedaaec48242ffc1bc33078e7df72a2b4cc240b508215f61ef3c4"
						  }
						],
						"cursor": "AAHO",
						"has_more": true
					  }
					`))
			}

			if strings.Contains(r.RequestURI, "files/download") {
				actual.DownloadBook = true
				w.Write([]byte(`file`))
				cancel()
			}
		},
	))
	defer testServer.Close()

	require.NoError(t, waitForSystemUnderTestReady(t, testServer.URL))

	mb := &mockBridge{RootDir: rootDir}

	err := os.MkdirAll(mb.GetConfigDir(), os.ModePerm)
	require.NoError(t, err)

	// buf := bytes.NewBuffer([]byte("AppID=0000\n"))
	// helperCreateFile(t, filepath.Join(mb.GetConfigDir(), app.FileConfig), buf)

	log := logger.New("DEBUG", os.Stdout)
	instance := &app.App{
		Logger:   log,
		API:      mb,
		IsTest:   true,
		TestAddr: testServer.Listener.Addr().String(),
	}

	err = instance.Init(ctx)
	require.NoError(t, err)

	for {
		select {
		case <-ctx.Done():
			instance.Close()
			assert.Equal(t, expected, actual)
			return
		case <-time.After(time.Second):
			t.Fatal("timeout")
			return
		}
	}
}

func helperPrepareRootDir(t *testing.T) string {
	val, err := os.MkdirTemp("", "dropbox_test")
	require.NoError(t, err)

	return val
}

func helperCreateFile(t *testing.T, filename string, src io.Reader) {
	file, err := os.Create(filename)
	require.NoError(t, err)

	_, err = io.Copy(file, src)
	require.NoError(t, err)
}

func waitForSystemUnderTestReady(t *testing.T, url string) error {
	req, err := http.NewRequest(http.MethodGet, url+"/healthcheck", nil)
	require.NoError(t, err)

	for attempts := 30; attempts > 0; attempts-- {
		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		if err == nil && resp != nil && resp.StatusCode == http.StatusNoContent {
			return nil
		}

		log.Printf("ATTEMPTING TO CONNECT")

		<-time.After(time.Second)
	}

	return errors.New("SUT is not ready for tests")
}

type mockBridge struct {
	RootDir string
}

func (mock *mockBridge) GetFlashDir() string {
	return filepath.Join(mock.RootDir, "/mnt/ext1")
}
func (mock *mockBridge) GetSDCardDir() string {
	return filepath.Join(mock.RootDir, "/mnt/ext2")
}

func (mock *mockBridge) GetConfigDir() string {
	return filepath.Join(mock.GetFlashDir(), "/system", "/config")
}

func (mock *mockBridge) GetLangDir() string {
	return filepath.Join(mock.GetFlashDir(), "/system", "/language")
}

func (mock *mockBridge) GetCacheDir() string {
	return filepath.Join(mock.GetFlashDir(), "/system", "/cache")
}

func (mock *mockBridge) GetPathNetAgent() string {
	return filepath.Join(mock.RootDir, "/ebrmain", "/bin/netagent")
}
func (mock *mockBridge) GetAppDir() string {
	return filepath.Join(mock.GetFlashDir(), "/applications")
}

func (mock *mockBridge) GetGlobalConfigFilename() string {
	return filepath.Join(mock.GetConfigDir(), "/global.cfg")
}

func (mock *mockBridge) GetKeyboardNumeric() int { return 0 }

func (mock *mockBridge) OpenKeyboard(string, string, int) <-chan string {
	ch := make(chan string, 1)
	defer close(ch)

	ch <- "0000"

	return ch
}

func (mock *mockBridge) ScreenSize() image.Point {
	return image.Pt(800, 1024)
}

func (mock *mockBridge) GetLogo() api.Bitmaper { return &mockBitmap{} }

func (mock *mockBridge) DrawBitmap(x, y int, bm api.Bitmaper) {}

func (mock *mockBridge) DrawPixel(p image.Point, cl color.Color) {}

func (mock *mockBridge) PartialUpdate(r image.Rectangle) {}

func (mock *mockBridge) OpenProgressbar(string, string, int, int) {}

func (mock *mockBridge) UpdateProgressbar(string, int) {}

func (mock *mockBridge) CloseProgressbar() {}

type mockBitmap struct{}

func (*mockBitmap) Size() image.Point { return image.Pt(100, 100) }
