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

package main

import (
	"fmt"
	"os/exec"
	"pb-dropbox-downloader/app"
	"pb-dropbox-downloader/bridge"
	"pb-dropbox-downloader/lifecycle"
	"pb-dropbox-downloader/logger"
	"pb-dropbox-downloader/pbsdk"

	ink "github.com/dennwc/inkview"
)

const dialogIconError = "4"

func main() {
	defer func() {
		if p := recover(); p != nil {
			err := exec.Command(
				"/ebrmain/bin/dialog",
				dialogIconError,
				"",
				fmt.Sprintf("Application error:\n%s", p),
			).Run()
			if err != nil {
				panic(err.Error())
			}
		}
	}()

	logFile, err := app.CreateLogFile(pbsdk.FlashDir)
	if err != nil {
		panic(err.Error())
	}

	brg := bridge.Bridge{}

	err = ink.Run(&lifecycle.Lifecycle{
		Wrapped: &app.App{
			Logger: logger.New("DEBUG", logFile),
			API:    brg,
		},
	})
	if err != nil {
		panic(err.Error())
	}
}
