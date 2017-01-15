// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package controller

import "github.com/tmacychen/UFG/framework"

type InputEventHandlerOuter interface{}

type InputEventHandler struct {
	outer         InputEventHandlerOuter
	isMouseOver   bool
	isMouseDown   map[framework.MouseButton]bool
	onClick       framework.Event
	onDoubleClick framework.Event
	onKeyPress    framework.Event
	onKeyStroke   framework.Event
	onMouseMove   framework.Event
	onMouseEnter  framework.Event
	onMouseExit   framework.Event
	onMouseDown   framework.Event
	onMouseUp     framework.Event
	onMouseScroll framework.Event
	onKeyDown     framework.Event
	onKeyUp       framework.Event
	onKeyRepeat   framework.Event
}

func (m *InputEventHandler) getOnClick() framework.Event {
	if m.onClick == nil {
		m.onClick = CreateEvent(m.Click)
	}
	return m.onClick
}

func (m *InputEventHandler) getOnDoubleClick() framework.Event {
	if m.onDoubleClick == nil {
		m.onDoubleClick = CreateEvent(m.DoubleClick)
	}
	return m.onDoubleClick
}

func (m *InputEventHandler) getOnKeyPress() framework.Event {
	if m.onKeyPress == nil {
		m.onKeyPress = CreateEvent(m.KeyPress)
	}
	return m.onKeyPress
}

func (m *InputEventHandler) getOnKeyStroke() framework.Event {
	if m.onKeyStroke == nil {
		m.onKeyStroke = CreateEvent(m.KeyStroke)
	}
	return m.onKeyStroke
}

func (m *InputEventHandler) getOnMouseMove() framework.Event {
	if m.onMouseMove == nil {
		m.onMouseMove = CreateEvent(m.MouseMove)
	}
	return m.onMouseMove
}

func (m *InputEventHandler) getOnMouseEnter() framework.Event {
	if m.onMouseEnter == nil {
		m.onMouseEnter = CreateEvent(m.MouseEnter)
	}
	return m.onMouseEnter
}

func (m *InputEventHandler) getOnMouseExit() framework.Event {
	if m.onMouseExit == nil {
		m.onMouseExit = CreateEvent(m.MouseExit)
	}
	return m.onMouseExit
}

func (m *InputEventHandler) getOnMouseDown() framework.Event {
	if m.onMouseDown == nil {
		m.onMouseDown = CreateEvent(m.MouseDown)
	}
	return m.onMouseDown
}

func (m *InputEventHandler) getOnMouseUp() framework.Event {
	if m.onMouseUp == nil {
		m.onMouseUp = CreateEvent(m.MouseUp)
	}
	return m.onMouseUp
}

func (m *InputEventHandler) getOnMouseScroll() framework.Event {
	if m.onMouseScroll == nil {
		m.onMouseScroll = CreateEvent(m.MouseScroll)
	}
	return m.onMouseScroll
}

func (m *InputEventHandler) getOnKeyDown() framework.Event {
	if m.onKeyDown == nil {
		m.onKeyDown = CreateEvent(m.KeyDown)
	}
	return m.onKeyDown
}

func (m *InputEventHandler) getOnKeyUp() framework.Event {
	if m.onKeyUp == nil {
		m.onKeyUp = CreateEvent(m.KeyUp)
	}
	return m.onKeyUp
}

func (m *InputEventHandler) getOnKeyRepeat() framework.Event {
	if m.onKeyRepeat == nil {
		m.onKeyRepeat = CreateEvent(m.KeyRepeat)
	}
	return m.onKeyRepeat
}

func (m *InputEventHandler) Init(outer InputEventHandlerOuter) {
	m.outer = outer
	m.isMouseDown = make(map[framework.MouseButton]bool)
}

func (m *InputEventHandler) Click(ev framework.MouseEvent) (consume bool) {
	m.getOnClick().Fire(ev)
	return false
}

func (m *InputEventHandler) DoubleClick(ev framework.MouseEvent) (consume bool) {
	m.getOnDoubleClick().Fire(ev)
	return false
}

