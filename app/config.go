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
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"pb-dropbox-downloader/config"
	"pb-dropbox-downloader/httpclient"
)

func (app *App) loadConfig() {
	filename := filepath.Join(app.API.GetConfigDir(), FileConfig)
	conf, err := config.Load(app.API.GetGlobalConfigFilename(), filename)
	if err != nil {
		app.error("failed to load a config, %s", err)
	}

	rawConf, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		app.Logger.Errorf("failed to marshal a config, %s", err)
	}

	app.Logger.Debugf("Loaded config:\n%s", string(rawConf))

	app.conf = conf
	app.conf.IsTest = app.IsTest
}

func (app *App) downloadConfig() error {
	file, err := os.Create(filepath.Join(app.API.GetConfigDir(), FileConfig))
	if err != nil {
		return fmt.Errorf("failed to create a config, %w", err)
	}

	ctx := context.Background()
	args := httpclient.DownloadArgs{
		URL:    app.conf.ConfigURL,
		Client: app.client,
		Src:    file,
	}

	if err := httpclient.Download(ctx, args); err != nil {
		return fmt.Errorf("failed to download a file, %w", err)
	}

	app.Logger.Debugf("downloaded a config file")

	return nil
}
