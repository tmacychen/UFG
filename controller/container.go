// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"github.com/tmacychen/UFG/framework"
	"github.com/tmacychen/UFG/framework/outer"
	"github.com/tmacychen/UFG/math"
)

type ContainerNoControlOuter interface {
	framework.Container
	outer.PaintChilder
	outer.Painter
	outer.LayoutChildren
}

type ContainerControlableOuter interface {
	framework.Control
	ContainerNoControlOuter
}

type ContainerOuter interface {
	outer.Attachable
	outer.IsVisibler
	outer.LayoutChildren
	outer.Parenter
	outer.Sized
	framework.Container
}



type ContainerControlable struct {
	Attachable
	Container
	DrawPaint
	InputEventHandler
	Layoutable
	Paddable
	PaintChildren
	Visible
	Parentable

}


func (c *ContainerControlable) Init(outer ContainerControlableOuter, theme framework.Theme) {
	c.Attachable.Init(outer)
	c.DrawPaint.Init(outer, theme)
	c.InputEventHandler.Init(outer)
	c.Layoutable.Init(outer, theme)
	c.Paddable.Init(outer)
	c.PaintChildren.Init(outer)
	c.Parentable.Init(outer)
	c.Visible.Init(outer)
	c.Container.Init(outer,theme)
	// Interface compliance test
	_ = framework.Container(c)
}

type Container struct {
	outer              ContainerOuter
	children           framework.Children
	isMouseEventTarget bool
	relayoutSuspended  bool
}
func (c *Container) Init(outer ContainerOuter, theme framework.Theme) {

	c.outer = outer
	c.children = framework.Children{}
	outer.OnAttach(func() {
		for _, v := range c.children {
			v.Control.Attach()
		}
	})
	outer.OnDetach(func() {
		for _, v := range c.children {
			v.Control.Detach()
		}
	})
}


func (c *Container) SetMouseEventTarget(mouseEventTarget bool) {
	c.isMouseEventTarget = mouseEventTarget
}

func (c *Container) IsMouseEventTarget() bool {
	return c.isMouseEventTarget
}

// RelayoutSuspended returns true if adding or removing a child Control to this
// Container will not trigger a relayout of this Container. The default is false
// where any mutation will trigger a relayout.
func (c *Container) RelayoutSuspended() bool {
	return c.relayoutSuspended
}

// SetRelayoutSuspended enables or disables relayout of the Container on
// adding or removing a child Control to this Container.
func (c *Container) SetRelayoutSuspended(enable bool) {
	c.relayoutSuspended = true
}

// framework.Parent compliance
func (c *Container) Children() framework.Children {
	return c.children
}

// framework.Container compliance
func (c *Container) AddChild(child framework.Control) *framework.Child {
	return c.outer.AddChildAt(len(c.children), child)
}

// index 是child的位置，control是要加入的子组件
func (c *Container) AddChildAt(index int, control framework.Control) *framework.Child {
	if control.Parent() != nil {
		panic("Child already has a parent")
	}
	if index < 0 || index > len(c.children) {
		panic(fmt.Errorf("Index %d is out of bounds. Acceptable range: [%d - %d]",
			index, 0, len(c.children)))
	}

	child := &framework.Child{Control: control}

	c.children = append(c.children, nil)
	copy(c.children[index+1:], c.children[index:])
	c.children[index] = child

	control.SetParent(c.outer)
	if c.outer.Attached() {
		control.Attach()
	}
	if !c.relayoutSuspended {
		c.outer.Relayout()
	}
	println("Add a child")
	return child
}

func (c *Container) RemoveChild(control framework.Control) {
	for i := range c.children {
		if c.children[i].Control == control {
			c.outer.RemoveChildAt(i)
			return
		}
	}
	panic("Child not part of container")
}

func (c *Container) RemoveChildAt(index int) {
	child := c.children[index]
	c.children = append(c.children[:index], c.children[index+1:]...)
	child.Control.SetParent(nil)
	if c.outer.Attached() {
		child.Control.Detach()
	}
	if !c.relayoutSuspended {
		c.outer.Relayout()
	}
}

func (c *Container) RemoveAll() {
	for i := len(c.children) - 1; i >= 0; i-- {
		c.outer.RemoveChildAt(i)
	}
}

func (c *Container) ContainsPoint(p math.Point) bool {
	if !c.outer.IsVisible() || !c.outer.Size().Rect().Contains(p) {
		return false
	}
	for _, v := range c.children {
		if v.Control.ContainsPoint(p.Sub(v.Offset)) {
			return true
		}
	}
	if c.IsMouseEventTarget() {
		return true
	}
	return false
}
