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

package config

import (
	"fmt"
	"time"

	"gopkg.in/ini.v1"
)

type Config struct {
	cfg      *ini.File
	filename string

	Language string `ini:",omitempty"`
	Timezone string `ini:",omitempty"`
	Font     string `ini:",omitempty"`

	ConfigURL string
	L10nURL   string
	TestURL   string

	AppID           string
	DropboxAuthURL  string
	DropboxTokenURL string
	RedirectURL     string

	RefreshToken         string
	AccessToken          string
	AccessTokenExpiresIn time.Time

	Timeout               time.Duration
	MaxIdleConns          int
	MaxConnsPerHost       int
	MaxIdleConnsPerHost   int
	DialerTimeout         time.Duration
	BackoffMaxInterval    time.Duration
	BackoffMaxElapsedTime time.Duration

	// IsShowStatusBar bool
	HasSDCard bool `ini:"-"`
	IsTest    bool `ini:"-"`
}

func Load(global, local string) (*Config, error) {
	conf := defaultConfig()
	conf.filename = local

	opts := ini.LoadOptions{Loose: true, Insensitive: true}
	cfg, err := ini.LoadSources(opts, global, local)
	if err != nil {
		return nil, fmt.Errorf("failed to load config, %w", err)
	}

	conf.cfg = cfg

	if err := cfg.Section("").MapTo(conf); err != nil {
		return nil, fmt.Errorf("failed to assign value, %w", err)
	}

	return conf, nil
}

func (conf *Config) Save() error {
	cfg := ini.Empty(ini.LoadOptions{InsensitiveKeys: true})
	if err := ini.ReflectFrom(cfg, conf); err != nil {
		return fmt.Errorf("failed to reflect a config, %w", err)
	}

	// reset duplicates from global config
	cfg.Section("").DeleteKey("language")
	cfg.Section("").DeleteKey("font")
	cfg.Section("").DeleteKey("timezone")

	if err := cfg.SaveTo(conf.filename); err != nil {
		return fmt.Errorf("failed to write a file, %w", err)
	}

	return nil
}

func defaultConfig() *Config {
	repo := "https://raw.githubusercontent.com/imega/pb-dropbox-downloader/main"
	return &Config{
		AppID:    "fuwgr8q8src9lk8",
		Language: "en",
		L10nURL:  repo + "/translations/%s.ftl",
		DropboxAuthURL: "https://www.dropbox.com/oauth2/authorize?" +
			"response_type=code&code_challenge_method=S256&" +
			"token_access_type=offline&redirect_uri=https://pb.imega.ru" +
			"&client_id=%s&code_challenge=%s",
		DropboxTokenURL: "https://api.dropbox.com/oauth2/token",
		ConfigURL:       repo + "/config/pb-dropbox-downloader.cfg",
		TestURL:         "http://clients3.google.com/generate_204",
		RedirectURL:     "https://pb.imega.ru",

		Timeout:               15 * time.Second,
		MaxIdleConns:          100,
		MaxConnsPerHost:       100,
		MaxIdleConnsPerHost:   100,
		DialerTimeout:         15 * time.Second,
		BackoffMaxInterval:    300 * time.Millisecond,
		BackoffMaxElapsedTime: 10 * time.Second,
	}
}
