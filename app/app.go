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

package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"path/filepath"
	"pb-dropbox-downloader/api"
	"pb-dropbox-downloader/auth"
	"pb-dropbox-downloader/config"
	"pb-dropbox-downloader/httpclient"
	"pb-dropbox-downloader/l10n"
	"pb-dropbox-downloader/logger"
	"time"

	"golang.org/x/text/message"
)

const FileConfig = "pb-dropbox-downloader.cfg"

var ErrConn = errors.New("failed to connect the network, try connecting manually")

type App struct {
	Logger *logger.Logger
	API    api.Pocketbook

	TestAddr string
	IsTest   bool

	ctx     context.Context
	cancel  context.CancelFunc
	client  *http.Client
	conf    *config.Config
	printer *message.Printer

	syncCh         chan struct{}
	accessTokenCh  chan struct{}
	refreshTokenCh chan struct{}
}

func (app *App) Init(ctx context.Context) error {
	newCtx, cancel := context.WithCancel(ctx)
	app.ctx = newCtx
	app.cancel = cancel

	app.Logger.Debugf("Starting an application")

	app.syncCh = make(chan struct{}, 1)
	app.accessTokenCh = make(chan struct{}, 1)
	app.refreshTokenCh = make(chan struct{}, 1)

	app.showLogo()

	app.loadConfig()
	app.initHTTPClient()

	if !app.connect() {
		app.error(ErrConn.Error())
	}

	// TODO: to inititialize certificates
	// if err := ink.InitCerts(); err != nil {
	// 	app.error("failed to inititialize certificates, %s", err)
	// }

	if app.conf.AppID == "" {
		if err := app.downloadConfig(); err != nil {
			app.error("failed to download a config, %s", err)
		}

		app.loadConfig()
		app.initHTTPClient()
	}

	app.initLocalization()

	if !app.IsTest {
		app.hasSDCard()
	}
	app.Logger.Debugf("Has SDCard: %v", app.conf.HasSDCard)

	go app.process(app.ctx)

	app.syncCh <- struct{}{}

	return nil
}

func (app *App) Close() {
	app.Logger.Infof("======CLOSE======")
	app.cancel()
}

func (app *App) Sync() {
	app.Logger.Infof("============")
	// app.syncCh <- struct{}{}
}

func (app *App) error(format string, args ...interface{}) {
	app.Logger.Errorf(format, args...)
	// Panic is used here because the application cannot send
	// an error message. This code prevents the error from being displayed.
	// https://github.com/dennwc/inkview/blob/8f626286e32bb7cc2c42fce86acfb60786e1d496/inkview.go#L54
	panic(fmt.Errorf(format, args...))
}

func (app *App) initHTTPClient() {
	httpClientConf := httpclient.Config{
		UserAgent:             "github.com/evg4b/pb-dropbox-downloader",
		Timeout:               app.conf.Timeout,
		MaxIdleConns:          app.conf.MaxIdleConns,
		MaxConnsPerHost:       app.conf.MaxConnsPerHost,
		MaxIdleConnsPerHost:   app.conf.MaxIdleConnsPerHost,
		DialerTimeout:         app.conf.DialerTimeout,
		BackoffMaxInterval:    app.conf.BackoffMaxInterval,
		BackoffMaxElapsedTime: app.conf.BackoffMaxElapsedTime,
	}

	if app.IsTest {
		httpClientConf.TestHost = app.TestAddr
	}

	app.client = httpclient.New(httpClientConf, app.Logger)

	app.Logger.Debugf("initialize http-client")
}

func (app *App) initLocalization() {
	var rLang io.ReadCloser

	rLang, err := app.loadLang()
	if err != nil {
		app.Logger.Errorf("failed to load a lang, %s", err)
	}

	if errors.Is(err, fs.ErrNotExist) {
		if err := app.downloadLang(); err != nil {
			app.Logger.Errorf("failed to download a lang, %s", err)
		}

		rLang, err = app.loadLang()
		if err != nil {
			app.Logger.Errorf("failed to load a lang, %s", err)
		}
	}

	printer, err := l10n.New(app.conf.Language, rLang)
	if err != nil {
		app.Logger.Errorf("failed to create a lang-printer, %s", err)
	}

	app.printer = printer
}

