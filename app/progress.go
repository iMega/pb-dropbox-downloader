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

type progress struct {
	Text  string
	Value int
}

func (app *App) NewProgressBar(title, text string, ch <-chan progress) {
	app.API.OpenProgressbar(title, text, 3, 0)

	go func() {
		defer app.API.CloseProgressbar()

		for p := range ch {
			app.API.UpdateProgressbar(p.Text, p.Value)

			if p.Value >= 100 {
				return
			}
		}
	}()
}