func (m *InputEventHandler) KeyPress(ev KeyboardEvent) (consume bool) {
	m.getOnKeyPress().Fire(ev)
	return false
}

func (m *InputEventHandler) KeyStroke(ev KeyStrokeEvent) (consume bool) {
	m.getOnKeyStroke().Fire(ev)
	return false
}

func (m *InputEventHandler) MouseScroll(ev framework.MouseEvent) (consume bool) {
	m.getOnMouseScroll().Fire(ev)
	return false
}

func (m *InputEventHandler) MouseMove(ev framework.MouseEvent) {
	m.getOnMouseMove().Fire(ev)
}

func (m *InputEventHandler) MouseEnter(ev framework.MouseEvent) {
	m.isMouseOver = true
	m.getOnMouseEnter().Fire(ev)
}

func (m *InputEventHandler) MouseExit(ev framework.MouseEvent) {
	m.isMouseOver = false
	m.getOnMouseExit().Fire(ev)
}

func (m *InputEventHandler) MouseDown(ev framework.MouseEvent) {
	m.isMouseDown[ev.Button] = true
	m.getOnMouseDown().Fire(ev)
}

func (m *InputEventHandler) MouseUp(ev framework.MouseEvent) {
	m.isMouseDown[ev.Button] = false
	m.getOnMouseUp().Fire(ev)
}

func (m *InputEventHandler) KeyDown(ev KeyboardEvent) {
	m.getOnKeyDown().Fire(ev)
}

func (m *InputEventHandler) KeyUp(ev KeyboardEvent) {
	m.getOnKeyUp().Fire(ev)
}

func (m *InputEventHandler) KeyRepeat(ev KeyboardEvent) {
	m.getOnKeyRepeat().Fire(ev)
}

func (m *InputEventHandler) OnClick(f func(framework.MouseEvent)) framework.EventSubscription {
	return m.getOnClick().Listen(f)
}

func (m *InputEventHandler) OnDoubleClick(f func(framework.MouseEvent)) framework.EventSubscription {
	return m.getOnDoubleClick().Listen(f)
}

func (m *InputEventHandler) OnKeyPress(f func(KeyboardEvent)) framework.EventSubscription {
	return m.getOnKeyPress().Listen(f)
}

func (m *InputEventHandler) OnKeyStroke(f func(KeyStrokeEvent)) framework.EventSubscription {
	return m.getOnKeyStroke().Listen(f)
}

func (m *InputEventHandler) OnMouseMove(f func(framework.MouseEvent)) framework.EventSubscription {
	return m.getOnMouseMove().Listen(f)
}

func (m *InputEventHandler) OnMouseEnter(f func(framework.MouseEvent)) framework.EventSubscription {
	return m.getOnMouseEnter().Listen(f)
}

func (m *InputEventHandler) OnMouseExit(f func(framework.MouseEvent)) framework.EventSubscription {
	return m.getOnMouseExit().Listen(f)
}

func (m *InputEventHandler) OnMouseDown(f func(framework.MouseEvent)) framework.EventSubscription {
	return m.getOnMouseDown().Listen(f)
}

func (m *InputEventHandler) OnMouseUp(f func(framework.MouseEvent)) framework.EventSubscription {
	return m.getOnMouseUp().Listen(f)
}

func (m *InputEventHandler) OnMouseScroll(f func(framework.MouseEvent)) framework.EventSubscription {
	return m.getOnMouseScroll().Listen(f)
}

func (m *InputEventHandler) OnKeyDown(f func(KeyboardEvent)) framework.EventSubscription {
	return m.getOnKeyDown().Listen(f)
}

func (m *InputEventHandler) OnKeyUp(f func(KeyboardEvent)) framework.EventSubscription {
	return m.getOnKeyUp().Listen(f)
}

func (m *InputEventHandler) OnKeyRepeat(f func(KeyboardEvent)) framework.EventSubscription {
	return m.getOnKeyRepeat().Listen(f)
}

func (m *InputEventHandler) IsMouseOver() bool {
	return m.isMouseOver
}

func (m *InputEventHandler) IsMouseDown(button framework.MouseButton) bool {
	return m.isMouseDown[button]
}
