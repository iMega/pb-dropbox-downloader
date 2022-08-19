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
	"errors"
	"fmt"
	"net"
	"net/http"
	"pb-dropbox-downloader/logger"
	"time"
)

func New(conf Config, log *logger.Logger) *http.Client {
	return &http.Client{
		Timeout: conf.Timeout,
		Transport: &userAgent{
			Value: conf.UserAgent,
			RoundTriper: &retrier{
				Retrier: NewDefaultRetrier(conf),
				Log:     log,
				RoundTriper: &http.Transport{
					MaxIdleConns:        conf.MaxIdleConns,
					MaxConnsPerHost:     conf.MaxConnsPerHost,
					MaxIdleConnsPerHost: conf.MaxIdleConnsPerHost,
					DialContext: (&net.Dialer{
						Timeout: conf.DialerTimeout,
					}).DialContext,
				},
			},
		},
	}
}

var errTemporaryNetworkProblem = errors.New("temporary network problem")

type retrier struct {
	Retrier     *Retrier
	Log         *logger.Logger
	RoundTriper http.RoundTripper
}

func (t *retrier) RoundTrip(req *http.Request) (*http.Response, error) {
	var resp *http.Response

	operation := func() error {
		r, err := t.RoundTriper.RoundTrip(req)
		if err != nil {
			return fmt.Errorf("failed to call TransportWithRetrier, %w", err)
		}

		if r.StatusCode >= http.StatusLocked {
			return errTemporaryNetworkProblem
		}

		resp = r

		return nil
	}

	notifyFn := func(err error, next time.Duration) {
		t.Log.Infof("%s, retrying in %s...", err, next)
	}

	if err := t.Retrier.Retry(req.Context(), operation, notifyFn); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return resp, nil
}
