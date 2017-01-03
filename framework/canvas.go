// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package framework

import (
	"github.com/tmacychen/UFG/math"
	"github.com/tmacychen/UFG/widget/tools"
)

type Canvas interface {
	Size() math.Size
	IsComplete() bool
	Complete()
	Push()
	Pop()
	AddClip(math.Rect)
	Clear(tools.Color)
	DrawCanvas(c Canvas, position math.Point)
	DrawTexture(t Texture, bounds math.Rect)
	DrawRunes(font Font, runes []rune, points []math.Point, color tools.Color)
	DrawLines(Polygon, tools.Pen)
	DrawPolygon(Polygon, tools.Pen, tools.Brush)
	DrawRect(math.Rect, tools.Brush)
	DrawRoundedRect(rect math.Rect, tl, tr, bl, br float32, p tools.Pen, b tools.Brush)
}
