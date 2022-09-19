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

package bridge

import (
	"image"
	"image/color"
	"pb-dropbox-downloader/api"
	"pb-dropbox-downloader/pbsdk"

	ink "github.com/dennwc/inkview"
)

type Bridge struct{}

func (Bridge) GetFlashDir() string     { return pbsdk.FlashDir }
func (Bridge) GetSDCardDir() string    { return pbsdk.SDCardDir }
func (Bridge) GetConfigDir() string    { return pbsdk.ConfigDir }
func (Bridge) GetLangDir() string      { return pbsdk.LangDir }
func (Bridge) GetCacheDir() string     { return pbsdk.CacheDir }
func (Bridge) GetPathNetAgent() string { return pbsdk.NetAgent }
func (Bridge) GetAppDir() string       { return pbsdk.AppDir }

func (Bridge) GetGlobalConfigFilename() string { return pbsdk.GlobalConfig }

func (Bridge) GetKeyboardNumeric() int { return int(pbsdk.KeyboardNumeric) }

func (Bridge) OpenKeyboard(title, buf string, keyboard int) <-chan string {
	ch := make(chan string)
	// defer close(ch)

	pbsdk.OpenKeyboard(title, buf, pbsdk.Keyboard(keyboard), ch)

	return ch
}

func (Bridge) ScreenSize() image.Point { return ink.ScreenSize() }

func (Bridge) GetLogo() api.Bitmaper { return pbsdk.GetLogo() }

func (Bridge) DrawBitmap(x, y int, bm api.Bitmaper) {
	if val, ok := bm.(*pbsdk.Bitmap); ok {
		pbsdk.DrawBitmap(x, y, val)
	}
}

func (Bridge) DrawPixel(p image.Point, cl color.Color) { ink.DrawPixel(p, cl) }

func (Bridge) PartialUpdate(r image.Rectangle) { ink.PartialUpdate(r) }

func (Bridge) OpenProgressbar(title, text string, icon, percent int) {
	pbsdk.OpenProgressbar(title, text, icon, percent)
}

func (Bridge) UpdateProgressbar(text string, percent int) {
	pbsdk.UpdateProgressbar(text, percent)
}

func (Bridge) CloseProgressbar() { pbsdk.CloseProgressbar() }
