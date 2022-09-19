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

package main

import (
	"image/png"
	"log"
	"os"

	"golang.org/x/image/bmp"
)

func main() {
	src, err := os.Open("logo.png")
	if err != nil {
		log.Fatalf("failed to open a src-file, %s", err)
	}
	defer src.Close()

	dst, err := os.Create("logo.bmp")
	if err != nil {
		log.Fatalf("failed to create a dst-file, %s", err)
	}
	defer dst.Close()

	img, err := png.Decode(src)
	if err != nil {
		log.Fatalf("failed to decode an image, %s", err)
	}

	if err := bmp.Encode(dst, img); err != nil {
		log.Fatalf("failed to encode an image, %s", err)
	}

	log.Printf("Done!")
}