func (app *App) process(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			app.Logger.Debugf("process: context done")

			if err := ctx.Err(); err != nil {
				app.Logger.Errorf("failed to cancel, %s", err)
			}

			return

		case _, ok := <-app.syncCh:
			app.Logger.Debugf("process: sync")

			if !ok {
				app.Logger.Errorf("channel was closed")

				return
			}

			if app.conf.AccessTokenExpiresIn.Before(time.Now()) || app.conf.AccessTokenExpiresIn.IsZero() {
				app.accessTokenCh <- struct{}{}

				continue
			}

			if !app.connect() {
				app.error(ErrConn.Error())
			}

			app.API.OpenProgressbar("Sync", "", 1, 0)

			synchronizer := app.getSynchronizer(
				func(text string, current, total int) {
					percent := 100 * current / total
					app.Logger.Debugf("downloaded: %s (%d)", text, percent)

					app.API.UpdateProgressbar(text, percent)
				},
			)

			folder := filepath.Join(app.API.GetFlashDir(), "dropbox")
			if app.conf.HasSDCard {
				folder = filepath.Join(app.API.GetSDCardDir(), "dropbox")
			}

			app.Logger.Debugf("start sync")

			if err := synchronizer.Sync(context.Background(), folder, true); err != nil {
				app.Logger.Errorf("failed to sync, %s", err)
			}

			app.API.CloseProgressbar()

			app.Logger.Debugf("sync complete")

		case _, ok := <-app.accessTokenCh:
			app.Logger.Debugf("process: getting access token")

			if !ok {
				app.Logger.Errorf("channel was closed")

				return
			}

			if app.conf.RefreshToken == "" {
				app.refreshTokenCh <- struct{}{}
			}

			if !app.connect() {
				app.error(ErrConn.Error())
			}

			app.getAccessToken(ctx)

		case _, ok := <-app.refreshTokenCh:
			app.Logger.Debugf("process: getting refresh token")

			if !ok {
				app.Logger.Errorf("channel was closed")

				return
			}

			app.getRefreshToken(ctx)
		}
	}
}

func (app *App) getAccessToken(ctx context.Context) {
	authClient := auth.New(app.conf.DropboxTokenURL, app.client)
	args := auth.TokenParameters{
		ClientID:     app.conf.AppID,
		RefreshToken: app.conf.RefreshToken,
	}

	token, err := authClient.GetAccessToken(ctx, args)
	if err != nil {
		app.error("%s", err)

		app.refreshTokenCh <- struct{}{}

		return
	}

	app.conf.AccessToken = token.AccessToken
	app.conf.AccessTokenExpiresIn = token.ExpiresIn.Date

	if err := app.conf.Save(); err != nil {
		app.Logger.Errorf("failed to save a config, %s", err)

		app.refreshTokenCh <- struct{}{}
	}

	app.Logger.Debugf("got a access token")

	app.syncCh <- struct{}{}
}

func (app *App) getRefreshToken(ctx context.Context) {
	pkce := auth.CreateAuthorizationURL(app.conf)

	app.showQRCode(pkce.AuthorizationURL)

	authClient := auth.New(app.conf.DropboxTokenURL, app.client)

	keyboard := app.API.GetKeyboardNumeric()
	title := "EnterPin"
	if app.conf.RedirectURL == "" {
		keyboard = app.API.GetKeyboardNumeric()
		title = "EnterCode"
	}

	keyboardInput := app.API.OpenKeyboard(
		app.printer.Sprint(title),
		"",
		keyboard,
	)

	go func() {
		app.Logger.Debugf("waiting input from user")

		input := <-keyboardInput

		app.Logger.Debugf("input is %s", input)

		if app.conf.RedirectURL != "" {
			code, err := app.exchangePinToCode(ctx, input)
			if err != nil {
				app.Logger.Errorf("failed to exchange a pin, %s", err)

				app.refreshTokenCh <- struct{}{}

				return
			}

			input = code
		}

		args := auth.TokenParameters{
			Code:         input,
			ClientID:     app.conf.AppID,
			CodeVerifier: pkce.CodeVerifier,
			RedirectURL:  app.conf.RedirectURL,
		}

		token, err := authClient.GetRefreshToken(ctx, args)
		if err != nil {
			app.error("%s", err)

			app.refreshTokenCh <- struct{}{}

			return
		}

		app.conf.RefreshToken = token.RefreshToken
		app.conf.AccessToken = token.AccessToken
		app.conf.AccessTokenExpiresIn = token.ExpiresIn.Date

		if err := app.conf.Save(); err != nil {
			app.Logger.Errorf("failed to save a config, %s", err)

			app.refreshTokenCh <- struct{}{}
		}

		app.Logger.Debugf("got a refresh token")

		app.syncCh <- struct{}{}
	}()
}
