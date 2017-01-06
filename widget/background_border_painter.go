// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package widget

import (
	"github.com/tmacychen/UFG/framework"
	"github.com/tmacychen/UFG/framework/outer"
	"github.com/tmacychen/UFG/math"
	"github.com/tmacychen/UFG/widget/tools"
)

type BackgroundBorderPainterOuter interface {
	outer.Redrawer
}

type BackgroundBorderPainter struct {
	outer BackgroundBorderPainterOuter
	brush tools.Brush
	pen   tools.Pen
}

func (b *BackgroundBorderPainter) Init(outer BackgroundBorderPainterOuter) {
	b.outer = outer
	b.brush = tools.DefaultBrush
	b.pen = tools.DefaultPen
}

func (b *BackgroundBorderPainter) PaintBackground(c framework.Canvas, r math.Rect) {
	if b.brush.Color.A != 0 {
		w := b.pen.Width
		c.DrawRoundedRect(r, w, w, w, w, tools.TransparentPen, b.brush)
	}
}

func (b *BackgroundBorderPainter) PaintBorder(c framework.Canvas, r math.Rect) {
	if b.pen.Color.A != 0 && b.pen.Width != 0 {
		w := b.pen.Width
		c.DrawRoundedRect(r, w, w, w, w, b.pen, tools.TransparentBrush)
	}
}

func (b *BackgroundBorderPainter) BackgroundBrush() tools.Brush {
	return b.brush
}

func (b *BackgroundBorderPainter) SetBackgroundBrush(brush tools.Brush) {
	if b.brush != brush {
		b.brush = brush
		println("SetBackgroundBrush redraw")
		b.outer.Redraw()
	}
}

func (b *BackgroundBorderPainter) BorderPen() tools.Pen {
	return b.pen
}

func (b *BackgroundBorderPainter) SetBorderPen(pen tools.Pen) {
	if b.pen != pen {
		b.pen = pen
		b.outer.Redraw()
	}
}
