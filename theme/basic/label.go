package basic

import (
	"github.com/tmacychen/UFG/framework"
	"github.com/tmacychen/UFG/widget"
	"github.com/tmacychen/UFG/math"
)

func createLabel(theme *Theme) framework.Label {
	l := &widget.Label{}
	l.Init(l, theme, theme.DefaultFont(), theme.LabelStyle.FontColor)
	l.SetMargin(math.Spacing{L: 3, T: 3, R: 3, B: 3})
	println("create label end")
	return l
}
