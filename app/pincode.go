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
	"bytes"
	"context"
	"fmt"
	"os"
	"pb-dropbox-downloader/httpclient"
)

func (app *App) exchangePinToCode(
	ctx context.Context,
	pin string,
) (string, error) {
	buf := bytes.NewBuffer(nil)

	args := httpclient.DownloadArgs{
		URL:    app.conf.RedirectURL + "?pin=" + pin,
		Client: app.client,
		Src:    buf,
	}

	if err := httpclient.Download(ctx, args); err != nil {
		return "", fmt.Errorf("failed to get code, %w", err)
	}

	if buf.Len() == 0 {
		return "", os.ErrInvalid
	}

	return buf.String(), nil
}
