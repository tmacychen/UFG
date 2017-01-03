// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package controller

import "UFG/framework"

type KeyStrokeEvent struct {
	Character rune
	Modifier  framework.KeyboardModifier
}