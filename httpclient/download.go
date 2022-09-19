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

package httpclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type DownloadArgs struct {
	URL    string
	Client *http.Client
	Src    io.Writer
}

func Download(ctx context.Context, args DownloadArgs) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, args.URL, nil)
	if err != nil {
		return fmt.Errorf("failed to create a request, %w", err)
	}

	resp, err := args.Client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send a request, %w", err)
	}

	defer resp.Body.Close()

	if _, err := io.Copy(args.Src, resp.Body); err != nil {
		return fmt.Errorf("failed to copy a data, %w", err)
	}

	return nil
}
