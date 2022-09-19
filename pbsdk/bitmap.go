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

package pbsdk

/*
#include "inkview.h"
#cgo CFLAGS: -pthread
#cgo LDFLAGS: -pthread -lpthread -linkview
*/
import "C" //nolint:typecheck
import "image"

type Bitmap struct {
	Src *C.ibitmap
}

func (bm *Bitmap) Size() image.Point {
	return image.Pt(int(bm.Src.width), int(bm.Src.height))
}

func LoadBitmap(filename string) *Bitmap {
	bm := C.LoadBitmap(C.CString(filename))
	if bm == nil {
		return nil
	}

	return &Bitmap{Src: bm}
}

func DrawBitmap(x, y int, bm *Bitmap) {
	C.DrawBitmap(C.int(x), C.int(y), bm.Src)
}

func DrawBitmapRect(bm *Bitmap, x, y, w, h, flags int) {
	C.DrawBitmapRect(C.int(x), C.int(y), C.int(w), C.int(h), bm.Src, C.int(flags))
}

func StretchBitmap(bm *Bitmap, x, y, w, h, flags int) {
	C.StretchBitmap(C.int(x), C.int(y), C.int(w), C.int(h), bm.Src, C.int(flags))
}
