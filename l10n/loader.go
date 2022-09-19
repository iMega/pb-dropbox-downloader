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

package l10n

import (
	"fmt"
	"io"
	"pb-dropbox-downloader/translations"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
	"gopkg.in/ini.v1"
)

func New(lang string, r io.ReadCloser) (*message.Printer, error) {
	lngTag, err := language.Parse(lang)
	if err != nil {
		return nil, fmt.Errorf("failed to parse a lang code, %w", err)
	}

	if r == nil {
		r = io.NopCloser(strings.NewReader(""))
	}

	cat, err := loadCatalog(lngTag, r)
	if err != nil {
		return nil, fmt.Errorf("failed to load a lang-catalog, %w", err)
	}

	message.DefaultCatalog = cat
	printer := message.NewPrinter(lngTag)

	return printer, nil
}

func loadCatalog(lang language.Tag, rd io.ReadCloser) (catalog.Catalog, error) {
	fallbackDict, err := loadDict(strings.NewReader(translations.Fallback()))
	if err != nil {
		return nil, fmt.Errorf("failed to load fallback dictionary, %w", err)
	}

	dict := map[string]catalog.Dictionary{
		language.English.String(): fallbackDict,
	}

	currentDict, err := loadDict(rd)
	if err != nil {
		return nil, fmt.Errorf("failed to load current dictionary, %w", err)
	}

	dict[lang.String()] = currentDict

	cat, err := catalog.NewFromMap(dict, catalog.Fallback(language.English))
	if err != nil {
		return nil, fmt.Errorf("failed to create catalog, %w", err)
	}

	return cat, nil
}

func loadDict(src interface{}) (*dictionary, error) {
	file, err := ini.Load(src)
	if err != nil {
		return nil, fmt.Errorf("failed to load dictionary, %w", err)
	}

	mapDict := file.Section("").KeysHash()

	return &dictionary{Src: mapDict}, nil
}
