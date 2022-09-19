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
	"image"

	"github.com/skip2/go-qrcode"
)

func (app *App) showQRCode(url string) {
	size := app.API.ScreenSize()

	app.createQRCode(image.Pt(int(size.X/4), 10), size.X/2, url)
	app.API.PartialUpdate(image.Rectangle{
		Min: image.Pt(int(size.X/4), 10),
		Max: image.Pt(size.X/4*3, size.X+10),
	})
}

func (app *App) createQRCode(point image.Point, width int, data string) {
	qrc, err := qrcode.New(data, qrcode.Low)
	if err != nil {
		app.Logger.Errorf("failed to create a qrcode, %s", err)
	}

	img := qrc.Image(width)
	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			clr := img.At(x, y)
			app.API.DrawPixel(image.Pt(point.X+x, point.Y+y), clr)
			// if clr == color.Black {
			// app.API.DrawPixel(image.Pt(point.X+x, point.Y+y), color.Black)
			// }
		}
	}
}
