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

void iv_keyboardhandler_cb(char *text);
*/
import "C"

var keyboardInput chan string

// Keyboard is a code for keyboard.
type Keyboard int

const (
	KeyboardNumeric Keyboard = Keyboard(C.KBD_NUMERIC)

	keyboardStringLength = 4
)

func OpenKeyboard(title string, buf string, kbrd Keyboard, input chan string) {
	keyboardInput = input
	C.OpenKeyboard(
		C.CString(title),
		C.CString(buf),
		keyboardStringLength,
		C.int(kbrd),
		C.iv_keyboardhandler((*[0]byte)(C.iv_keyboardhandler_cb)),
	)
}

//export iv_keyboardhandler_cb
func iv_keyboardhandler_cb(text *C.char) {
	keyboardInput <- C.GoString(text)
}
