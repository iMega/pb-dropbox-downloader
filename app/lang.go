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
	"fmt"
	"io"
	"os"
	"path/filepath"
	"pb-dropbox-downloader/httpclient"
)

func (app *App) loadLang() (io.ReadCloser, error) {
	filename := filepath.Join(app.API.GetLangDir(), app.conf.Language+".ftl")
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file, %w", err)
	}

	app.Logger.Debugf("loaded a localization file")

	return file, nil
}

func (app *App) downloadLang() error {
	if err := os.MkdirAll(app.API.GetLangDir(), 0755); err != nil {
		return fmt.Errorf("failed to create a langdir, %w", err)
	}

	filename := filepath.Join(app.API.GetLangDir(), app.conf.Language+".ftl")
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create a config, %w", err)
	}

	ctx := context.Background()
	args := httpclient.DownloadArgs{
		URL:    fmt.Sprintf(app.conf.L10nURL, app.conf.Language),
		Client: app.client,
		Src:    file,
	}

	if err := httpclient.Download(ctx, args); err != nil {
		return fmt.Errorf("failed to download a file, %w", err)
	}

	app.Logger.Debugf("downloaded a localization file")

	return nil
}
