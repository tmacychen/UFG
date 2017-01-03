// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package controller

import (
	"UFG/framework"
	"UFG/framework/outer"
	"UFG/math"
)

type ControlOuter interface {
	framework.Control
	outer.Painter
	outer.Redrawer
	outer.Relayouter
}

type Control struct {
	Attachable
	DrawPaint
	InputEventHandler
	Layoutable
	Parentable
	Visible
}

func (c *Control) Init(outer ControlOuter, theme framework.Theme) {
	c.Attachable.Init(outer)
	c.DrawPaint.Init(outer, theme)
	c.Layoutable.Init(outer, theme)
	c.InputEventHandler.Init(outer)
	c.Parentable.Init(outer)
	c.Visible.Init(outer)

	// Interface compliance test
	_ = framework.Control(c)
}

func (c *Control) DesiredSize(min, max math.Size) math.Size {
	return max
}

func (c *Control) ContainsPoint(p math.Point) bool {
	return c.IsVisible() && c.Size().Rect().Contains(p)
}
