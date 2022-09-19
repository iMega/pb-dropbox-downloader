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

package lifecycle

import (
	"context"
	"fmt"
	"pb-dropbox-downloader/app"

	ink "github.com/dennwc/inkview"
)

type Lifecycle struct {
	Wrapped *app.App
}

func (lc *Lifecycle) Init() error {
	ctx := context.TODO()
	if err := lc.Wrapped.Init(ctx); err != nil {
		return fmt.Errorf("failed to init an app, %w", err)
	}

	return nil
}

func (lc *Lifecycle) Close() error {
	lc.Wrapped.Close()

	return nil
}

func (lc *Lifecycle) Draw() {}

func (lc *Lifecycle) Key(e ink.KeyEvent) bool { return false }

func (lc *Lifecycle) Pointer(e ink.PointerEvent) bool {
	lc.Wrapped.Sync()
	// ink.Exit()

	return true
}

func (lc *Lifecycle) Touch(e ink.TouchEvent) bool { return false }

func (lc *Lifecycle) Orientation(o ink.Orientation) bool { return false }
