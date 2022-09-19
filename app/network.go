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
	"os/exec"
	"pb-dropbox-downloader/httpclient"
)

func (app *App) openConnections() {
	err := exec.Command(app.API.GetPathNetAgent(), "net on").Run()
	if err != nil {
		app.Logger.Errorf("failed to turn on network, %s", err)
	}

	err = exec.Command(app.API.GetPathNetAgent(), "connect").Run()
	if err != nil {
		app.Logger.Errorf("failed to connect, %s", err)
	}
}

func (app *App) connect() bool {
	app.Logger.Debugf("initialize connection")

	for attempt := 0; attempt < 3; attempt++ {
		connetionStatus, err := httpclient.CheckConnection(
			app.client,
			app.conf.TestURL,
		)
		if err != nil {
			app.Logger.Debugf("CheckConnection: %s", err)
			app.openConnections()
		}

		if connetionStatus {
			return true
		}

		app.Logger.Infof("failed to connect to network, attempt: %d", attempt)
	}

	return false
}
