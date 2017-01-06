// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/tmacychen/UFG/framework"
	"github.com/tmacychen/UFG/framework/outer"
)

type PaintChildrenOuter interface {
	framework.Container
	outer.PaintChilder
	outer.Sized
}

type PaintChildren struct {
	outer PaintChildrenOuter
}

func (p *PaintChildren) Init(outer PaintChildrenOuter) {
	p.outer = outer
}

func (p *PaintChildren) Paint(c framework.Canvas) {
	for i, v := range p.outer.Children() {
		if v.Control.IsVisible() {
			c.Push()
			c.AddClip(v.Control.Size().Rect().Offset(v.Offset))
			p.outer.PaintChild(c, v, i)
			c.Pop()
			println("print children")
		}
	}
}

func (p *PaintChildren) PaintChild(c framework.Canvas, child *framework.Child, idx int) {
	if canvas := child.Control.Draw(); canvas != nil {
		c.DrawCanvas(canvas, child.Offset)
	}
}
