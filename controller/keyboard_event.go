// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package controller

import "github.com/tmacychen/UFG/framework"

type KeyboardEvent struct {
	Key      framework.KeyboardKey
	Modifier framework.KeyboardModifier
}
