// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package widget

import (
	"UFG/widget/tools"
)

type Style struct {
	FontColor tools.Color
	Brush     tools.Brush
	Pen       tools.Pen
}

func CreateStyle(fontColor, brushColor, penColor tools.Color, penWidth float32) Style {
	return Style{
		FontColor: fontColor,
		Pen:       tools.CreatePen(penWidth, penColor),
		Brush:     tools.CreateBrush(brushColor),
	}
}
