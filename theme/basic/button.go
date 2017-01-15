package basic

import (
	"github.com/tmacychen/UFG/framework"
	"github.com/tmacychen/UFG/math"
	"github.com/tmacychen/UFG/widget"

	"github.com/tmacychen/UFG/widget/tools"

)

type Button struct {
	widget.Button
	theme *Theme
}

func createButton(theme *Theme) framework.Button {
	b := &Button{}
	b.Init(b, theme)
	b.theme = theme
	b.SetPadding(math.Spacing{L: 3, T: 3, R: 3, B: 3})
	b.SetMargin(math.Spacing{L: 3, T: 3, R: 3, B: 3})
	b.SetBackgroundBrush(theme.ButtonDefaultStyle.Brush)
	b.SetBorderPen(theme.ButtonDefaultStyle.Pen)
	b.OnMouseEnter(func(framework.MouseEvent) { b.Redraw() })
	b.OnMouseExit(func(framework.MouseEvent) { b.Redraw() })
	b.OnMouseDown(func(framework.MouseEvent) { b.Redraw() })
	b.OnMouseUp(func(framework.MouseEvent) { b.Redraw() })
	b.OnGainedFocus(b.Redraw)
	b.OnLostFocus(b.Redraw)
	return b
}

func (b *Button) Paint(c framework.Canvas) {
	pen := b.BorderPen()
	brush := b.BackgroundBrush()
	fontColor := b.theme.ButtonDefaultStyle.FontColor

	switch {
	case b.IsMouseDown(framework.MouseButtonLeft) && b.IsMouseOver():
		pen = b.theme.ButtonPressedStyle.Pen
		brush = b.theme.ButtonPressedStyle.Brush
		fontColor = b.theme.ButtonPressedStyle.FontColor
	case b.IsMouseOver():
		pen = b.theme.ButtonOverStyle.Pen
		brush = b.theme.ButtonOverStyle.Brush
		fontColor = b.theme.ButtonOverStyle.FontColor
	}

	if l := b.Label(); l != nil {
		l.SetColor(fontColor)
		l.SetSize(b.Size())
		l.SetVerticalAlignment(framework.AlignMiddle)
	}

	r := b.Size().Rect()

	c.DrawRoundedRect(r, 2, 2, 2, 2, tools.TransparentPen, brush)
	b.PaintChildren.Paint(c)
	c.DrawRoundedRect(r, 2, 2, 2, 2, pen, tools.TransparentBrush)

	if b.IsChecked() {
		pen = b.theme.HighlightStyle.Pen
		brush = b.theme.HighlightStyle.Brush
		c.DrawRoundedRect(r, 2.0, 2.0, 2.0, 2.0, pen, brush)
	}

	if b.HasFocus() {
		pen = b.theme.FocusedStyle.Pen
		brush = b.theme.FocusedStyle.Brush
		c.DrawRoundedRect(r.ContractI(int(pen.Width)), 3.0, 3.0, 3.0, 3.0, pen, brush)
	}

}
