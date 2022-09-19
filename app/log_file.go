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
	"fmt"
	"os"
	"path/filepath"
)

const logFileName = "pb-dropbox-downloader.txt"

func CreateLogFile(path string) (*os.File, error) {
	file, err := os.OpenFile(
		filepath.Join(path, logFileName),
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0755,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file, %w", err)
	}

	return file, nil
}
