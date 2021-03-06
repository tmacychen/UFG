// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package widget

import (
	"github.com/tmacychen/UFG/controller"
	"github.com/tmacychen/UFG/framework"
	"github.com/tmacychen/UFG/math"
)

type ButtonOuter interface {
	LinearLayoutOuter
	IsChecked() bool
	SetChecked(bool)
}

type Button struct {
	LinearLayout
	controller.Focusable

	outer      ButtonOuter
	label      framework.Label
	theme      framework.Theme
	buttonType framework.ButtonType
	checked    bool
}

func (b *Button) Init(outer ButtonOuter, theme framework.Theme) {
	b.LinearLayout.Init(outer, theme)
	b.Focusable.Init(outer)

	b.buttonType = framework.PushButton
	b.theme = theme
	b.outer = outer

	// Interface compliance test
	_ = framework.Button(b)
}

func (b *Button) Label() framework.Label {
	return b.label
}

func (b *Button) Text() string {
	if b.label != nil {
		return b.label.Text()
	} else {
		return ""
	}
}

func (b *Button) SetText(text string) {
	if b.Text() == text {
		return
	}
	if text == "" {
		if b.label != nil {
			b.RemoveChild(b.label)
			b.label = nil
		}
	} else {
		if b.label == nil {
			b.label = b.theme.CreateLabel()
			b.label.SetMargin(math.ZeroSpacing)

			b.AddChild(b.label)

		}
		b.label.SetText(text)
	}
}

func (b *Button) Type() framework.ButtonType {
	return b.buttonType
}

func (b *Button) SetType(buttonType framework.ButtonType) {
	if buttonType != b.buttonType {
		b.buttonType = buttonType
		b.outer.Redraw()
	}
}

func (b *Button) IsChecked() bool {
	return b.checked
}

func (b *Button) SetChecked(checked bool) {
	if checked != b.checked {
		b.checked = checked
		b.outer.Redraw()
	}
}

// InputEventHandler override
func (b *Button) Click(ev framework.MouseEvent) (consume bool) {
	println("button click")
	if ev.Button == framework.MouseButtonLeft {
		println("buttonLeft")
		if b.buttonType == framework.ToggleButton {
			b.outer.SetChecked(!b.outer.IsChecked())
		}
		b.InputEventHandler.Click(ev)
		return true
	}
	return b.InputEventHandler.Click(ev)
}

// InputEventHandler override
func (b *Button) KeyPress(ev framework.KeyboardEvent) (consume bool) {
	consume = b.InputEventHandler.KeyPress(ev)
	if ev.Key == framework.KeySpace || ev.Key == framework.KeyEnter {
		me := framework.MouseEvent{
			Button: framework.MouseButtonLeft,
		}
		return b.Click(me)
	}
	return
}
