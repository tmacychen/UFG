// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package framework

import "UFG/widget/tools"

type Label interface {
	Control
	Text() string
	SetText(string)
	Font() Font
	SetFont(Font)
	Color() tools.Color
	SetColor(tools.Color)
	Multiline() bool
	SetMultiline(bool)
	SetHorizontalAlignment(HorizontalAlignment)
	HorizontalAlignment() HorizontalAlignment
	SetVerticalAlignment(VerticalAlignment)
	VerticalAlignment() VerticalAlignment
}