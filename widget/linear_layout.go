// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package widget

import (
	"github.com/tmacychen/UFG/controller"
	"github.com/tmacychen/UFG/framework"
	"github.com/tmacychen/UFG/math"
	"github.com/tmacychen/UFG/widget/tools"
)

type LinearLayoutOuter interface {
	controller.ContainerOuter
}

type LinearLayout struct {
	controller.Container
	BackgroundBorderPainter
	outer               LinearLayoutOuter
	direction           framework.Direction
	sizeMode            framework.SizeMode
	horizontalAlignment framework.HorizontalAlignment
	verticalAlignment   framework.VerticalAlignment
}

func (l *LinearLayout) Init(outer LinearLayoutOuter, theme framework.Theme) {

	l.Container.Init(outer, theme)
	l.BackgroundBorderPainter.Init(outer)
	l.outer = outer
	l.SetMouseEventTarget(true)
	println("linearlayout SetBackgroundBrush!")
	l.SetBackgroundBrush(tools.TransparentBrush)
	l.SetBorderPen(tools.TransparentPen)

	// Interface compliance test
	_ = framework.LinearLayout(l)
}

func (l *LinearLayout) LayoutChildren() {
	s := l.outer.Size().Contract(l.outer.Padding())
	o := l.outer.Padding().LT()
	children := l.outer.Children()
	major := 0
	if l.direction.RightToLeft() || l.direction.BottomToTop() {
		if l.direction.RightToLeft() {
			major = s.W
		} else {
			major = s.H
		}
	}
	for _, c := range children {
		cm := c.Control.Margin()
		cs := c.Control.DesiredSize(math.ZeroSize, s.Contract(cm).Max(math.ZeroSize))
		c.Control.SetSize(cs)

		// Calculate minor-axis alignment
		var minor int
		switch l.direction.Orientation() {
		case framework.Horizontal:
			switch l.verticalAlignment {
			case framework.AlignTop:
				minor = cm.T
			case framework.AlignMiddle:
				minor = (s.H - cs.H) / 2
			case framework.AlignBottom:
				minor = s.H - cs.H
			}
		case framework.Vertical:
			switch l.horizontalAlignment {
			case framework.AlignLeft:
				minor = cm.L
			case framework.AlignCenter:
				minor = (s.W - cs.W) / 2
			case framework.AlignRight:
				minor = s.W - cs.W
			}
		}

		// Peform layout
		switch l.direction {
		case framework.LeftToRight:
			major += cm.L
			c.Offset = math.Point{X: major, Y: minor}.Add(o)
			major += cs.W
			major += cm.R
			s.W -= cs.W + cm.W()
		case framework.RightToLeft:
			major -= cm.R
			c.Offset = math.Point{X: major - cs.W, Y: minor}.Add(o)
			major -= cs.W
			major -= cm.L
			s.W -= cs.W + cm.W()
		case framework.TopToBottom:
			major += cm.T
			c.Offset = math.Point{X: minor, Y: major}.Add(o)
			major += cs.H
			major += cm.B
			s.H -= cs.H + cm.H()
		case framework.BottomToTop:
			major -= cm.B
			c.Offset = math.Point{X: minor, Y: major - cs.H}.Add(o)
			major -= cs.H
			major -= cm.T
			s.H -= cs.H + cm.H()
		}
	}
}

func (l *LinearLayout) DesiredSize(min, max math.Size) math.Size {
	if l.sizeMode.Fill() {
		return max
	}

	bounds := min.Rect()
	children := l.outer.Children()

	horizontal := l.direction.Orientation().Horizontal()
	offset := math.Point{X: 0, Y: 0}
	for _, c := range children {
		cs := c.Control.DesiredSize(math.ZeroSize, max)
		cm := c.Control.Margin()
		cb := cs.Expand(cm).Rect().Offset(offset)
		if horizontal {
			offset.X += cb.W()
		} else {
			offset.Y += cb.H()
		}
		bounds = bounds.Union(cb)
	}

	return bounds.Size().Expand(l.outer.Padding()).Clamp(min, max)
}

func (l *LinearLayout) Direction() framework.Direction {
	return l.direction
}

func (l *LinearLayout) SetDirection(d framework.Direction) {
	if l.direction != d {
		l.direction = d
		l.outer.Relayout()
	}
}

func (l *LinearLayout) SizeMode() framework.SizeMode {
	return l.sizeMode
}

func (l *LinearLayout) SetSizeMode(mode framework.SizeMode) {
	if l.sizeMode != mode {
		l.sizeMode = mode
		l.outer.Relayout()
	}
}

func (l *LinearLayout) HorizontalAlignment() framework.HorizontalAlignment {
	return l.horizontalAlignment
}

func (l *LinearLayout) SetHorizontalAlignment(alignment framework.HorizontalAlignment) {
	if l.horizontalAlignment != alignment {
		l.horizontalAlignment = alignment
		l.outer.Relayout()
	}
}

func (l *LinearLayout) VerticalAlignment() framework.VerticalAlignment {
	return l.verticalAlignment
}

func (l *LinearLayout) SetVerticalAlignment(alignment framework.VerticalAlignment) {
	if l.verticalAlignment != alignment {
		l.verticalAlignment = alignment
		l.outer.Relayout()
	}
}
func (l *LinearLayout) Paint(c framework.Canvas) {
	r := l.Size().Rect()
	l.BackgroundBorderPainter.PaintBackground(c, r)
	l.PaintChildren.Paint(c)
	l.BackgroundBorderPainter.PaintBorder(c, r)
}
