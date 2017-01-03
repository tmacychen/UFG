// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
package widget

//
//import (
//	"UFG/framework"
//	"UFG/math"
//)
//
//type ButtonOuter interface {
//	LinearLayoutOuter
//	IsChecked() bool
//	SetChecked(bool)
//}
//
//type Button struct {
//	LinearLayout
//	Focusable
//
//	outer ButtonOuter
//	//	theme      theme.Theme
//	label      framework.Label
//	buttonType framework.ButtonType
//	checked    bool
//}
//
//func CreateButton(theme *Theme) framework.Button {
//	b := &Button{}
//	b.Init(b, theme)
//	//	b.theme = theme
//	b.SetPadding(math.Spacing{L: 3, T: 3, R: 3, B: 3})
//	b.SetMargin(math.Spacing{L: 3, T: 3, R: 3, B: 3})
//	b.SetBackgroundBrush(b.theme.ButtonDefaultStyle.Brush)
//	b.SetBorderPen(b.theme.ButtonDefaultStyle.Pen)
//	b.OnMouseEnter(func(framework.MouseEvent) { b.Redraw() })
//	b.OnMouseExit(func(framework.MouseEvent) { b.Redraw() })
//	b.OnMouseDown(func(framework.MouseEvent) { b.Redraw() })
//	b.OnMouseUp(func(framework.MouseEvent) { b.Redraw() })
//	b.OnGainedFocus(b.Redraw)
//	b.OnLostFocus(b.Redraw)
//	return b
//}
//
//// Button internal overrides
//func (b *Button) Paint(c framework.Canvas) {
//	pen := b.Button.BorderPen()
//	brush := b.Button.BackgroundBrush()
//	fontColor := b.theme.ButtonDefaultStyle.FontColor
//
//	switch {
//	case b.IsMouseDown(framework.MouseButtonLeft) && b.IsMouseOver():
//		pen = b.theme.ButtonPressedStyle.Pen
//		brush = b.theme.ButtonPressedStyle.Brush
//		fontColor = b.theme.ButtonPressedStyle.FontColor
//	case b.IsMouseOver():
//		pen = b.theme.ButtonOverStyle.Pen
//		brush = b.theme.ButtonOverStyle.Brush
//		fontColor = b.theme.ButtonOverStyle.FontColor
//	}
//
//	if l := b.Label(); l != nil {
//		l.SetColor(fontColor)
//	}
//
//	r := b.Size().Rect()
//
//	c.DrawRoundedRect(r, 2, 2, 2, 2, framework.TransparentPen, brush)
//
//	b.PaintChildren.Paint(c)
//
//	c.DrawRoundedRect(r, 2, 2, 2, 2, pen, framework.TransparentBrush)
//
//	if b.IsChecked() {
//		pen = b.theme.HighlightStyle.Pen
//		brush = b.theme.HighlightStyle.Brush
//		c.DrawRoundedRect(r, 2.0, 2.0, 2.0, 2.0, pen, brush)
//	}
//
//	if b.HasFocus() {
//		pen = b.theme.FocusedStyle.Pen
//		brush = b.theme.FocusedStyle.Brush
//		c.DrawRoundedRect(r.ContractI(int(pen.Width)), 3.0, 3.0, 3.0, 3.0, pen, brush)
//	}
//}
//func (b *Button) Init(outer ButtonOuter, theme framework.Theme) {
//	b.LinearLayout.Init(outer, theme)
//	b.Focusable.Init(outer)
//
//	b.buttonType = framework.PushButton
//	//	b.theme = theme
//	b.outer = outer
//
//	// Interface compliance test
//	_ = framework.Button(b)
//}
//
//func (b *Button) Label() framework.Label {
//	return b.label
//}
//
//func (b *Button) Text() string {
//	if b.label != nil {
//		return b.label.Text()
//	} else {
//		return ""
//	}
//}
//
//func (b *Button) SetText(text string) {
//	if b.Text() == text {
//		return
//	}
//	if text == "" {
//		if b.label != nil {
//			b.RemoveChild(b.label)
//			b.label = nil
//		}
//	} else {
//		if b.label == nil {
//			b.label = b.theme.CreateLabel()
//			b.label.SetMargin(math.ZeroSpacing)
//			b.AddChild(b.label)
//		}
//		b.label.SetText(text)
//	}
//}
//
//func (b *Button) Type() framework.ButtonType {
//	return b.buttonType
//}
//
//func (b *Button) SetType(buttonType framework.ButtonType) {
//	if buttonType != b.buttonType {
//		b.buttonType = buttonType
//		b.outer.Redraw()
//	}
//}
//
//func (b *Button) IsChecked() bool {
//	return b.checked
//}
//
//func (b *Button) SetChecked(checked bool) {
//	if checked != b.checked {
//		b.checked = checked
//		b.outer.Redraw()
//	}
//}
//
//// InputEventHandler override
//func (b *Button) Click(ev framework.MouseEvent) (consume bool) {
//	if ev.Button == framework.MouseButtonLeft {
//		if b.buttonType == framework.ToggleButton {
//			b.outer.SetChecked(!b.outer.IsChecked())
//		}
//		b.LinearLayout.Click(ev)
//		return true
//	}
//	return b.LinearLayout.Click(ev)
//}
//
//func (b *Button) KeyPress(ev framework.KeyboardEvent) (consume bool) {
//	consume = b.LinearLayout.KeyPress(ev)
//	if ev.Key == framework.KeySpace || ev.Key == framework.KeyEnter {
//		me := framework.MouseEvent{
//			Button: framework.MouseButtonLeft,
//		}
//		return b.Click(me)
//	}
//	return
//}
