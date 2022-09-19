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
	"path/filepath"
	"pb-dropbox-downloader/datastorage"
	"pb-dropbox-downloader/dropbox"
	"pb-dropbox-downloader/synchroniser"

	"github.com/go-git/go-billy/v5/osfs"
	dropboxLib "github.com/tj/go-dropbox"
)

const DBFileName = "pb-dropbox-downloader.bin"

func (app *App) getSynchronizer(
	progressFn synchroniser.ProgressFn,
) *synchroniser.DropboxSynchroniser {
	dbConf := &dropboxLib.Config{
		AccessToken: app.conf.AccessToken,
		HTTPClient:  app.client,
	}

	dropboxClient := dropbox.NewClient(
		dropbox.WithGoDropbox(
			dropboxLib.New(dbConf),
		),
	)

	fs := osfs.New("")
	storage := datastorage.NewFileStorage(
		datastorage.WithFilesystem(fs),
		datastorage.WithConfigPath(
			filepath.Join(app.API.GetCacheDir(), DBFileName),
		),
	)

	return synchroniser.NewSynchroniser(
		synchroniser.WithStorage(storage),
		synchroniser.WithFileSystem(fs),
		synchroniser.WithDropboxClient(dropboxClient),
		synchroniser.WithProgress(progressFn),
		synchroniser.WithOutput(app.Logger.GetWriter()),
		// synchroniser.WithMaxParallelism(parallelism),
		// synchroniser.WithVersion(version),
	)
}
