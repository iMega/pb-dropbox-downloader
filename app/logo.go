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

import "image"

func (app *App) showLogo() {
	size := app.API.ScreenSize()
	bm := app.API.GetLogo()
	logo := bm.Size()
	newImg := resizeImageByWidth(logo, size.X/2)
	x1 := size.X / 4

	app.API.DrawBitmap(x1, 0, bm)
	app.API.PartialUpdate(image.Rectangle{
		Min: image.Pt(x1, 0),
		Max: image.Pt(x1+newImg.X, newImg.Y),
	})
}

func resizeImageByWidth(src image.Point, width int) image.Point {
	ratio := float64(src.X / src.Y)

	return image.Pt(width, int(float64(width)/ratio))
}
