// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package widget

import (
	"strings"
	"github.com/tmacychen/UFG/controller"
	"github.com/tmacychen/UFG/framework"
	"github.com/tmacychen/UFG/math"
	"github.com/tmacychen/UFG/widget/tools"

)

type LabelOuter interface {
	controller.ControlOuter
}

type Label struct {
	controller.Control

	outer               LabelOuter
	font                framework.Font
	color               tools.Color
	horizontalAlignment framework.HorizontalAlignment
	verticalAlignment   framework.VerticalAlignment
	multiline           bool
	text                string
}





func (l *Label) Init(outer LabelOuter, theme framework.Theme, font framework.Font, color tools.Color) {
	if font == nil {
		panic("Cannot create a label with a nil font")
	}
	l.Control.Init(outer, theme)
	l.outer = outer
	l.font = font
	l.color = color
	l.horizontalAlignment = framework.AlignLeft
	l.verticalAlignment = framework.AlignMiddle
	// Interface compliance test
	_ = framework.Label(l)
}

func (l *Label) Text() string {
	return l.text
}

func (l *Label) SetText(text string) {
	if l.text != text {
		l.text = text
		l.outer.Relayout()
	}
}

//
func (l *Label) Font() framework.Font {
	return l.font
}

//}
//
func (l *Label) SetFont(font framework.Font) {
	if l.font != font {
		l.font = font
		l.Relayout()
	}
}
func (l *Label) Color() tools.Color {
	return l.color
}

func (l *Label) SetColor(color tools.Color) {
	if l.color != color {
		l.color = color
		l.outer.Redraw()
	}
}

func (l *Label) Multiline() bool {
	return l.multiline
}

func (l *Label) SetMultiline(multiline bool) {
	if l.multiline != multiline {
		l.multiline = multiline
		l.outer.Relayout()
	}
}

func (l *Label) DesiredSize(min, max math.Size) math.Size {
	t := l.text
	if !l.multiline {
		t = strings.Replace(t, "\n", " ", -1)
	}
	s := l.font.Measure(&framework.TextBlock{Runes: []rune(t)})
	return s.Clamp(min, max)
}

func (l *Label) SetHorizontalAlignment(horizontalAlignment framework.HorizontalAlignment) {
	if l.horizontalAlignment != horizontalAlignment {
		l.horizontalAlignment = horizontalAlignment
		l.Redraw()
	}
}

func (l *Label) HorizontalAlignment() framework.HorizontalAlignment {
	return l.horizontalAlignment
}

func (l *Label) SetVerticalAlignment(verticalAlignment framework.VerticalAlignment) {
	if l.verticalAlignment != verticalAlignment {
		l.verticalAlignment = verticalAlignment
		l.Redraw()
	}
}

func (l *Label) VerticalAlignment() framework.VerticalAlignment {
	return l.verticalAlignment
}

// parts.DrawPaint overrides
func (l *Label) Paint(c framework.Canvas) {
	r := l.outer.Size().Rect()
	t := l.text
	if !l.multiline {
		t = strings.Replace(t, "\n", " ", -1)
	}

	runes := []rune(t)
	offsets := l.font.Layout(&framework.TextBlock{
		Runes:     runes,
		AlignRect: r,
		H:         l.horizontalAlignment,
		V:         l.verticalAlignment,
	})
	c.DrawRunes(l.font, runes, offsets, l.color)
}
