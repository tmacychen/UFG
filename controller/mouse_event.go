// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/tmacychen/UFG/framework"
	"github.com/tmacychen/UFG/math"
)

type MouseEvent struct {
	Button           framework.MouseButton
	State            framework.MouseState
	Point            math.Point // Local to the event receiver
	WindowPoint      math.Point
	Window           framework.Window
	ScrollX, ScrollY int
	Modifier         framework.KeyboardModifier
}
